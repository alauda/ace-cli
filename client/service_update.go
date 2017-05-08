package client

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alauda/alauda/client/rest"
)

// UpdateServiceData defines the request body for the ScaleService API.
type UpdateServiceData struct {
	ImageTag           string              `json:"image_tag"`
	Command            string              `json:"run_command"`
	Entrypoint         string              `json:"entrypoint"`
	InstanceSize       string              `json:"instance_size"`
	CustomInstanceSize ServiceInstanceSize `json:"custom_instance_size"`
	Env                map[string]string   `json:"instance_envvars"`
}

// UpdateService scales the service to the specified number of instances
func (client *Client) UpdateService(name string, data *UpdateServiceData) error {
	url := client.buildUpdateServiceURL(name)

	request, err := client.buildUpdateServiceRequest(data)
	if err != nil {
		return err
	}

	response, err := request.Put(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildUpdateServiceURL(name string) string {
	server := strings.TrimSuffix(client.APIServer(), "/")
	return fmt.Sprintf("%s/%s/%s/%s", server, "services", client.Namespace(), name)
}

func (client *Client) buildUpdateServiceRequest(data *UpdateServiceData) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	updated := make(map[string]interface{})

	if data.ImageTag != "" {
		updated["image_tag"] = data.ImageTag
	}
	if data.Command != "" {
		updated["run_command"] = data.Command
	}
	if data.Entrypoint != "" {
		updated["entrypoint"] = data.Entrypoint
	}
	if data.InstanceSize != "" {
		updated["instance_size"] = data.InstanceSize
	}
	if data.InstanceSize == "CUSTOMIZED" {
		updated["custom_instance_size"] = data.CustomInstanceSize
	}
	if data.Env != nil {
		updated["instance_envvars"] = data.Env
	}

	body, err := json.Marshal(updated)
	if err != nil {
		return nil, err
	}

	request.SetBody(body)

	return request, nil
}
