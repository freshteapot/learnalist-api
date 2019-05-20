package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/freshteapot/learnalist-api/api/uuid"
	"github.com/labstack/echo/v4"
)

type HttpShareListWithUserInput struct {
	UserUUID  string `json:"user_uuid"`
	AlistUUID string `json:"alist_uuid"`
	Action    string `json:"action"`
}

func (env *Env) V1ShareAlist(c echo.Context) error {
	user := c.Get("loggedInUser").(uuid.User)
	// TODO maybe we support an array
	var input = &HttpShareListWithUserInput{}

	defer c.Request().Body.Close()
	jsonBytes, _ := ioutil.ReadAll(c.Request().Body)

	err := json.Unmarshal(jsonBytes, input)
	if err != nil {
		// TODO
		response := HttpResponseMessage{
			Message: "Wrong format",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	aList, _ := env.Datastore.GetAlist(input.AlistUUID)
	if aList == nil {
		response := HttpResponseMessage{
			Message: "Wrong format",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if aList.User.Uuid == user.Uuid {
		if input.UserUUID != user.Uuid {
			if input.Action == "grant" {
				env.Acl.GrantListReadAccess(input.UserUUID, input.AlistUUID)
			}
			if input.Action == "revoke" {
				env.Acl.RevokeListReadAccess(input.UserUUID, input.AlistUUID)
			}
		}
	}
	return c.JSON(http.StatusOK, input)
}
