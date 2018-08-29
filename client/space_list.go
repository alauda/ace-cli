package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListSpacesParams defines the query parameters for the ListSpaces API.
type ListSpacesParams struct {
	Project string
}

// ListSpacesResult defines the response body for the ListSpaces API.
type ListSpacesResult struct {
	Spaces []Space
}

// ListSpaces returns all spaces in an account.
func (client *Client) ListSpaces(params *ListSpacesParams) (*ListSpacesResult, error) {
	url := client.buildURL("v1", "spaces", "")
	request := client.buildListSpacesRequest(params)

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListSpacesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListSpacesRequest(params *ListSpacesParams) *rest.Request {
	request := rest.NewRequest(client.Token())

	request.SetQueryParam("project_name", params.Project)

	return request
}

func parseListSpacesResult(response *rest.Response) (*ListSpacesResult, error) {
	result := ListSpacesResult{}

	err := json.Unmarshal(response.Body(), &result.Spaces)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
