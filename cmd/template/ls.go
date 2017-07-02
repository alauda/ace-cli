package template

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewLsCmd creates a new compose ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List app templates",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}

	return lsCmd
}

func doLs(alauda client.APIClient) error {
	fmt.Println("[alauda] Listing app templates")

	util.InitializeClient(alauda)

	result, err := alauda.ListAppTemplates()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListAppTemplatesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "ID", "TEMPLATE", "CREATED AT", "CREATED BY"}
}

func buildLsTableContent(result *client.ListAppTemplatesResult) [][]string {
	var content [][]string

	for _, template := range result.AppTemplates {
		content = append(content, []string{template.Name, template.ID, template.Template, template.CreatedAt, template.CreatedBy})
	}

	return content
}