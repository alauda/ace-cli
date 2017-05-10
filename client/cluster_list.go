package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// ListClustersResult defines the response body for the ListClusters API.
type ListClustersResult struct {
	Clusters []Cluster
}

// ListClusters returns all clusters in an account.
func (client *Client) ListClusters() (*ListClustersResult, error) {
	url := client.buildListClustersURL()
	request := client.buildListClustersRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListClustersResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListClustersURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/regions/%s", server, client.Namespace())
}

func (client *Client) buildListClustersRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListClustersResult(response *rest.Response) (*ListClustersResult, error) {
	result := ListClustersResult{}

	err := json.Unmarshal(response.Body(), &result.Clusters)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
