package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListRegistriesResult defines the response body for the ListRegistries API.
type ListRegistriesResult struct {
	Registries []Registry
}

// ListRegistries returns the list of all registries in the account.
func (client *Client) ListRegistries() (*ListRegistriesResult, error) {
	url := client.buildURL("registries", "")
	request := client.buildListRegistriesRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListRegistriesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListRegistriesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListRegistriesResult(response *rest.Response) (*ListRegistriesResult, error) {
	result := ListRegistriesResult{}

	err := json.Unmarshal(response.Body(), &result.Registries)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
