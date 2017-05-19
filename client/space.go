package client

// Space defines the response body for one space returned by the ListSpaces API.
type Space struct {
	ID          string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
}
