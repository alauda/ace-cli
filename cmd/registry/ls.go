package registry

import (
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Short: "List registries",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}

	return lsCmd
}

// NewRegistriesCmd creates a new alauda registries command, which is a shortcut to the registry ls command.
func NewRegistriesCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "registries"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient) error {
	fmt.Println("[alauda] Listing registries")

	util.InitializeClient(alauda)

	result, err := alauda.ListRegistries()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListRegistriesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "ID", "DESCRIPTION", "CLUSTER", "ENDPOINT", "IS PUBLIC", "CREATED BY", "CREATED AT"}
}

func buildLsTableContent(result *client.ListRegistriesResult) [][]string {
	var content [][]string

	for _, registry := range result.Registries {
		content = append(content, []string{registry.Name, registry.ID, registry.Description, registry.Cluster,
			registry.Endpoint, strconv.FormatBool(registry.IsPublic), registry.CreatedBy, registry.CreatedAt})
	}

	return content
}
