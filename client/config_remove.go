package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveConfig deletes the specified configuration.
func (client *Client) RemoveConfig(name string) error {
	url := client.buildURL("configs", name)
	request := client.buildRemoveConfigRequest()

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

func (client *Client) buildRemoveConfigRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
