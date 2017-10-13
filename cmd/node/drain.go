package node

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type drainOptions struct {
	cluster string
}

func newDrainCmd(alauda client.APIClient) *cobra.Command {
	var opts drainOptions

	drainCmd := &cobra.Command{
		Use:   "drain IP",
		Short: "Migrate containers off the node and make it unschedulable",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node drain expects IP")
			}
			return doDrain(alauda, args[0], &opts)
		},
	}

	drainCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return drainCmd
}

func doDrain(alauda client.APIClient, name string, opts *drainOptions) error {
	fmt.Println("[alauda] Draining", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	err = alauda.DrainNode(name, cluster)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
