package client

// Stack defines the response body of the InspectStack API.
type Stack struct {
	Name      string `json:"app_name"`
	ID        string `json:"uuid"`
	State     string `json:"current_status"`
	CreatedBy string `json:"created_by"`
}
