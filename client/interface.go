package client

// AlaudaClient is the interface implemented by the Alauda API client.
type AlaudaClient interface {
	APIServer() string
	Token() string
	SetAPIServer(string)
	SetToken(string)
	AuthClient
}

// AuthClient is the API client for authentication related APIs.
type AuthClient interface {
	Login(*LoginData) (*LoginResult, error)
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ AlaudaClient = &Client{}
