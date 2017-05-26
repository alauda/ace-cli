package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StartService starts the specified service.
func (client *Client) StartService(name string) error {
	url := client.buildURL("services", "%s/start", name)
	request := client.buildStartServiceRequest()

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

func (client *Client) buildStartServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
