package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StopApp stops a specific application
func (client *Client) StopApp(cluster string, namespace string, name string) error {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/%s/stop", cluster, namespace, name)
	request := client.buildStopAppRequest()

	response, err := request.Put(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()

	return err
}

func (client *Client) buildStopAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
