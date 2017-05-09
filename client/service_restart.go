package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// RestartService restarts the specified service.
func (client *Client) RestartService(name string) error {
	url := client.buildRestartServiceURL(name)
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

func (client *Client) buildRestartServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/services/%s/%s", server, client.Namespace(), name)
}

func (client *Client) buildRestartServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
