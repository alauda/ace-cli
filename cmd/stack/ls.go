package stack

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
}

// NewLsCmd creates a new stack ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List stacks",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return lsCmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing stacks")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	params := client.ListStacksParams{
		Cluster: cluster,
	}

	result, err := alauda.ListStacks(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(stacks *client.ListStacksResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(stacks)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "ID", "STATE", "CREATED BY"}
}

func buildLsTableContent(result *client.ListStacksResult) [][]string {
	var content [][]string

	for _, stack := range result.Stacks {
		content = append(content, []string{stack.Name, stack.ID, stack.State, stack.CreatedBy})
	}

	return content
}
