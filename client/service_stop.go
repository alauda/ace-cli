package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StopService stops the specified service.
func (client *Client) StopService(name string, params *ServiceParams) error {
	url := client.buildURL("services", "%s/stop", name)
	request := client.buildStopServiceRequest(params)

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

func (client *Client) buildStopServiceRequest(params *ServiceParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.App != "" {
		request.SetQueryParam("application", params.App)
	}

	return request
}
