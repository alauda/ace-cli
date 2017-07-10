package client

// Service defines the response body for one service returned in the ListServices API.
type Service struct {
	Name             string              `json:"service_name"`
	ID               string              `json:"uuid"`
	ImageName        string              `json:"image_name"`
	ImageTag         string              `json:"image_tag"`
	Command          string              `json:"run_command"`
	Entrypoint       string              `json:"entrypoint"`
	CreatedAt        string              `json:"created_at"`
	Size             ServiceInstanceSize `json:"custom_instance_size"`
	Ports            []int               `json:"ports"`
	TargetInstances  int                 `json:"target_num_instances"`
	HealthyInstances int                 `json:"healthy_num_instances"`
	State            string              `json:"current_status"`
	NetworkMode      string              `json:"network_mode"`
	Env              map[string]string   `json:"instance_envvars"`
	Volumes          []ServiceVolume     `json:"volumes"`
}

// ServiceParams defines the base query parameters for various service APIs.
type ServiceParams struct {
	App string
}

// ServiceInstanceSize defines the size of the service instances.
type ServiceInstanceSize struct {
	Memory int     `json:"mem"`
	CPU    float64 `json:"cpu"`
}

// ServiceLoadBalancer defines the load balancer data in the CreateService request.
type ServiceLoadBalancer struct {
	ID        string            `json:"load_balancer_id"`
	Type      string            `json:"type"`
	Listeners []ServiceListener `json:"listeners"`
}

// ServiceListener defines the load balancer listener data in the CreateService request.
type ServiceListener struct {
	ListenerPort  int    `json:"listener_port"`
	ContainerPort int    `json:"container_port"`
	Protocol      string `json:"protocol"`
}

// ServiceVolume defines the volume data in the CreateService request.
type ServiceVolume struct {
	Path       string `json:"app_volume_dir"`
	VolumeName string `json:"volume_name"`
	VolumeID   string `json:"volume_id"`
}

// ServiceConfig defines the configuration data in the CreateService request.
type ServiceConfig struct {
	Type  string             `json:"type"`
	Path  string             `json:"path"`
	Value ServiceConfigValue `json:"value"`
}

// ServiceConfigValue defines the configuration value to be injected into the service.
type ServiceConfigValue struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
