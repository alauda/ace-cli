package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StartService starts the specified service.
func (client *Client) StartService(name string, params *ServiceParams) error {
	url := client.buildURL("services", "%s/start", name)
	request := client.buildStartServiceRequest(params)

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

func (client *Client) buildStartServiceRequest(params *ServiceParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.App != "" {
		request.SetQueryParam("application", params.App)
	}

	return request
}
