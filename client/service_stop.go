package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// StopService stops the specified service.
func (client *Client) StopService(name string) error {
	url := client.buildStopServiceURL(name)
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

func (client *Client) buildStopServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/services/%s/%s/stop", server, client.Namespace(), name)
}

func (client *Client) buildStopServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
