package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListSpacesResult defines the response body for the ListSpaces API.
type ListSpacesResult struct {
	Spaces []Space
}

// ListSpaces returns all spaces in an account.
func (client *Client) ListSpaces() (*ListSpacesResult, error) {
	url := client.buildURL("v1", "spaces", "")
	request := client.buildListSpacesRequest()

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

func (client *Client) buildListSpacesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListSpacesResult(response *rest.Response) (*ListSpacesResult, error) {
	result := ListSpacesResult{}

	err := json.Unmarshal(response.Body(), &result.Spaces)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
