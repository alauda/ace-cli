package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListVolumesParams defines the query parameters for the ListVolumes API.
type ListVolumesParams struct {
	ClusterID string
}

// ListVolumesResult defines the response body for the ListVolumes API.
type ListVolumesResult struct {
	Volumes []Volume
}

// ListVolumes returns all volumes in a cluster.
func (client *Client) ListVolumes(params *ListVolumesParams) (*ListVolumesResult, error) {
	url := client.buildURL("v1", "storage", "volumes")
	request := client.buildListVolumesRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListVolumesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListVolumesRequest(params *ListVolumesParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	if params.ClusterID != "" {
		request.SetQueryParam("region_id", params.ClusterID)
	}

	return request
}

func parseListVolumesResult(response *rest.Response) (*ListVolumesResult, error) {
	result := ListVolumesResult{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
