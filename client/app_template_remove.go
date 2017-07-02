package client

import (
	"github.com/alauda/alauda/client/rest"
)

// RemoveAppTemplate deletes the specified app template.
func (client *Client) RemoveAppTemplate(name string) error {
	url := client.buildURL("application-templates", name)
	request := client.buildRemoveAppTemplatesRequest()

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

func (client *Client) buildRemoveAppTemplatesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
