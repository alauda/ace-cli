package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// StartService starts the specified service.
func (client *Client) StartService(name string) error {
	url := client.buildStartServiceURL(name)
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

func (client *Client) buildStartServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/services/%s/%s/start", server, client.Namespace(), name)
}

func (client *Client) buildStartServiceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
