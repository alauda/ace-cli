package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListAppTemplatesResult defines the response body for the ListAppTemplates API.
type ListAppTemplatesResult struct {
	AppTemplates []AppTemplate
}

// ListAppTemplates returns all app templates.
func (client *Client) ListAppTemplates() (*ListAppTemplatesResult, error) {
	url := client.buildURL("application-templates", "")
	request := client.buildListAppTemplatesRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListAppTemplatesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListAppTemplatesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListAppTemplatesResult(response *rest.Response) (*ListAppTemplatesResult, error) {
	result := ListAppTemplatesResult{}

	err := json.Unmarshal(response.Body(), &result.AppTemplates)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
