package client

// Service defines the response body for one service returned in the ListServices API.
type Service struct {
	Name             string              `json:"service_name"`
	ImageName        string              `json:"image_name"`
	ImageTag         string              `json:"image_tag"`
	Command          string              `json:"run_command"`
	Created          string              `json:"created_at"`
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
