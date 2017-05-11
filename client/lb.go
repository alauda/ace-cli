package client

// LoadBalancer defines the response body of the InspectLoadBalancer API.
type LoadBalancer struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	AddressType string `json:"address_type"`
	CreatedWith string `json:"create_type"`
	CreatedAt   string `json:"created_at"`
}
