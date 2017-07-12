package client

// Image defines the response body for one image from the ListImages API.
type Image struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"uuid"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	PulledAt    string `json:"pulled_at"`
	PushedAt    string `json:"pushed_at"`
	Uploads     int    `json:"upload"`
	Downloads   int    `json:"download"`
	IsPublic    bool   `json:"is_public"`
}
