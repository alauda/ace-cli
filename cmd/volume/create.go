package volume

import (
	"errors"
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type createOptions struct {
	cluster string
	size    int
}

// NewCreateCmd creates a new create volume command.
func NewCreateCmd(alauda client.APIClient) *cobra.Command {
	var opts createOptions

	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create a new volume",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("volume create expects NAME")
			}
			return doCreate(alauda, args[0], &opts)
		},
	}

	createCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster to create the volume in")
	createCmd.Flags().IntVarP(&opts.size, "size", "s", 10, "Volume size (GB)")

	return createCmd
}

func doCreate(alauda client.APIClient, name string, opts *createOptions) error {
	fmt.Println("[alauda] Creating", name)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	clusterID, err := getClusterID(alauda, cluster)
	if err != nil {
		return err
	}

	err = validateSize(opts.size)
	if err != nil {
		return err
	}

	data := client.CreateVolumeData{
		Name:      name,
		Driver:    "glusterfs",
		ClusterID: clusterID,
		Size:      opts.size,
		Namespace: alauda.Namespace(),
	}

	err = alauda.CreateVolume(&data)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}

func validateSize(size int) error {
	if size < 10 || size > 1024 {
		return errors.New("supported volume size (GB): [10, 1024]")
	}

	return nil
}
