package client

import (
	"github.com/alauda/alauda/client/rest"
)

// DeleteApp deletes a specific application
func (client *Client) DeleteApp(cluster string, namespace string, name string) error {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/%s", cluster, namespace, name)
	request := client.buildDeleteAppRequest()

	response, err := request.Delete(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()

	return err
}

func (client *Client) buildDeleteAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
