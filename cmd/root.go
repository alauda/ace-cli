package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/app"
	"github.com/alauda/alauda/cmd/cluster"
	"github.com/alauda/alauda/cmd/image"
	"github.com/alauda/alauda/cmd/lb"
	"github.com/alauda/alauda/cmd/node"
	"github.com/alauda/alauda/cmd/registry"
	"github.com/alauda/alauda/cmd/space"
	"github.com/alauda/alauda/cmd/util"
	"github.com/alauda/alauda/cmd/volume"
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

		// Adding image related shortcuts.
		image.NewImagesCmd(alauda),

		// Adding cluster related shortcuts
		cluster.NewClustersCmd(alauda),

		// Adding node related shortcuts
		node.NewNodesCmd(alauda),

		// Adding lb related shortcuts
		lb.NewLbsCmd(alauda),

		// Adding registry related shortcuts
		registry.NewRegistriesCmd(alauda),

		// Adding space related shortcuts
		space.NewSpacesCmd(alauda),

		// Adding volume related shortcuts
		volume.NewVolumesCmd(alauda),

		app.NewAppCmd(alauda),
		space.NewSpaceCmd(alauda),
		cluster.NewClusterCmd(alauda),
		lb.NewLBCmd(alauda),
		volume.NewVolumeCmd(alauda),
		registry.NewRegistryCmd(alauda),
		image.NewImageCmd(alauda),
		node.NewNodeCmd(alauda),
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
