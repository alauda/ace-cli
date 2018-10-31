package app

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type startOptions struct {
	cluster   string
	namespace string
}

// NewStartCmd creates a new alauda app start command.
func NewStartCmd(alauda client.APIClient) *cobra.Command {
	var opts startOptions

	startCmd := &cobra.Command{
		Use:   "start NAME",
		Short: "Start an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("app start expects NAME")
			}
			return doStart(alauda, &opts, args[0])
		},
	}

	startCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	startCmd.Flags().StringVarP(&opts.namespace, "namespace", "n", "", "Namespace")

	return startCmd
}

func doStart(alauda client.APIClient, opts *startOptions, name string) error {
	fmt.Println("[alauda] Starting", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	namespace, err := util.ConfigNamespace(opts.namespace)
	if err != nil {
		return err
	}

	err = alauda.StartApp(cluster, namespace, name)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
