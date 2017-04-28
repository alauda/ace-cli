package client

// Client is the API client for the Alauda platform.
type Client struct {
	apiServer string
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

// Token field.
func (client *Client) Token() string {
	return client.token
}

// SetAPIServer sets the apiServer field.
func (client *Client) SetAPIServer(apiServer string) {
	client.apiServer = apiServer
}

// SetToken sets the token field.
func (client *Client) SetToken(token string) {
	client.token = token
}
