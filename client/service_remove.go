package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// RemoveService starts the specified service.
func (client *Client) RemoveService(name string) error {
	url := client.buildRemoveServiceURL(name)
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

func (client *Client) buildRemoveServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s/%s/%s", server, "services", client.Namespace(), name)
}

func (client *Client) buildRemoveServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
