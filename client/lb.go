package client

// LoadBalancer defines the response body of the InspectLoadBalancer API.
type LoadBalancer struct {
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Address     string     `json:"address"`
	AddressType string     `json:"address_type"`
	CreatedWith string     `json:"create_type"`
	CreatedAt   string     `json:"created_at"`
	Listeners   []Listener `json:"listeners"`
}

// Listener defines the response body for one listener from the InspectLoadBalancer API.
type Listener struct {
	ServiceID     string `json:"service_id"`
	ServiceName   string `json:"service_name"`
	Protocol      string `json:"protocol"`
	ListenerPort  int    `json:"listener_port"`
	ContainerPort int    `json:"container_port"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
