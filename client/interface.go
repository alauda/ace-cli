package client

// APIClient is the interface implemented by the Alauda API client.
type APIClient interface {
	APIServer() string
	Namespace() string
	Token() string
	Initialize(string, string, string)
	AuthClient
	ServiceClient
	SpaceClient
	ClusterClient
	LoadBalancerClient
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
	RestartService(string) error
	ScaleService(string, *ScaleServiceData) error
	UpdateService(string, *UpdateServiceData) error
}

// SpaceClient is the API client for space related APIs.
type SpaceClient interface {
	ListSpaces() (*ListSpacesResult, error)
	InspectSpace(string) (*Space, error)
}

// ClusterClient is the API client for cluster related APIs.
type ClusterClient interface {
	ListClusters() (*ListClustersResult, error)
	InspectCluster(string) (*Cluster, error)
}

// LoadBalancerClient is the API client for LB related APIs.
type LoadBalancerClient interface {
	ListLoadBalancers(*ListLoadBalancersParams) (*ListLoadBalancersResult, error)
	InspectLoadBalancer(string) (*LoadBalancer, error)
	UpdateLoadBalancer(string, *UpdateLoadBalancerData) error
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ APIClient = &Client{}
