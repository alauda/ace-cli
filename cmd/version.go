package cmd

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/spf13/cobra"
)

var version = "1.3.0"

func newVersionCmd(alauda client.APIClient) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Display version of Alauda CLI",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[alauda]", version)
		},
	}

	return versionCmd
}
