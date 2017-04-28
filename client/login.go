package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// LoginOptions defines the input options for the login API.
type LoginOptions struct {
	Server   string
	Account  string
	Username string
	Password string
}

type loginRequestBody struct {
	Organization string `json:"organization"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

// LoginSuccess contains the token after a successful login.
type LoginSuccess struct {
	Token string
}

type loginResponseBody struct {
	Namespace string
	Username  string
	Email     string
	Token     string
}

// Login authenticates against the Alauda server.
func (client *Client) Login(opts *LoginOptions) (*LoginSuccess, error) {
	url := buildURL(opts)

	body, err := buildBody(opts)
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

	result := loginResponseBody{}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &LoginSuccess{
		Token: result.Token,
	}, nil
}

func buildURL(opts *LoginOptions) string {
	server := strings.TrimSuffix(opts.Server, "/")
	return fmt.Sprintf("%s/%s", server, "generate-api-token")
}

func buildBody(opts *LoginOptions) ([]byte, error) {
	body := &loginRequestBody{
		Organization: opts.Account,
		Username:     opts.Username,
		Password:     opts.Password,
	}

	marshalled, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}
