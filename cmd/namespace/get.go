package namespace

import (
	"fmt"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newGetCmd(alauda client.APIClient) *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get the current cluster namespace",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doGet(alauda)
		},
	}
	return getCmd
}

func doGet(alauda client.APIClient) error {
	fmt.Println("[alauda] Getting the current cluster namespace")

	space := viper.Get(util.SettingNamespace)

	fmt.Println(space)

	return nil
}
