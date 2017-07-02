package template

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type downloadOptions struct {
	filePath string
}

// NewDownloadCmd creates a new template download command.
func NewDownloadCmd(alauda client.APIClient) *cobra.Command {
	var opts downloadOptions

	downloadCmd := &cobra.Command{
		Use:   "download NAME",
		Short: "Download an app template",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("template download expects NAME")
			}
			return doDownload(alauda, args[0], &opts)
		},
	}

	downloadCmd.Flags().StringVarP(&opts.filePath, "file", "f", "./alauda-yml-out.yml", "Output file to download the template into")

	return downloadCmd
}

func doDownload(alauda client.APIClient, name string, opts *downloadOptions) error {
	fmt.Println("[alauda] Downloading", name)

	util.InitializeClient(alauda)

	err := alauda.DownloadAppTemplate(name, opts.filePath)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
