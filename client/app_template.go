package client

// AppTemplate defines the response body of the InspectAppTemplate API.
type AppTemplate struct {
	Name        string    `json:"app_name"`
	ID          string    `json:"uuid"`
	Description string    `json:"description"`
	Template    string    `json:"template"`
	CreatedAt   string    `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   string    `json:"updated_at"`
	Services    []Service `json:"services"`
}
