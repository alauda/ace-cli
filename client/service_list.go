package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// ListServicesParams defines the query parameters for the ListServices API.
type ListServicesParams struct {
	Cluster string
}

// ListServicesResult defines the response body for the ListServices API.
type ListServicesResult struct {
	Count   int       `json:"count"`
	Results []Service `json:"results"`
}

// Service defines the response body for one service returned in the ListServices API.
type Service struct {
	Name             string              `json:"service_name"`
	ImageName        string              `json:"image_name"`
	ImageTag         string              `json:"image_tag"`
	Command          string              `json:"run_command"`
	Created          string              `json:"created_at"`
	Size             ServiceInstanceSize `json:"custom_instance_size"`
	TargetInstances  int                 `json:"target_num_instances"`
	HealthyInstances int                 `json:"healthy_num_instances"`
	Status           string              `json:"current_status"`
}

// ServiceInstanceSize defines the size of the service instances
type ServiceInstanceSize struct {
	Memory int `json:"mem"`
	CPU    int `json:"cpu"`
}

// ListServices returns all services deployed.
func (client *Client) ListServices(params *ListServicesParams) (*ListServicesResult, error) {
	url := client.buildListServicesURL()
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

func (client *Client) buildListServicesURL() string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s/%s", server, "services", client.Namespace())
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
