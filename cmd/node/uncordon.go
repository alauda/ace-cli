package node

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type uncordonOptions struct {
	cluster string
}

func newUncordonCmd(alauda client.APIClient) *cobra.Command {
	var opts uncordonOptions

	cordonCmd := &cobra.Command{
		Use:   "uncordon IP",
		Short: "Make a node schedulable again",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node uncordon expects IP")
			}
			return doUncordon(alauda, args[0], &opts)
		},
	}

	cordonCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")

	return cordonCmd
}

func doUncordon(alauda client.APIClient, name string, opts *uncordonOptions) error {
	fmt.Println("[alauda] Uncordoning", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	err = alauda.UncordonNode(name, cluster)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
