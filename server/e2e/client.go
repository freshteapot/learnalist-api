package e2e

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/freshteapot/learnalist-api/server/api/alist"
	"github.com/freshteapot/learnalist-api/server/api/api"
)

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

type RegisterResponse struct {
	Username  string `json:"username"`
	Uuid      string `json:"uuid"`
	BasicAuth string
}

type AlistUuidResponse struct {
	Uuid string `json:"uuid"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
}

type Client struct {
	server     string
	httpClient *http.Client
}

func NewClient(_server string) Client {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 1 * time.Millisecond,
		}).Dial,
		TLSHandshakeTimeout: 1 * time.Millisecond,
	}
	var netClient = &http.Client{
		Timeout:   time.Millisecond * 5,
		Transport: netTransport,
	}

	return Client{
		server:     _server,
		httpClient: netClient,
	}
}

func (c Client) getServerURL() string {
	return c.server
}

func getBasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c Client) Register(username string, password string) RegisterResponse {
	fmt.Println("Registering user via Register")
	body := strings.NewReader(fmt.Sprintf(`
{
    "username":"%s",
    "password":"%s"
}
`, username, password))

	url := fmt.Sprintf("%s/api/v1/register", c.getServerURL())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
		fmt.Println("Failed NewRequest")
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// handle err
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle err
		panic(err)
	}
	var response RegisterResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		// handle err
		panic(err)
	}
	response.BasicAuth = getBasicAuth(username, password)
	return response
}

func (c Client) RawPostListV1(userInfo RegisterResponse, input string) (*http.Response, error) {
	var response *http.Response
	body := strings.NewReader(input)
	url := fmt.Sprintf("%s/api/v1/alist", c.getServerURL())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c Client) PostListV1(userInfo RegisterResponse, input string) (alist.Alist, error) {
	fmt.Println("Posting a list via PostListV1")
	var response alist.Alist
	resp, err := c.RawPostListV1(userInfo, input)
	if err != nil {
		// handle err
		return response, nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response)
	return response, nil
}

func (c Client) RawPutListV1(userInfo RegisterResponse, uuid string, input string) (*http.Response, error) {
	fmt.Println("Updating a list via RawPutListV1")
	var response *http.Response
	body := strings.NewReader(input)
	url := fmt.Sprintf("%s/api/v1/alist/%s", c.getServerURL(), uuid)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c Client) PutListV1(userInfo RegisterResponse, uuid string, input string) (alist.Alist, error) {
	fmt.Println("Updating a list via PutListV1")
	var response alist.Alist
	resp, err := c.RawPutListV1(userInfo, uuid, input)

	if err != nil {
		// handle err
		return response, nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response)
	return response, nil
}

func (c Client) SetListShare(userInfo RegisterResponse, alistUUID string, action string) MessageResponse {
	body := strings.NewReader(fmt.Sprintf(`{
  "alist_uuid": "%s",
  "action": "%s"
}`, alistUUID, action))
	url := fmt.Sprintf("%s/api/v1/share/alist", c.getServerURL())

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	var response MessageResponse
	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response)
	return response
}

func (c Client) GetListByUUID(userInfo RegisterResponse, uuid string) HttpResponse {
	url := fmt.Sprintf("%s/api/v1/alist/%s", c.getServerURL(), uuid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// handle err
	}

	defer resp.Body.Close()

	var response HttpResponse
	response.StatusCode = resp.StatusCode
	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response.Body)
	return response
}

func (c Client) RawPostLabelV1(userInfo RegisterResponse, label string) (*http.Response, error) {
	fmt.Println("Posting a list via RawPostLabelV1")
	input := api.HttpLabelInput{
		Label: label,
	}
	var response *http.Response
	b, _ := json.Marshal(input)
	body := strings.NewReader(string(b))
	url := fmt.Sprintf("%s/api/v1/labels", c.getServerURL())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c Client) PostLabelV1(userInfo RegisterResponse, label string) ([]string, error) {
	fmt.Println("Posting a list via PostLabelV1")
	var response []string

	resp, err := c.RawPostLabelV1(userInfo, label)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c Client) RawGetLabelsByMeV1(userInfo RegisterResponse) (*http.Response, error) {
	var response *http.Response
	fmt.Println("GET  labels via RawGetLabelsByMeV1")
	url := fmt.Sprintf("%s/api/v1/labels/by/me", c.getServerURL())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c Client) GetLabelsByMeV1(userInfo RegisterResponse) ([]string, error) {
	fmt.Println("GET  labels via GetLabelsByMeV1")
	var response []string

	resp, err := c.RawGetLabelsByMeV1(userInfo)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c Client) RawDeleteLabelV1(userInfo RegisterResponse, label string) (*http.Response, error) {
	fmt.Println("Posting a list via RawDeleteLabelV1")
	var response *http.Response
	url := fmt.Sprintf("%s/api/v1/labels/%s", c.getServerURL(), label)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}

func (c Client) RawGetListsByMe(userInfo RegisterResponse) (*http.Response, error) {
	var response *http.Response
	url := fmt.Sprintf("%s/api/v1/alist/by/me", c.getServerURL())
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// handle err
		return response, nil
	}
	req.Header.Set("Authorization", "Basic "+userInfo.BasicAuth)
	req.Header.Set("Content-Type", "application/json")

	return c.httpClient.Do(req)
}
