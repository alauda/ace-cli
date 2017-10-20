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
	VolumeClient
	AppClient
	AppTemplateClient
	ConfigClient
	RegistryClient
	ImageClient
	NodeClient
}

// AuthClient is the API client for authentication related APIs.
type AuthClient interface {
	Login(*LoginData) (*LoginResult, error)
}

// ServiceClient is the API client for service related APIs.
type ServiceClient interface {
	CreateService(*CreateServiceData) error
	ListServices(*ListServicesParams) (*ListServicesResult, error)
	StartService(string, *ServiceParams) error
	StopService(string, *ServiceParams) error
	RemoveService(string, *ServiceParams) error
	InspectService(string, *ServiceParams) (*Service, error)
	RestartService(string, *ServiceParams) error
	ScaleService(string, *ScaleServiceData, *ServiceParams) error
	UpdateService(string, *UpdateServiceData, *ServiceParams) error
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

// VolumeClient is the API client for volume related APIs.
type VolumeClient interface {
	ListVolumes(*ListVolumesParams) (*ListVolumesResult, error)
	InspectVolume(string) (*Volume, error)
	CreateVolume(*CreateVolumeData) error
	RemoveVolume(string) error
}

// AppClient is the API client for the app related APIs.
type AppClient interface {
	CreateApp(*CreateAppData, string) error
	ListApps(*ListAppsParams) (*ListAppsResult, error)
	InspectApp(string) (*App, error)
	StartApp(string) error
	StopApp(string) error
	RemoveApp(string) error
}

// AppTemplateClient is the API client for the app template related APIs.
type AppTemplateClient interface {
	CreateAppTemplate(*CreateAppTemplateData, string) error
	ListAppTemplates() (*ListAppTemplatesResult, error)
	InspectAppTemplate(string) (*AppTemplate, error)
	DownloadAppTemplate(string, string) error
	UpdateAppTemplate(string, *UpdateAppTemplateData, string) error
	RemoveAppTemplate(string) error
}

// ConfigClient is the API client for the config related APIs.
type ConfigClient interface {
	CreateConfig(*CreateConfigData) error
	ListConfigs() (*ListConfigsResult, error)
	InspectConfig(string) (*Config, error)
	RemoveConfig(string) error
}

// RegistryClient is the API client for the registry related APIs.
type RegistryClient interface {
	ListRegistries() (*ListRegistriesResult, error)
	ListRegistryProjects(string) (*ListRegistryProjectsResult, error)
}

// ImageClient is the API client for the image related APIs.
type ImageClient interface {
	ListImages(string, string) (*ListImagesResult, error)
	ListImageTags(string, string, string) (*ListImageTagsResult, error)
}

// NodeClient is the API client for node related APIs.
type NodeClient interface {
	ListNodes(string) (*ListNodesResult, error)
	InspectNode(string, string) (*Node, error)
	CordonNode(string, string) error
	UncordonNode(string, string) error
	DrainNode(string, string) error
	SetNodeLabels(string, string, *SetNodeLabelsData) error
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ APIClient = &Client{}
