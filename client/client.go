package client

import "net/http"

// Client is the API client for the Alauda platform.
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new Alauda API client.
func NewClient() (*Client, error) {

	return &Client{
		httpClient: http.DefaultClient,
	}, nil
}
