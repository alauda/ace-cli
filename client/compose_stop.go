package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StopApp stops the specified application.
func (client *Client) StopApp(name string) error {
	url := client.buildURL("applications", "%s/stop", name)
	request := client.buildStopAppRequest()

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

func (client *Client) buildStopAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
