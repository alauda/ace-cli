package labels

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type setOptions struct {
	cluster string
	labels  []string
}

func newSetCmd(alauda client.APIClient) *cobra.Command {
	var opts setOptions

	setCmd := &cobra.Command{
		Use:   "set IP",
		Short: "Set labels of a node",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("node labels set expects IP")
			}
			return doSet(alauda, args[0], &opts)
		},
	}

	setCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster")
	setCmd.Flags().StringSliceVarP(&opts.labels, "labels", "l", []string{}, "Labels to set")

	return setCmd
}

func doSet(alauda client.APIClient, name string, opts *setOptions) error {
	fmt.Println("[alauda] Setting labels for", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	labels, err := util.ParseKeyValues(opts.labels)
	if err != nil {
		return err
	}

	data := client.SetNodeLabelsData(labels)

	err = alauda.SetNodeLabels(name, cluster, &data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
