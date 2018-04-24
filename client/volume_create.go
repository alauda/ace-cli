package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// CreateVolumeData defines the request body for the CreateVolume API.
type CreateVolumeData struct {
	Name       string `json:"name"`
	Driver     string `json:"driver_name"`
	Size       int    `json:"size"`
	ClusterID  string `json:"region_id"`
	Space      string `json:"space_name"`
	VolumeType string `json:"volume_type"`
	Namespace  string `json:"namespace"`
}

// CreateVolume creates a new volume.
func (client *Client) CreateVolume(data *CreateVolumeData) error {
	url := client.buildURL("v1", "storage", "volumes")

	request, err := client.buildCreateVolumeRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildCreateVolumeRequest(data *CreateVolumeData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	data.Namespace = client.Namespace()

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
