package node

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type cordonOptions struct {
	cluster string
}

func newCordonCmd(alauda client.APIClient) *cobra.Command {
	var opts cordonOptions

	cordonCmd := &cobra.Command{
		Use:   "cordon IP",
		Short: "Cordon a node",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node cordon expects IP")
			}
			return doCordon(alauda, args[0], &opts)
		},
	}

	cordonCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return cordonCmd
}

func doCordon(alauda client.APIClient, name string, opts *cordonOptions) error {
	fmt.Println("[alauda] Cordoning", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	data := client.CordonNodeData{
		Action: "cordon",
	}

	err = alauda.CordonNode(name, cluster, &data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
