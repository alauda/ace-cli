package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveService deletes the specified service.
func (client *Client) RemoveService(name string) error {
	url := client.buildURL("services", name)
	request := client.buildRemoveServiceRequest()

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

func (client *Client) buildRemoveServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
