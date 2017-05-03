package service

import (
	"errors"

	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

// NewRunCmd creates a new run service command.
func NewRunCmd(alauda client.APIClient) *cobra.Command {
	var opts createOptions

	runCmd := &cobra.Command{
		Use:   "run NAME IMAGE",
		Short: "Create and start a new service",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("service run expects NAME and IMAGE")
			}
			return doRun(alauda, args[0], args[1], &opts)
		},
	}

	runCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster to run the service in")
	runCmd.Flags().StringVarP(&opts.space, "space", "s", "", "Space to run the service in")

	return runCmd
}

func doRun(alauda client.APIClient, name string, image string, opts *createOptions) error {
	return doCreate(alauda, name, image, opts, true)
}
