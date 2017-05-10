package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// InspectSpace retrieves details about the specified space.
func (client *Client) InspectSpace(name string) (*Space, error) {
	url := client.buildInspectSpaceURL(name)
	request := client.buildInspectSpaceRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseInspectSpaceResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildInspectSpaceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/spaces/%s/space/%s", server, client.Namespace(), name)
}

func (client *Client) buildInspectSpaceRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseInspectSpaceResult(response *rest.Response) (*Space, error) {
	result := Space{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
