package client

// Node defines the response body for one node returned by the ListNodes API.
type Node struct {
	IP        string        `json:"private_ip"`
	State     string        `json:"state"`
	Type      string        `json:"type"`
	Resources NodeResources `json:"resources"`
}

// NodeResources describes the total and available resources on a node.
type NodeResources struct {
	AvailableMemory string `json:"available_mem"`
	AvailableCPUs   string `json:"avaliable_cpus"`
	TotalMemory     string `json:"total_mem"`
	TotalCPUs       string `json:"total_cpus"`
}
