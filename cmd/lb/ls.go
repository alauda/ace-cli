package lb

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type lsOptions struct {
	cluster string
	service string
}

// NewLsCmd creates a new lb ls command.
func NewLsCmd(alauda client.APIClient) *cobra.Command {
	var opts lsOptions

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "List load balancers",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doLs(alauda, &opts)
		},
	}

	lsCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	lsCmd.Flags().StringVarP(&opts.service, "service", "s", "", "Service")

	return lsCmd
}

func doLs(alauda client.APIClient, opts *lsOptions) error {
	fmt.Println("[alauda] Listing load balancers")

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	params := client.ListLoadBalancersParams{
		Cluster: cluster,
		Service: opts.service,
	}

	result, err := alauda.ListLoadBalancers(&params)
	if err != nil {
		return err
	}

	printLsResult(result)

	fmt.Println("[alauda] OK")

	return nil
}

func printLsResult(lbs *client.ListLoadBalancersResult) {
	header := buildLsTableHeader()
	content := buildLsTableContent(lbs)

	util.PrintTable(header, content)
}

func buildLsTableHeader() []string {
	return []string{"NAME", "TYPE", "ADDRESS", "ADDRESS TYPE", "CREATED WITH", "CREATED AT"}
}

func buildLsTableContent(result *client.ListLoadBalancersResult) [][]string {
	var content [][]string

	for i := 0; i < len(result.LoadBalancers); i++ {
		lb := result.LoadBalancers[i]
		content = append(content, []string{lb.Name, lb.Type, lb.Address, lb.AddressType, lb.CreatedWith, lb.CreatedAt})
	}

	return content
}
