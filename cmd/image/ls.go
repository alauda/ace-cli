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
	project  string
}

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		// Use: this will be specified in the wrapper commands.
		Short: "List images",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.registry, "registry", "r", "alauda_public_registry", "Registry")
	lsCmd.Flags().StringVarP(&opts.project, "project", "p", "", "Registry project")

	return lsCmd
}

// NewImagesCmd creates a new alauda images command, which is a shortcut to the image ls command.
func NewImagesCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "images"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing images")

	util.InitializeClient(alauda)

	result, err := alauda.ListImages(opts.registry, opts.project)
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
