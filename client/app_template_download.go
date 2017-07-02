package client

import (
	"io"
	"net/http"
	"os"
)

// DownloadAppTemplate downloads a specific app template.
func (client *Client) DownloadAppTemplate(name string, filePath string) error {
	template, err := client.InspectAppTemplate(name)
	if err != nil {
		return err
	}

	resp, err := http.Get(template.Template)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
