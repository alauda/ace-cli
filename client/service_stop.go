package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StopService stops the specified service.
func (client *Client) StopService(name string) error {
	url := client.buildURL("services", "%s/stop", name)
	request := client.buildStopServiceRequest()

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

func (client *Client) buildStopServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
