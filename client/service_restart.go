package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RestartServiceParams defines the query parameters for the RestartService API.
type RestartServiceParams struct {
	App string
}

// RestartService restarts the specified service.
func (client *Client) RestartService(name string, params *RestartServiceParams) error {
	url := client.buildURL("services", name)
	request := client.buildRestartServiceRequest(params)

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

func (client *Client) buildRestartServiceRequest(params *RestartServiceParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.App != "" {
		request.SetQueryParam("application", params.App)
	}

	return request
}
