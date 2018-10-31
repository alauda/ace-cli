package app

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type stopOptions struct {
	cluster   string
	namespace string
}

// NewStopCmd creates a new alauda app stop command.
func NewStopCmd(alauda client.APIClient) *cobra.Command {
	var opts stopOptions

	stopCmd := &cobra.Command{
		Use:   "stop NAME",
		Short: "Stop an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("app stop expects NAME")
			}
			return doStop(alauda, &opts, args[0])
		},
	}

	stopCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	stopCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return stopCmd
}

func doStop(alauda client.APIClient, opts *stopOptions, name string) error {
	fmt.Println("[alauda] Stopping", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	err = alauda.StopApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
