package service

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type psOptions struct {
	cluster string
}

// NewPsCmd creates a new service ps command.
func NewPsCmd(alauda client.AlaudaClient) *cobra.Command {
	var opts psOptions

	psCmd := &cobra.Command{
		Use:   "ps",
		Short: "List services",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doPs(alauda, opts)
		},
	}

	psCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return psCmd
}

func doPs(alauda client.AlaudaClient, opts psOptions) error {
	util.InitializeClient(alauda)

	params := client.ListServicesParams{
		Cluster: opts.cluster,
	}

	result, err := alauda.ListServices(&params)
	if err != nil {
		return err
	}

	for i := 0; i < result.Count; i++ {
		fmt.Println(result.Results[i].Name, result.Results[i].Image, result.Results[i].Status)
	}
	return nil
}
