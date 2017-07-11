package client

// Registry defines the response body for the InspectRegistry API.
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
