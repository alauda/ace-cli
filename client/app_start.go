package client

import (
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// StartApp starts the specified application.
func (client *Client) StartApp(name string) error {
	url := client.buildStartAppURL(name)
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

func (client *Client) buildStartAppURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/applications/%s/%s/start", server, client.Namespace(), name)
}

func (client *Client) buildStartAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
