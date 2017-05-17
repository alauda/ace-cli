package client

// Service defines the response body for one service returned in the ListServices API.
type Service struct {
	Name             string              `json:"service_name"`
	ImageName        string              `json:"image_name"`
	ImageTag         string              `json:"image_tag"`
	Command          string              `json:"run_command"`
	Entrypoint       string              `json:"entrypoint"`
	CreatedAt        string              `json:"created_at"`
	Size             ServiceInstanceSize `json:"custom_instance_size"`
	Ports            []int               `json:"ports"`
	TargetInstances  int                 `json:"target_num_instances"`
	HealthyInstances int                 `json:"healthy_num_instances"`
	Status           string              `json:"current_status"`
	NetworkMode      string              `json:"network_mode"`
	Env              map[string]string   `json:"instance_envvars"`
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
