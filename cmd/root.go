package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/app"
	"github.com/alauda/alauda/cmd/cluster"
	"github.com/alauda/alauda/cmd/image"
	"github.com/alauda/alauda/cmd/namespace"
	"github.com/alauda/alauda/cmd/project"
	"github.com/alauda/alauda/cmd/registry"
	"github.com/alauda/alauda/cmd/space"
	"github.com/alauda/alauda/cmd/util"
	"github.com/alauda/alauda/cmd/kubectl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// NewRootCmd creates a new root command for the Alauda CLI.
func NewRootCmd(alauda client.APIClient) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "alauda",
		Short: "Alauda CLI",
		Long:  ``,
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: $HOME/.alauda.yml)")

	addCommands(rootCmd, alauda)

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func addCommands(cmd *cobra.Command, alauda client.APIClient) {
	cmd.AddCommand(
		newVersionCmd(alauda),
		newLoginCmd(alauda),
		newLogoutCmd(alauda),

		// Adding application related commands.
		app.NewAppCmd(alauda),
		app.NewAppsCmd(alauda),
		app.NewLsCmd(alauda),
		app.NewInspectCmd(alauda),
		app.NewYamlCmd(alauda),
		app.NewStartCmd(alauda),
		app.NewStopCmd(alauda),
		app.NewRmCmd(alauda),
		app.NewRunCmd(alauda),

		// Adding image related commands.
		image.NewImageCmd(alauda),
		image.NewImagesCmd(alauda),

		// Adding cluster related commands.
		cluster.NewClusterCmd(alauda),
		cluster.NewClustersCmd(alauda),

		// Adding node related commands.
		// node.NewNodeCmd(alauda),
		// node.NewNodesCmd(alauda),

		// Adding lb related commands.
		// lb.NewLBCmd(alauda),
		// lb.NewLbsCmd(alauda),

		// Adding project related commands.
		project.NewProjectCmd(alauda),
		project.NewProjectsCmd(alauda),

		// Adding registry related commands.
		registry.NewRegistryCmd(alauda),
		registry.NewRegistriesCmd(alauda),

		// Adding space related commands.
		space.NewSpaceCmd(alauda),
		space.NewSpacesCmd(alauda),

		// Adding namespace related commands.
		namespace.NewNamespaceCmd(alauda),
		namespace.NewNamespacesCmd(alauda),

		// Adding embedded kubectl commands.
		kubectl.NewKubectlCmd(),
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		// Use default.
		cfgFile = filepath.Join(os.Getenv("HOME"), util.ConfigFileName)
	}
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	viper.SetDefault(util.SettingServer, util.DefaultAPIServer)

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("[alauda] Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("[alauda] WARNING: Unable to read config file:", cfgFile)
	}
}
