package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ScaleServiceData defines the request body for the ScaleService API.
type ScaleServiceData struct {
	TargetInstances int `json:"target_num_instances"`
}

// ScaleService scales the service to the specified number of instances
func (client *Client) ScaleService(name string, data *ScaleServiceData, params *ServiceParams) error {
	url := client.buildURL("services", name)

	request, err := client.buildScaleServiceRequest(data, params)
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

func (client *Client) buildScaleServiceRequest(data *ScaleServiceData, params *ServiceParams) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	if params.App != "" {
		request.SetQueryParam("application", params.App)
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
