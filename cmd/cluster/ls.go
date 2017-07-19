package cluster

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Short: "List clusters",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}
	return lsCmd
}

// NewClustersCmd creates a new alauda clusters command, which is a shortcut to the cluster ls command.
func NewClustersCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "clusters"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient) error {
	fmt.Println("[alauda] Listing clusters")

	util.InitializeClient(alauda)

	result, err := alauda.ListClusters()
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListClustersResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "DISPLAY NAME", "TYPE", "CLOUD", "REGION", "CREATED", "STATE"}
}

func buildLsTableContent(result *client.ListClustersResult) [][]string {
	var content [][]string

	for _, cluster := range result.Clusters {
		content = append(content, []string{cluster.Name, cluster.DisplayName, cluster.Type,
			cluster.Attributes.Cloud.Name, cluster.Attributes.Cloud.Region, cluster.CreatedAt, cluster.State})
	}

	return content
}
