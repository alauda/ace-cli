package client

// APIClient is the interface implemented by the Alauda API client.
type APIClient interface {
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
	CreateService(*CreateServiceData) error
	ListServices(*ListServicesParams) (*ListServicesResult, error)
	StartService(string) error
	StopService(string) error
	RemoveService(string) error
	InspectService(string) (*Service, error)
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ APIClient = &Client{}
