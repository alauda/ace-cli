package image

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newTagsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	tagsCmd := &cobra.Command{
		Use:   "tags",
		Short: "List image tags",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("image tags expects NAME")
			}
			return doTags(alauda, args[0], &opts)
		},
	}

	tagsCmd.Flags().StringVarP(&opts.registry, "registry", "r", "alauda_public_registry", "Registry")
	tagsCmd.Flags().StringVarP(&opts.project, "project", "p", "", "Registry project")

	return tagsCmd
}

func doTags(alauda client.APIClient, imageName string, opts *lsOptions) error {
	fmt.Println("[alauda] Listing image tags")

	util.InitializeClient(alauda)

	result, err := alauda.ListImageTags(opts.registry, opts.project, imageName)
	if err != nil {
		return err
	}

	printTagsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printTagsResult(result *client.ListImageTagsResult) {
	header := buildTagsTableHeader()
	content := buildTagsTableContent(result)

	util.PrintTable(header, content)
}

func buildTagsTableHeader() []string {
	return []string{"NAME"}
}

func buildTagsTableContent(result *client.ListImageTagsResult) [][]string {
	var content [][]string

	for _, tag := range result.Tags {
		content = append(content, []string{tag})
	}

	return content
}
