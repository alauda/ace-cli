package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// CreateVolumeData defines the request body for the CreateVolume API.
type CreateVolumeData struct {
	Name       string `json:"name"`
	Driver     string `json:"driver_name"`
	Size       int    `json:"size"`
	ClusterID  string `json:"region_id"`
	VolumeType string `json:"volume_type"`
	Namespace  string `json:"namespace"`
}

// CreateVolume creates a new volume.
func (client *Client) CreateVolume(data *CreateVolumeData) error {
	url := client.buildCreateVolumeURL()

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

func (client *Client) buildCreateVolumeURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/storage/%s/volumes", server, client.Namespace())
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
