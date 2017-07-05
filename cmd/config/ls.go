package config

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewLsCmd creates a new config ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List configurations",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}

	return lsCmd
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
