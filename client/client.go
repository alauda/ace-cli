package client

import (
	"fmt"
	"strings"
)

// Client is the API client for the Alauda platform.
type Client struct {
	apiServer string
	namespace string
	token     string
}

// NewClient creates a new Alauda API client.
func NewClient() (*Client, error) {
	return &Client{}, nil
}

// APIServer field.
func (client *Client) APIServer() string {
	return client.apiServer
}

// Namespace field.
func (client *Client) Namespace() string {
	return client.namespace
}

// Token field.
func (client *Client) Token() string {
	return client.token
}

// Initialize should be called before using the client.
func (client *Client) Initialize(apiServer string, namespace string, token string) {
	client.apiServer = apiServer
	client.namespace = namespace
	client.token = token
}

func (client *Client) buildURL(route string, format string, a ...interface{}) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	path := fmt.Sprintf(format, a...)

	if route != "" {
		namespacedRoute := fmt.Sprintf("%s/%s", route, client.Namespace())
		return fmt.Sprintf("%s/%s/%s", server, namespacedRoute, path)
	}

	return fmt.Sprintf("%s/%s", server, path)
}
