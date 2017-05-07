package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// ScaleServiceData defines the request body for the ScaleService API.
type ScaleServiceData struct {
	TargetInstances int `json:"target_num_instances"`
}

// ScaleService scales the service to the specified number of instances
func (client *Client) ScaleService(name string, data *ScaleServiceData) error {
	url := client.buildScaleServiceURL(name)

	request, err := client.buildScaleServiceRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Put(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildScaleServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s/%s/%s", server, "services", client.Namespace(), name)
}

func (client *Client) buildScaleServiceRequest(data *ScaleServiceData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
