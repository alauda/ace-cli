package cluster

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

// NewLsCmd creates a new cluster ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List clusters",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda)
		},
	}
	return lsCmd
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

func printLsResult(spaces *client.ListClustersResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(spaces)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "DISPLAY NAME", "TYPE", "CREATED", "STATE"}
}

func buildLsTableContent(result *client.ListClustersResult) [][]string {
	var content [][]string

	for i := 0; i < len(result.Clusters); i++ {
		cluster := result.Clusters[i]
		content = append(content, []string{cluster.Name, cluster.DisplayName, cluster.Type, cluster.CreatedAt, cluster.State})
	}

	return content
}
