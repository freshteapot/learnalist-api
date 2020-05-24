package api_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/freshteapot/learnalist-api/server/api/api"
	"github.com/freshteapot/learnalist-api/server/api/uuid"
	"github.com/freshteapot/learnalist-api/server/mocks"
	"github.com/freshteapot/learnalist-api/server/pkg/event"
	"github.com/freshteapot/learnalist-api/server/pkg/oauth"
	"github.com/freshteapot/learnalist-api/server/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing user delete endpoint", func() {
	var (
		logger         *logrus.Logger
		hook           *test.Hook
		userUUID       string
		endpoint       string
		datastore      *mocks.Datastore
		userManagement *mocks.Management
		session        user.UserSession
		user           *uuid.User
		manager        *api.Manager
		req            *http.Request
		rec            *httptest.ResponseRecorder
		e              *echo.Echo
		c              echo.Context
	)

	AfterEach(emptyDatabase)

	BeforeEach(func() {
		logger, hook = test.NewNullLogger()
		datastore = &mocks.Datastore{}
		userManagement = &mocks.Management{}
		acl := &mocks.Acl{}
		oauthHandlers := oauth.Handlers{}
		testHugoHelper := &mocks.HugoSiteBuilder{}

		manager = api.NewManager(datastore, userManagement, acl, "", testHugoHelper, oauthHandlers, logger)

		userUUID = "fake-123"
		session.Token = "fake-token"
		session.UserUUID = userUUID
		endpoint = fmt.Sprintf("/api/v1/user/delete/%s", userUUID)

		user = &uuid.User{
			Uuid: "fake-123",
		}

		req, rec = setupFakeEndpoint(http.MethodPost, endpoint, "")
		e = echo.New()
		c = e.NewContext(req, rec)
		c.SetPath("/api/v1/alist/:uuid")
		c.Set("loggedInUser", *user)
	})

	It("The user to delete is not the same as the user logged in", func() {
		c.SetParamNames("uuid")
		c.SetParamValues("fake-345")
		manager.V1DeleteUser(c)
		Expect(rec.Code).To(Equal(http.StatusForbidden))
	})

	When("User is deleting themselves", func() {
		It("Issue deleting from the system", func() {
			c.SetParamNames("uuid")
			c.SetParamValues(userUUID)
			want := errors.New("fail")
			userManagement.On("DeleteUserFromDB", userUUID).Return(want)
			manager.V1DeleteUser(c)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Sadly, our service has taken a nap."}`))
		})

		It("Successfully deleted user", func() {
			c.SetParamNames("uuid")
			c.SetParamValues(userUUID)
			userManagement.On("DeleteUserFromDB", userUUID).Return(nil)
			manager.V1DeleteUser(c)

			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"User has been removed"}`))
			// Check event
			Expect(hook.LastEntry().Data["event"]).To(Equal(event.UserDeleted))
			Expect(hook.LastEntry().Data["user_uuid"]).To(Equal(userUUID))
		})
	})
})
