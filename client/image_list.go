package client

import (
	"encoding/json"

	"github.com/alauda/alauda/client/rest"
)

// ListImagesResult defines the response body for the ListImages API.
type ListImagesResult struct {
	Images []Image
}

// ListImages returns the list of all images, scoped by the parent registry.
func (client *Client) ListImages(registryName string, projectName string) (*ListImagesResult, error) {
	var url string
	if projectName != "" {
		url = client.buildURL("registries", "%s/projects/%s/repositories", registryName, projectName)
	} else {
		url = client.buildURL("registries", "%s/repositories", registryName)
	}

	request := client.buildListImagesRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListImagesResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListImagesRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListImagesResult(response *rest.Response) (*ListImagesResult, error) {
	result := ListImagesResult{}

	err := json.Unmarshal(response.Body(), &result.Images)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
