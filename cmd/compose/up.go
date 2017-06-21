package compose

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
)

type upOptions struct {
	cluster  string
	filePath string
	strict   bool
	timeout  int
}

// NewUpCmd creates a new compose up command.
func NewUpCmd(alauda client.APIClient) *cobra.Command {
	var opts upOptions

	upCmd := &cobra.Command{
		Use:   "up NAME",
		Short: "Creates and starts an application",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("compose up expects NAME")
			}
			return doUp(alauda, args[0], &opts)
		},
	}

	upCmd.Flags().StringVarP(&opts.cluster, "cluster", "c", "", "Cluster to create the application in")
	upCmd.Flags().StringVarP(&opts.filePath, "file", "f", "./alauda-compose.yml", "Compose yaml file")
	upCmd.Flags().BoolVarP(&opts.strict, "strict", "s", false, "Start services in strict dependency order")
	upCmd.Flags().IntVarP(&opts.timeout, "timeout", "", 150, "Timeout")

	return upCmd
}

func doUp(alauda client.APIClient, name string, opts *upOptions) error {
	fmt.Println("[alauda] Launching application using compose file", opts.filePath)

	util.InitializeClient(alauda)

	cluster, err := util.ConfigCluster(opts.cluster)
	if err != nil {
		return err
	}

	data := client.CreateAppData{
		Name:      name,
		Cluster:   cluster,
		Namespace: alauda.Namespace(),
		Strict:    opts.strict,
		Timeout:   opts.timeout,
	}

	absPath, err := filepath.Abs(opts.filePath)
	if err != nil {
		return err
	}

	err = alauda.CreateApp(&data, absPath)
	if err != nil {
		return err
	}

	fmt.Println("[alauda] OK")

	return nil
}
