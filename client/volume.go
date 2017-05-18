package client

// Volume defines the response body of the InspectVolume API.
type Volume struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Driver    string `json:"driver_name"`
	State     string `json:"state"`
	Size      int    `json:"size"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}
