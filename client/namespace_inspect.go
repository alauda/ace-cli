package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectNamespace retrieves details about the specified namespace.
func (client *Client) InspectNamespace(cluster string, name string) (*Namespace, error) {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/namespaces/%s", cluster, name)
	request := client.buildInspectNamespaceRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectNamespaceResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectNamespaceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectNamespaceResult(response *rest.Response) (*Namespace, error) {
	result := Namespace{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
