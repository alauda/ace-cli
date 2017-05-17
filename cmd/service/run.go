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
	runCmd.Flags().IntSliceVarP(&opts.expose, "expose", "", []int{}, "Ports exposed")
	runCmd.Flags().Float64VarP(&opts.cpu, "cpu", "", 0.125, "CPU (cores) (default: 0.125)")
	runCmd.Flags().IntVarP(&opts.memory, "memory", "", 256, "Memory (MB) (default: 256)")
	runCmd.Flags().IntVarP(&opts.number, "num-instances", "n", 1, "Number of instances (default: 1)")
	runCmd.Flags().StringSliceVarP(&opts.env, "env", "e", []string{}, "Environment variables")
	runCmd.Flags().StringVarP(&opts.cmd, "run-command", "r", "", "Command to run")
	runCmd.Flags().StringVarP(&opts.entrypoint, "entrypoint", "", "", "Entrypoint for the container")
	runCmd.Flags().StringSliceVarP(&opts.publish, "publish", "p", []string{}, "Ports to publish on the load balancer ([lb:][listenerPort:]containerPort[/protocol]")

	return runCmd
}

func doRun(alauda client.APIClient, name string, image string, opts *createOptions) error {
	return doCreate(alauda, name, image, opts, true)
}
