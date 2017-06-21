package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveApp deletes the specified app.
func (client *Client) RemoveApp(name string) error {
	url := client.buildURL("applications", name)
	request := client.buildRemoveAppRequest()

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

func (client *Client) buildRemoveAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
