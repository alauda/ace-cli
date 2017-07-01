package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// InspectAppTemplate retrieves details about a specific app template.
func (client *Client) InspectAppTemplate(name string) (*AppTemplate, error) {
	url := client.buildURL("application-templates", name)
	request := client.buildInspectAppTemplateRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectAppTemplateResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectAppTemplateRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectAppTemplateResult(response *rest.Response) (*AppTemplate, error) {
	result := AppTemplate{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
