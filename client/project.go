package client

// Project defines the response body for one project returned by the ListProjects API.
type Project struct {
	ID        string `json:"uuid"`
	Name      string `json:"name"`
	State     string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
