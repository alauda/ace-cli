package client

// Node defines the response body for one node returned by the ListNodes API.
type Node struct {
	IP         string         `json:"private_ip"`
	State      string         `json:"state"`
	Type       string         `json:"type"`
	Attributes NodeAttributes `json:"attr"`
	Resources  NodeResources  `json:"resources"`
	Labels     []NodeLabel    `json:"labels"`
}

// NodeResources describes the total and available resources on a node.
type NodeResources struct {
	AvailableMemory string `json:"available_mem"`
	AvailableCPUs   string `json:"avaliable_cpus"`
	TotalMemory     string `json:"total_mem"`
	TotalCPUs       string `json:"total_cpus"`
}

// NodeAttributes describes the additional attributes of the node.
type NodeAttributes struct {
	Schedulable bool `json:"schedulable"`
}

// NodeLabel describes a label on the node.
type NodeLabel struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Editable bool   `json:"editable"`
}
