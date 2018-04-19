package config

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Short: "List configurations",
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
	fmt.Println("[alauda] Listing configurations")

	util.InitializeClient(alauda)

	result, err := alauda.ListConfigs()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListConfigsResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "DESCRIPTION", "CREATED BY", "CREATED AT"}
}

func buildLsTableContent(result *client.ListConfigsResult) [][]string {
	var content [][]string

	for _, config := range result.Configs {
		content = append(content, []string{config.Name, config.Description, config.CreatedBy, config.CreatedAt})
	}

	return content
}
