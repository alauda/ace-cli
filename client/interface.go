package client

// AlaudaClient is the interface implemented by the Alauda API client.
type AlaudaClient interface {
	APIServer() string
	Namespace() string
	Token() string
	Initialize(string, string, string)
	AuthClient
	ServiceClient
}

// AuthClient is the API client for authentication related APIs.
type AuthClient interface {
	Login(*LoginData) (*LoginResult, error)
}

// ServiceClient is the API client for service related APIs.
type ServiceClient interface {
	ListServices(*ListServicesParams) (*ListServicesResult, error)
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ AlaudaClient = &Client{}
