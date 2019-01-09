package project

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Short: "List projects",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}

	return lsCmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient) error {
	fmt.Println("[alauda] Listing projects")

	util.InitializeClient(alauda)

	result, err := alauda.ListProjects()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListProjectsResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "CREATED", "STATE"}
}

func buildLsTableContent(result *client.ListProjectsResult) [][]string {
	var content [][]string

	for _, project := range result.Projects {
		content = append(content, []string{project.Name, project.CreatedAt, project.State})
	}

	return content
}
