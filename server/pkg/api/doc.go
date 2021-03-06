package api

import "github.com/freshteapot/learnalist-api/server/api/i18n"

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}

type HTTPUserRegisterResponse struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
}

type HTTPResponseMessage struct {
	Message string `json:"message"`
}

type HTTPLabelInput struct {
	Label string `json:"label"`
}

type HTTPGetVersionResponse struct {
	GitHash string `json:"gitHash"`
	GitDate string `json:"gitDate"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type HTTPShareListInput struct {
	AlistUUID string `json:"alist_uuid"`
	Action    string `json:"action"`
}

type HTTPShareListWithUserInput struct {
	UserUUID  string `json:"user_uuid"`
	AlistUUID string `json:"alist_uuid"`
	Action    string `json:"action"`
}

type HTTPLogoutRequest struct {
	Kind     string `json:"kind"`
	UserUUID string `json:"user_uuid"`
	Token    string `json:"token"`
}

var (
	HTTPErrorResponse = HTTPResponseMessage{
		Message: i18n.InternalServerErrorFunny,
	}

	HTTPAccessDeniedResponse = HTTPResponseMessage{
		Message: i18n.AclHttpAccessDeny,
	}
)
