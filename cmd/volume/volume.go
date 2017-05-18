package volume

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewVolumeCmd creates a new volume command.
func NewVolumeCmd(alauda client.APIClient) *cobra.Command {
	volumeCmd := &cobra.Command{
		Use:   "volume",
		Short: "Manage volumes",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	volumeCmd.AddCommand(
		NewLsCmd(alauda),
		NewInspectCmd(alauda),
	)

	return volumeCmd
}

func getClusterID(alauda client.APIClient, name string) (string, error) {
	cluster, err := alauda.InspectCluster(name)
	if err != nil {
		return "", err
	}

	return cluster.ID, nil
}

func getVolumeID(alauda client.APIClient, name string) (string, error) {
	cluster := viper.GetString(util.SettingCluster)

	if cluster != "" {
		clusterID, err := getClusterID(alauda, cluster)
		if err != nil {
			return "", err
		}

		volume, err := getVolumeInCluster(alauda, name, clusterID)
		if err != nil {
			return "", err
		}

		return volume.ID, nil
	}

	result, err := alauda.ListClusters()
	if err != nil {
		return "", err
	}

	for _, cluster := range result.Clusters {
		volume, err := getVolumeInCluster(alauda, name, cluster.ID)
		if err == nil {
			return volume.ID, nil
		}
	}

	return "", fmt.Errorf("volume %s not found", name)
}

func getVolumeInCluster(alauda client.APIClient, name string, clusterID string) (*client.Volume, error) {
	params := client.ListVolumesParams{
		ClusterID: clusterID,
	}

	result, err := alauda.ListVolumes(&params)
	if err != nil {
		return nil, err
	}

	for _, volume := range result.Volumes {
		if volume.Name == name {
			return &volume, nil
		}
	}

	return nil, fmt.Errorf("volume %s not found", name)
}
