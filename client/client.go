package client

import "net/http"

// Client is the API client for the Alauda platform.
type Client struct {
	httpClient *http.Client
	apiServer  string
}

// NewClient creates a new Alauda API client.
func NewClient(apiServer string) (*Client, error) {

	return &Client{
		httpClient: http.DefaultClient,
		apiServer:  apiServer,
	}, nil
}
