package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveService deletes the specified service.
func (client *Client) RemoveService(name string, params *ServiceParams) error {
	url := client.buildURL("services", name)
	request := client.buildRemoveServiceRequest(params)

	response, err := request.Delete(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildRemoveServiceRequest(params *ServiceParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.App != "" {
		request.SetQueryParam("application", params.App)
	}

	return request
}
