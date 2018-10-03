package client

import (
	"github.com/alauda/alauda/client/rest"
)

// GetAppYaml retrieves the YAML of a specific application
func (client *Client) GetAppYaml(cluster string, namespace string, name string) (string, error) {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/applications/%s/%s/yaml", cluster, namespace, name)
	request := client.buildGetAppYamlRequest()

	response, err := request.Get(url)
	if err != nil {
		return "", err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return "", err
	}

	return string(response.Body()), nil
}

func (client *Client) buildGetAppYamlRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}
