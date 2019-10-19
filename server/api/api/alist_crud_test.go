package api_test

//2, as I am being really lazy :(, once all moved over to ginkgo remove.
import (
	"encoding/json"
	"errors"
	"net/http"

	mockHugo "github.com/freshteapot/learnalist-api/server/alists/pkg/hugo/mocks"
	"github.com/freshteapot/learnalist-api/server/api/alist"
	"github.com/freshteapot/learnalist-api/server/api/i18n"
	mockModels "github.com/freshteapot/learnalist-api/server/api/models/mocks"
	"github.com/freshteapot/learnalist-api/server/api/uuid"
	mockAcl "github.com/freshteapot/learnalist-api/server/pkg/acl/mocks"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Testing Alist endpoints", func() {
	AfterEach(emptyDatabase)

	When("Basic crud", func() {
		var userUUID string
		var datastore *mockModels.Datastore
		var acl *mockAcl.Acl
		var user *uuid.User
		var method string
		var uri string
		BeforeEach(func() {
			method = http.MethodPost
			uri = "/api/v1/alist"
			testHugoHelper := &mockHugo.HugoSiteBuilder{}
			testHugoHelper.On("Write", mock.Anything)
			testHugoHelper.On("Remove", mock.Anything)
			m.HugoHelper = testHugoHelper

			datastore = &mockModels.Datastore{}
			acl = &mockAcl.Acl{}
			m.Datastore = datastore
			m.Acl = acl

			user = &uuid.User{
				Uuid: "fake-123",
			}
		})
		Context("Save a list", func() {
			It("Reject if wrong method", func() {
				method = "DELETE"
				input := ""
				req, rec := setupFakeEndpoint(method, uri, input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"This method is not supported."}`))
			})

			It("Get, is not accepted", func() {
				//datastore.On("GetUserByCredentials", mock.Anything).Return(user, nil)
				input := ""
				req, rec := setupFakeEndpoint(method, uri, input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"Your input is invalid json."}`))
			})

			It("Post, success", func() {
				savedList := alist.NewTypeV1()
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(savedList, nil)
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	}
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}

				req, rec := setupFakeEndpoint(method, uri, input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusCreated))
				b, _ := json.Marshal(savedList)

				Expect(cleanEchoJSONResponse(rec)).To(Equal(string(b)))
			})

			It("Post, fail, due to ownership", func() {
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(nil, errors.New(i18n.InputSaveAlistOperationOwnerOnly))
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	}
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}

				req, rec := setupFakeEndpoint(method, uri, input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusForbidden))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"Only the owner of the list can modify it."}`))
			})

			It("PUT, fail, due to list uuid not being found", func() {
				method := http.MethodPut
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(nil, errors.New(i18n.SuccessAlistNotFound))
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	}
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}
				uri = uri + "/1234"
				req, rec := setupFakeEndpoint(method, "/", input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.SetPath("/api/v1/alist/:uuid")
				c.Set("loggedInUser", *user)
				c.SetParamNames("uuid")
				c.SetParamValues("")
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"Please refer to the documentation on lists"}`))
			})

			It("PUT, fail, due to list uuid not being found", func() {
				method := http.MethodPut
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(nil, errors.New(i18n.SuccessAlistNotFound))
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	}
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}

				req, rec := setupFakeEndpoint(method, "/", input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				c.SetPath("/alist/:uuid")
				c.Set("loggedInUser", *user)
				c.SetParamNames("uuid")
				c.SetParamValues("1234")
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusNotFound))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"List not found."}`))
			})

			It("PUT, fail, due to uuid in uri not matching in the list", func() {
				method := http.MethodPut
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(nil, errors.New(i18n.SuccessAlistNotFound))
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	},
				"uuid": "fake-456"
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}

				req, rec := setupFakeEndpoint(method, "/", input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				c.SetPath("/alist/:uuid")
				c.Set("loggedInUser", *user)
				c.SetParamNames("uuid")
				c.SetParamValues("1234")
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"The list uuid in the uri doesnt match that in the payload"}`))
			})

			It("Post, fail, due to internal issues", func() {
				datastore.On("SaveAlist", mock.Anything, mock.Anything).Return(nil, errors.New("Failed"))
				input := `
      {
      	"data": ["car"],
      	"info": {
      		"title": "Days of the Week",
      		"type": "v1"
      	}
      }
      `
				user := &uuid.User{
					Uuid: userUUID,
				}

				req, rec := setupFakeEndpoint(method, uri, input)
				e := echo.New()
				c := e.NewContext(req, rec)
				c.Set("loggedInUser", *user)
				m.V1SaveAlist(c)
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
				Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"Failed"}`))
			})
		})

	})

})
