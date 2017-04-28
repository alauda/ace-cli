package client

// APIClient is the interface implemented by the Alauda API client.
type APIClient interface {
	AuthAPIClient
}

// AuthAPIClient is the API client for authentication related APIs.
type AuthAPIClient interface {
	Login(opts *LoginOptions) (*LoginSuccess, error)
}

// Type checking to ensure Client correctly implements APIClient.
var _ APIClient = &Client{}
