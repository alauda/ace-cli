package volume

import (
	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
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
