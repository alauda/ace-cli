package client

// App defines the response body of the InspectApp API.
type App struct {
	Name      string `json:"app_name"`
	ID        string `json:"uuid"`
	State     string `json:"current_status"`
	CreatedBy string `json:"created_by"`
}
