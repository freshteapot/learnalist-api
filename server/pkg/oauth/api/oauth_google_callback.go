package api

import (
	"bytes"
	"net/http"

	"github.com/freshteapot/learnalist-api/server/api/i18n"
	"github.com/freshteapot/learnalist-api/server/pkg/api"
	"github.com/freshteapot/learnalist-api/server/pkg/authenticate"
	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/oauth"
	"github.com/freshteapot/learnalist-api/server/pkg/openapi"
	"github.com/freshteapot/learnalist-api/server/pkg/user"
	"github.com/freshteapot/learnalist-api/server/pkg/utils"
	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func (s OauthService) V1OauthGoogleCallback(c echo.Context) error {
	logContext := s.logContext.WithFields(logrus.Fields{
		"endpoint": "callback",
		"idp":      oauth.IDPKeyGoogle,
	})

	googleConfig := s.oauthHandlers.Google
	userSession := s.userSession
	userFromIDP := s.userFromIDP

	r := c.Request()
	challenge := r.FormValue("state")
	code := r.FormValue("code")
	// Confirm the challenge is valid
	has, err := userSession.IsChallengeValid(challenge)
	if err != nil {
		logContext.WithFields(logrus.Fields{
			"error": err,
		}).Error("Invalid challenge")
		return c.String(http.StatusInternalServerError, i18n.InternalServerErrorFunny)
	}

	if !has {
		return c.String(http.StatusBadRequest, "Invalid code / challenge, please try to login again")
	}

	// Exchange the code for the token
	token, err := googleConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		// Log
		response := "Exhange of code to token failed"
		logContext.WithFields(logrus.Fields{
			"error": err,
		}).Error(response)

		return c.String(http.StatusBadRequest, response)
	}
	// At this point we have the id_token, we can extract out the extUserUUID (sub).
	// Could just extract it out via

	extUserUUID, err := googleConfig.GetUserUUIDFromIDP(openapi.HttpUserLoginIdpInput{
		IdToken: token.Extra("id_token").(string),
		Idp:     oauth.IDPKeyGoogle,
	})
	if err != nil {
		logContext.WithFields(logrus.Fields{
			"event":  "idp-token-verification",
			"method": "google.GetUserUUIDFromIDP",
			"error":  err,
		}).Error("Issue in google callback")
		return c.JSON(http.StatusForbidden, api.HTTPAccessDeniedResponse)
	}

	//// Look up the user based on their id and association with google.
	userUUID, err := userFromIDP.Lookup(oauth.IDPKeyGoogle, user.IDPKindUserID, extUserUUID)
	if err != nil {
		if err != utils.ErrNotFound {
			logContext.WithFields(logrus.Fields{
				"event":  "idp-lookup-user-info",
				"method": "userFromIDP.Lookup",
				"error":  err,
			}).Error("Issue in google callback")
			return c.String(http.StatusInternalServerError, i18n.InternalServerErrorFunny)
		}

		// Create a user
		userUUID, err = userFromIDP.Register(oauth.IDPKeyGoogle, user.IDPKindUserID, extUserUUID, []byte(``))
		if err != nil {
			logContext.WithFields(logrus.Fields{
				"event": "idp-register-user",
				"error": err,
			}).Error("Failed to register new user via idp")
			return c.String(http.StatusInternalServerError, i18n.InternalServerErrorFunny)
		}

		// TODO I wonder what I do want here?
		pref := user.UserPreference{}
		event.GetBus().Publish(event.TopicMonolog, event.Eventlog{
			Kind: event.ApiUserRegister,
			Data: event.EventNewUser{
				UUID: userUUID,
				Kind: event.KindUserRegisterIDPGoogle,
				Data: pref,
			},
		})
	}

	// Create a session for the user
	userSessionToken := guuid.New()
	session := user.UserSession{
		Token:     userSessionToken.String(),
		UserUUID:  userUUID,
		Challenge: challenge,
	}

	err = userSession.Activate(session)
	if err != nil {
		logContext.WithFields(logrus.Fields{
			"event": "idp-session-activate",
			"error": err,
		}).Error("Failed to activate session")
		return c.String(http.StatusInternalServerError, i18n.InternalServerErrorFunny)
	}

	event.GetBus().Publish(event.TopicMonolog, event.Eventlog{
		Kind: event.ApiUserLogin,
		Data: event.EventUser{
			UUID: userUUID,
			Kind: event.KindUserLoginIDPGoogle,
		},
	})

	// Removed logic around storing refresh tokens, currently not in use
	vars := make(map[string]interface{})
	vars["token"] = session.Token
	vars["userUUID"] = userUUID
	vars["refreshRedirectURL"] = "/welcome.html"
	vars["idp"] = oauth.IDPKeyGoogle

	var tpl bytes.Buffer
	oauthCallbackHtml200.Execute(&tpl, vars)

	cookie := authenticate.NewLoginCookie(session.Token)
	c.SetCookie(cookie)
	return c.HTMLBlob(http.StatusOK, tpl.Bytes())
}
