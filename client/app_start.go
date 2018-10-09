package client

import (
	"github.com/alauda/alauda/client/rest"
)

// StartApp starts a specific application
func (client *Client) StartApp(cluster string, namespace string, name string) error {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/%s/start", cluster, namespace, name)
	request := client.buildStartAppRequest()

	response, err := request.Put(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()

	return err
}

func (client *Client) buildStartAppRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
