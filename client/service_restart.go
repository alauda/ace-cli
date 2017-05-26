package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RestartService restarts the specified service.
func (client *Client) RestartService(name string) error {
	url := client.buildURL("services", name)
	request := client.buildRestartServiceRequest()

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

func (client *Client) buildRestartServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
