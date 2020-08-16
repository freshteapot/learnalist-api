package api_test

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/freshteapot/learnalist-api/server/api/alist"
	"github.com/freshteapot/learnalist-api/server/api/i18n"
	"github.com/freshteapot/learnalist-api/server/api/uuid"
	"github.com/freshteapot/learnalist-api/server/mocks"
	aclKeys "github.com/freshteapot/learnalist-api/server/pkg/acl/keys"
	"github.com/freshteapot/learnalist-api/server/pkg/api"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Testing Sharing endpoints", func() {
	AfterEach(emptyDatabase)

	When("/share/readaccess", func() {
		var datastore *mocks.Datastore
		var acl *mocks.Acl
		var userA *uuid.User
		var userB *uuid.User
		var method string
		var uri string
		var e *echo.Echo

		BeforeEach(func() {
			datastore = &mocks.Datastore{}
			acl = &mocks.Acl{}
			m.Datastore = datastore
			m.Acl = acl

			testHugoHelper := &mocks.HugoSiteBuilder{}
			testHugoHelper.On("WriteList", mock.Anything)
			testHugoHelper.On("WriteListsByUser", mock.Anything, mock.Anything)
			testHugoHelper.On("WritePublicLists", mock.Anything)
			testHugoHelper.On("DeleteList", mock.Anything).Return(nil)
			m.HugoHelper = testHugoHelper

			userA = &uuid.User{
				Uuid: "fake-123",
			}
			userB = &uuid.User{
				Uuid: "fake-456",
			}

			method = http.MethodPost
			uri = "/api/v1/share/readaccess"
			e = echo.New()
		})

		It("Invalid json input", func() {
			input := ""
			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)
			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Your input is invalid json."}`))
		})

		It("Valid json input, invalid action", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userB.Uuid,
				AlistUUID: "fakeList",
				Action:    "keys-to-the-castle",
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)
			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)
			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Please refer to the documentation on sharing a list"}`))
		})

		It("Server error: failed to store in storage", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userB.Uuid,
				AlistUUID: "fakeList",
				Action:    aclKeys.ActionGrant,
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			datastore.On("GetAlist", mock.Anything).Return(alist.Alist{}, errors.New("Fail"))

			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Sadly, our service has taken a nap."}`))
		})

		It("List not found", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userB.Uuid,
				AlistUUID: "fakeList",
				Action:    aclKeys.ActionGrant,
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			datastore.On("GetAlist", mock.Anything).Return(alist.Alist{}, errors.New(i18n.SuccessAlistNotFound))

			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusNotFound))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"List not found."}`))
		})

		It("List found, but the user setting the share is not the owner of the list", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userB.Uuid,
				AlistUUID: "fakeList",
				Action:    aclKeys.ActionGrant,
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			aList := alist.NewTypeV1()
			aList.User.Uuid = userB.Uuid
			datastore.On("GetAlist", mock.Anything).Return(aList, nil)

			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusForbidden))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Access Denied"}`))
		})

		It("List found, you cant share with yourself", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userA.Uuid,
				AlistUUID: "fakeList",
				Action:    aclKeys.ActionGrant,
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			aList := alist.NewTypeV1()
			aList.User.Uuid = userA.Uuid
			aList.Info.SharedWith = aclKeys.SharedWithPublic
			datastore.On("GetAlist", mock.Anything).Return(aList, nil)

			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusUnprocessableEntity))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Today, we dont let you share with yourself"}`))
		})

		It("List found, the user you want to share with doesnt exist", func() {
			inputGrant := &api.HttpShareListWithUserInput{
				UserUUID:  userB.Uuid,
				AlistUUID: "fakeList",
				Action:    aclKeys.ActionGrant,
			}
			a, _ := json.Marshal(inputGrant)
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			aList := alist.NewTypeV1()
			aList.User.Uuid = userA.Uuid
			aList.Info.SharedWith = aclKeys.SharedWithPublic
			datastore.On("GetAlist", mock.Anything).Return(aList, nil)
			datastore.On("UserExists", userB.Uuid).Return(false)

			m.V1ShareListReadAccess(c)

			Expect(rec.Code).To(Equal(http.StatusNotFound))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"User not found."}`))
		})

		Context("Success, we will share", func() {
			It("Grant user read access", func() {
				inputGrant := &api.HttpShareListWithUserInput{
					UserUUID:  userB.Uuid,
					AlistUUID: "fakeList",
					Action:    aclKeys.ActionGrant,
				}
				a, _ := json.Marshal(inputGrant)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.SharedWithPublic
				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("UserExists", userB.Uuid).Return(true)
				acl.On("GrantUserListReadAccess", inputGrant.AlistUUID, inputGrant.UserUUID).Return(nil)

				m.V1ShareListReadAccess(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"user_uuid":"fake-456","alist_uuid":"fakeList","action":"grant"}`))
			})

			It("Revoke user read access", func() {
				inputRevoke := &api.HttpShareListWithUserInput{
					UserUUID:  userB.Uuid,
					AlistUUID: "fakeList",
					Action:    aclKeys.ActionRevoke,
				}
				a, _ := json.Marshal(inputRevoke)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.SharedWithPublic
				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("UserExists", userB.Uuid).Return(true)
				acl.On("RevokeUserListReadAccess", inputRevoke.AlistUUID, inputRevoke.UserUUID).Return(nil)

				m.V1ShareListReadAccess(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"user_uuid":"fake-456","alist_uuid":"fakeList","action":"revoke"}`))
			})
		})
	})
	When("/share", func() {
		var datastore *mocks.Datastore
		var acl *mocks.Acl
		var userA *uuid.User
		var userB *uuid.User
		var method string
		var uri string
		var e *echo.Echo

		BeforeEach(func() {
			datastore = &mocks.Datastore{}
			acl = &mocks.Acl{}
			m.Datastore = datastore
			m.Acl = acl

			testHugoHelper := &mocks.HugoSiteBuilder{}
			testHugoHelper.On("WriteList", mock.Anything)
			testHugoHelper.On("WriteListsByUser", mock.Anything, mock.Anything)
			testHugoHelper.On("WritePublicLists", mock.Anything)
			testHugoHelper.On("DeleteList", mock.Anything).Return(nil)
			m.HugoHelper = testHugoHelper

			userA = &uuid.User{
				Uuid: "fake-123",
			}
			userB = &uuid.User{
				Uuid: "fake-456",
			}

			method = http.MethodPost
			uri = "/api/v1/share/alist"
			e = echo.New()
		})

		It("Invalid json input", func() {
			input := ""
			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)
			m.V1ShareAlist(c)

			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Your input is invalid json."}`))
		})

		It("Valid json input, invalid action", func() {
			inputBadAction := &api.HttpShareListInput{
				AlistUUID: "fakeList",
				Action:    "keys-to-the-castle",
			}
			a, _ := json.Marshal(inputBadAction)
			input := string(a)
			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)
			m.V1ShareAlist(c)

			Expect(rec.Code).To(Equal(http.StatusBadRequest))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Please refer to the documentation on sharing a list"}`))
		})

		It("List not found", func() {
			a, _ := json.Marshal(&api.HttpShareListInput{
				AlistUUID: "fakeList",
				Action:    aclKeys.SharedWithPublic,
			})
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			datastore.On("GetAlist", mock.Anything).Return(alist.Alist{}, errors.New(i18n.SuccessAlistNotFound))

			m.V1ShareAlist(c)

			Expect(rec.Code).To(Equal(http.StatusNotFound))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"List not found."}`))
		})

		It("List found, but the user setting the share is not the owner of the list", func() {
			a, _ := json.Marshal(&api.HttpShareListInput{
				AlistUUID: "fakeList",
				Action:    aclKeys.SharedWithPublic,
			})
			input := string(a)

			req, rec := setupFakeEndpoint(method, uri, input)
			c := e.NewContext(req, rec)
			c.Set("loggedInUser", *userA)

			aList := alist.NewTypeV1()
			aList.Uuid = "fake-123"
			aList.User.Uuid = userB.Uuid
			datastore.On("GetAlist", mock.Anything).Return(aList, nil)

			m.V1ShareAlist(c)

			Expect(rec.Code).To(Equal(http.StatusForbidden))
			Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"Access Denied"}`))
		})

		Context("Success, we will share the list", func() {
			It("With the public", func() {
				inputObject := &api.HttpShareListInput{
					AlistUUID: "fakeList",
					Action:    aclKeys.SharedWithPublic,
				}
				a, _ := json.Marshal(inputObject)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.Uuid = inputObject.AlistUUID
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.NotShared

				returnAlist := alist.NewTypeV1()
				returnAlist.Uuid = inputObject.AlistUUID
				returnAlist.User.Uuid = userA.Uuid
				returnAlist.Info.SharedWith = aclKeys.SharedWithPublic

				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("SaveAlist", http.MethodPut, returnAlist).Return(returnAlist, nil)
				datastore.On("GetAllListsByUser", userA.Uuid).Return([]alist.ShortInfo{}, nil)
				datastore.On("GetPublicLists").Return([]alist.ShortInfo{})

				m.V1ShareAlist(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				CheckMessageResponse(rec, i18n.ApiShareListSuccessWithPublic)
			})

			It("Privately", func() {
				inputObject := &api.HttpShareListInput{
					AlistUUID: "fakeList",
					Action:    aclKeys.NotShared,
				}
				a, _ := json.Marshal(inputObject)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.Uuid = inputObject.AlistUUID
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.SharedWithPublic

				returnAlist := alist.NewTypeV1()
				returnAlist.Uuid = inputObject.AlistUUID
				returnAlist.User.Uuid = userA.Uuid
				returnAlist.Info.SharedWith = aclKeys.NotShared

				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("SaveAlist", http.MethodPut, returnAlist).Return(returnAlist, nil)
				datastore.On("GetAllListsByUser", userA.Uuid).Return([]alist.ShortInfo{}, nil)
				datastore.On("GetPublicLists").Return([]alist.ShortInfo{}, nil)
				m.V1ShareAlist(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"List is now private to the owner"}`))
			})

			It("With friends", func() {
				inputObject := &api.HttpShareListInput{
					AlistUUID: "fakeList",
					Action:    aclKeys.SharedWithFriends,
				}
				a, _ := json.Marshal(inputObject)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.Uuid = inputObject.AlistUUID
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.NotShared

				returnAlist := alist.NewTypeV1()
				returnAlist.Uuid = inputObject.AlistUUID
				returnAlist.User.Uuid = userA.Uuid
				returnAlist.Info.SharedWith = aclKeys.SharedWithFriends

				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("SaveAlist", http.MethodPut, returnAlist).Return(returnAlist, nil)
				datastore.On("GetAllListsByUser", userA.Uuid).Return([]alist.ShortInfo{}, nil)
				datastore.On("GetPublicLists").Return([]alist.ShortInfo{}, nil)
				m.V1ShareAlist(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"List is now private to the owner and those granted access"}`))
			})

			It("With friends when already set", func() {
				inputObject := &api.HttpShareListInput{
					AlistUUID: "fakeList",
					Action:    aclKeys.SharedWithFriends,
				}
				a, _ := json.Marshal(inputObject)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.Uuid = inputObject.AlistUUID
				aList.User.Uuid = userA.Uuid
				aList.Info.SharedWith = aclKeys.SharedWithFriends

				returnAlist := alist.NewTypeV1()
				returnAlist.Uuid = inputObject.AlistUUID
				returnAlist.User.Uuid = userA.Uuid
				returnAlist.Info.SharedWith = aclKeys.SharedWithFriends

				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("SaveAlist", http.MethodPut, returnAlist).Return(returnAlist, nil)
				datastore.On("GetAllListsByUser", userA.Uuid).Return([]alist.ShortInfo{}, nil)
				datastore.On("GetPublicLists").Return([]alist.ShortInfo{}, nil)
				m.V1ShareAlist(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"message":"No change made"}`))
			})
		})
		/*
			It("", func() {

				inputRevoke := &api.HttpShareListWithUserInput{
					UserUUID:  userB.Uuid,
					AlistUUID: "fakeList",
					Action:    aclKeys.ActionRevoke,
				}
				a, _ := json.Marshal(inputRevoke)
				input := string(a)

				req, rec := setupFakeEndpoint(method, uri, input)
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *userA)

				aList := alist.NewTypeV1()
				aList.User.Uuid = userA.Uuid
				datastore.On("GetAlist", mock.Anything).Return(aList, nil)
				datastore.On("UserExists", userB.Uuid).Return(true)
				acl.On("RevokeUserListReadAccess", inputRevoke.AlistUUID, inputRevoke.UserUUID).Return(nil)

				m.V1ShareAlist(c)

				Expect(rec.Code).To(Equal(http.StatusOK))
				Expect(cleanEchoResponse(rec)).To(Equal(`{"user_uuid":"fake-456","alist_uuid":"fakeList","action":"revoke"}`))
			})
		*/
	})
})
