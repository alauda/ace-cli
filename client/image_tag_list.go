package client

import (
	"encoding/json"
	"fmt"

	"github.com/alauda/alauda/client/rest"
)

// ListImageTagsResult defines the response body for the ListImageTags API.
type ListImageTagsResult struct {
	Tags []string
}

// ListImageTags returns the list of tags of an image.
func (client *Client) ListImageTags(registryName string, projectName string, imageName string) (*ListImageTagsResult, error) {
	var url string
	if projectName != "" {
		url = client.buildURL("registries", "%s/projects/%s/repositories/%s/tags", registryName, projectName, imageName)
	} else {
		url = client.buildURL("registries", "%s/repositories/%s/tags", registryName, imageName)
	}

	request := client.buildListImageTagsRequest()

	response, err := request.Get(url)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseListImageTagsResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) buildListImageTagsRequest() *rest.Request {
	return rest.NewRequest(client.Token())
}

func parseListImageTagsResult(response *rest.Response) (*ListImageTagsResult, error) {
	result := ListImageTagsResult{}
	fmt.Println(string(response.Body()))

	err := json.Unmarshal(response.Body(), &result.Tags)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
