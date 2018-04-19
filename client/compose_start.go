package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StartApp starts the specified application.
func (client *Client) StartApp(name string) error {
	url := client.buildURL("applications", "%s/start", name)
	request := client.buildStartAppRequest()

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

func (client *Client) buildStartAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
