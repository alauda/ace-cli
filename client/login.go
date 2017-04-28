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

// LoginResponse defines the response body for the login API.
type LoginResult struct {
	Namespace string `json:"namespace"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

// Login authenticates against the Alauda server.
func (client *Client) Login(data *LoginData) (*LoginResult, error) {
	url := buildRequestURL(client.APIServer())

	body, err := buildRequestBody(data)
	if err != nil {
		return nil, err
	}

	request := rest.NewRequest()

	response, err := request.Post(url, body)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	return buildResult(response)
}

func buildRequestURL(apiServer string) string {
	server := strings.TrimSuffix(apiServer, "/")
	return fmt.Sprintf("%s/%s", server, "generate-api-token")
}

func buildRequestBody(data *LoginData) ([]byte, error) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}

func buildResult(response *rest.Response) (*LoginResult, error) {
	result := LoginResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
