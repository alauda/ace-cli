package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListNamespacesResult defines the response body for the ListNamespaces API.
type ListNamespacesResult struct {
	Namespaces []Namespace
}

// ListNamespaces returns all namespaces in a cluster.
func (client *Client) ListNamespaces(cluster string) (*ListNamespacesResult, error) {
	url := client.buildURL("v2", "kubernetes", "clusters/%s/namespaces/", cluster)
	request := client.buildListNamespacesRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListNamespacesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListNamespacesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListNamespacesResult(response *rest.Response) (*ListNamespacesResult, error) {
	result := ListNamespacesResult{}

	err := json.Unmarshal(response.Body(), &result.Namespaces)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
