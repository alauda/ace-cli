package client

// Registry defines the response body for one registry from the ListRegistries API.
type Registry struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	ID          string `json:"uuid"`
	Endpoint    string `json:"endpoint"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	Audience    string `json:"audience"`
	IsPublic    bool   `json:"is_public"`
	Cluster     string `json:"region_id"`
	Issuer      string `json:"issuer"`
}

// RegistryProject defines the response body for one registry project from the ListRegistryProjects API.
type RegistryProject struct {
	Name      string `json:"project_name"`
	ID        string `json:"project_id"`
	CreatedBy string `json:"created_by"`
	RepoCount int    `json:"repo_count"`
}
