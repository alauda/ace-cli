package node

import (
	"fmt"
	"strconv"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
}

func newBaseLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Short: "List nodes",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return lsCmd
}

// NewNodesCmd creates a new alauda nodes command, which is a shortcut to the node ls command.
func NewNodesCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "nodes"
	return cmd
}

func newLsCmd(alauda client.APIClient) *cobra.Command {
	cmd := newBaseLsCmd(alauda)
	cmd.Use = "ls"
	return cmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing nodes")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	result, err := alauda.ListNodes(cluster)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(result *client.ListNodesResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(result)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"IP", "STATE", "TYPE", "SCHEDULABLE", "RESOURCES"}
}

func buildLsTableContent(result *client.ListNodesResult) [][]string {
	var content [][]string

	for _, node := range result.Nodes {
		schedulable := strconv.FormatBool(node.Attributes.Schedulable)
		resources := fmt.Sprintf("CPUs: %s/%s, Memory: %s/%s",
			node.Resources.AvailableCPUs, node.Resources.TotalCPUs,
			node.Resources.AvailableMemory, node.Resources.TotalMemory)
		content = append(content, []string{node.IP, node.State, node.Type, schedulable, resources})
	}

	return content
}
