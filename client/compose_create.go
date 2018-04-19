package client

import (
	"strconv"

	"github.com/alauda/alauda/client/rest"
)

// CreateAppData defines the request body for the CreateApp API.
type CreateAppData struct {
	Name      string `json:"app_name"`
	Cluster   string `json:"region"`
	Namespace string `json:"namespace"`
	Strict    bool   `json:"strict_mode"`
	Timeout   int    `json:"timeout"`
}

// CreateApp creates and starts the specified application.
func (client *Client) CreateApp(data *CreateAppData, filePath string) error {
	url := client.buildURL("applications", "")

	request, err := client.buildCreateAppRequest(data, filePath)
	if err != nil {
		return err
	}

	response, err := request.Post(url)
	if err != nil {
		return err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) buildCreateAppRequest(data *CreateAppData, configFile string) (*rest.Request, error) {
	request := rest.NewRequest(client.Token())

	request.SetFormData(map[string]string{
		"app_name":    data.Name,
		"region":      data.Cluster,
		"namespace":   data.Namespace,
		"strict_mode": strconv.FormatBool(data.Strict),
		"timeout":     strconv.Itoa(data.Timeout),
	})

	request.SetFile("services", configFile)

	return request, nil
}
