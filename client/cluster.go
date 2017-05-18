package client

// Cluster defines the response body for one cluster returned by the ListClusters API.
type Cluster struct {
	Name        string            `json:"name"`
	DisplayName string            `json:"display_name"`
	ID          string            `json:"id"`
	Type        string            `json:"container_manager"`
	State       string            `json:"state"`
	CreatedAt   string            `json:"created_at"`
	Attributes  ClusterAttributes `json:"attr"`
}

// ClusterAttributes contains the attributes of a cluster.
type ClusterAttributes struct {
	Cloud ClusterCloudInfo `json:"cloud"`
}

// ClusterCloudInfo contains information about the cloud/region the cluster is deployed in.
type ClusterCloudInfo struct {
	Name   string `json:"name"`
	Region string `json:"region_id"`
}
