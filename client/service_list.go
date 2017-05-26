package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListServicesParams defines the query parameters for the ListServices API.
type ListServicesParams struct {
	Cluster string
}

// ListServicesResult defines the response body for the ListServices API.
type ListServicesResult struct {
	Count    int       `json:"count"`
	Services []Service `json:"results"`
}

// ListServices returns all services deployed.
func (client *Client) ListServices(params *ListServicesParams) (*ListServicesResult, error) {
	url := client.buildURL("services", "")
	request := client.buildListServicesRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListServicesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListServicesRequest(params *ListServicesParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.Cluster != "" {
		request.SetQueryParam("region_name", params.Cluster)
	}

	request.SetQueryParam("detail", "true")

	return request
}

func parseListServicesResult(response *rest.Response) (*ListServicesResult, error) {
	result := ListServicesResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
