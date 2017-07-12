package image

import (
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	registry string
}

// NewLsCmd creates a new image ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List images",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.registry, "registry", "r", "alauda_public_registry", "Registry")

	return lsCmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing images")

	util.InitializeClient(alauda)

	result, err := alauda.ListImages(opts.registry)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListImagesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "ID", "DESCRIPTION", "IS PUBLIC", "CREATED BY", "CREATED AT"}
}

func buildLsTableContent(result *client.ListImagesResult) [][]string {
	var content [][]string

	for _, image := range result.Images {
		content = append(content, []string{image.Name, image.ID, image.Description,
			strconv.FormatBool(image.IsPublic), image.CreatedBy, image.CreatedAt})
	}

	return content
}
