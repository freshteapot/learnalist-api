package api_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	mockHugo "github.com/freshteapot/learnalist-api/server/alists/pkg/hugo/mocks"
	"github.com/freshteapot/learnalist-api/server/api/i18n"
	mockModels "github.com/freshteapot/learnalist-api/server/api/models/mocks"
	"github.com/freshteapot/learnalist-api/server/api/uuid"
	mockAcl "github.com/freshteapot/learnalist-api/server/pkg/acl/mocks"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Testing Api endpoints that get lists", func() {
	AfterEach(emptyDatabase)

	var datastore *mockModels.Datastore
	var acl *mockAcl.Acl
	var user *uuid.User
	var (
		method    string
		uri       string
		input     string
		alistUUID string
	)

	var c echo.Context
	var req *http.Request
	var rec *httptest.ResponseRecorder
	BeforeEach(func() {
		testHugoHelper := &mockHugo.HugoSiteBuilder{}
		testHugoHelper.On("Write", mock.Anything)
		testHugoHelper.On("Remove", mock.Anything)
		m.HugoHelper = testHugoHelper

		datastore = &mockModels.Datastore{}
		acl = &mockAcl.Acl{}
		m.Datastore = datastore
		m.Acl = acl

		alistUUID = "fake-list-123"
		method = http.MethodDelete
		input = ""
		user = &uuid.User{
			Uuid: "fake-123",
		}

		uri = fmt.Sprintf("/api/v1/alist/%s", alistUUID)
		req, rec = setupFakeEndpoint(method, uri, input)
		e := echo.New()
		c = e.NewContext(req, rec)
		c.SetPath("/api/v1/alist/:uuid")
		c.Set("loggedInUser", *user)
		c.SetParamNames("uuid")
		c.SetParamValues(alistUUID)
	})

	When("Remove a list", func() {
		It("List being removed is not found", func() {
			datastore.On("RemoveAlist", alistUUID, user.Uuid).Return(errors.New(i18n.SuccessAlistNotFound))
			m.V1RemoveAlist(c)
			Expect(rec.Code).To(Equal(http.StatusNotFound))
			Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"List not found."}`))
		})

		It("Only the owner of the list can remove it", func() {
			datastore.On("RemoveAlist", alistUUID, user.Uuid).Return(errors.New(i18n.InputDeleteAlistOperationOwnerOnly))
			m.V1RemoveAlist(c)
			Expect(rec.Code).To(Equal(http.StatusForbidden))
			Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"Only the owner of the list can remove it."}`))
		})

		It("An error occurred whilst trying to remove the list", func() {
			datastore.On("RemoveAlist", alistUUID, user.Uuid).Return(errors.New("Fail"))
			m.V1RemoveAlist(c)
			Expect(rec.Code).To(Equal(http.StatusInternalServerError))
			Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"We have failed to remove your list."}`))
		})

		It("Successfully removed a list", func() {
			datastore.On("RemoveAlist", alistUUID, user.Uuid).Return(nil)
			m.V1RemoveAlist(c)
			Expect(rec.Code).To(Equal(http.StatusOK))
			Expect(cleanEchoJSONResponse(rec)).To(Equal(`{"message":"List fake-list-123 was removed."}`))
		})
	})
})
