package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// LoginData defines the request body for the login API.
type LoginData struct {
	Organization string `json:"organization"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// LoginResult defines the response body for the login API.
type LoginResult struct {
	Namespace string `json:"namespace"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

// Login authenticates against the Alauda server.
func (client *Client) Login(data *LoginData) (*LoginResult, error) {
	url := client.buildLoginURL()

	request, err := client.buildLoginRequest(data)
	if err != nil {
		return nil, err
	}

	response, err := request.Post(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	return parseLoginResult(response)
}

func (client *Client) buildLoginURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s", server, "generate-api-token")
}

func (client *Client) buildLoginRequest(data *LoginData) (*rest.Request, error) {
	request := rest.NewRequest("")

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}

func parseLoginResult(response *rest.Response) (*LoginResult, error) {
	result := LoginResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
