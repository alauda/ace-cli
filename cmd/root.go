package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd/cluster"
	"github.com/alauda/alauda/cmd/compose"
	"github.com/alauda/alauda/cmd/config"
	"github.com/alauda/alauda/cmd/image"
	"github.com/alauda/alauda/cmd/lb"
	"github.com/alauda/alauda/cmd/node"
	"github.com/alauda/alauda/cmd/registry"
	"github.com/alauda/alauda/cmd/service"
	"github.com/alauda/alauda/cmd/space"
	"github.com/alauda/alauda/cmd/template"
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

		// Adding service related commands as top level defaults.
		service.NewServiceCmd(alauda),
		service.NewPsCmd(alauda),
		service.NewCreateCmd(alauda),
		service.NewRunCmd(alauda),
		service.NewStartCmd(alauda),
		service.NewStopCmd(alauda),
		service.NewRmCmd(alauda),
		service.NewInspectCmd(alauda),
		service.NewRestartCmd(alauda),
		service.NewScaleCmd(alauda),
		service.NewUpdateCmd(alauda),

		// Adding image related shortcuts.
		image.NewImagesCmd(alauda),

		// Adding cluster related shortcuts
		cluster.NewClustersCmd(alauda),

		// Adding node related shortcuts
		node.NewNodesCmd(alauda),

		// Adding app related shortcuts
		compose.NewAppsCmd(alauda),

		// Adding config related shortcuts
		config.NewConfigsCmd(alauda),

		space.NewSpaceCmd(alauda),
		cluster.NewClusterCmd(alauda),
		lb.NewLBCmd(alauda),
		volume.NewVolumeCmd(alauda),
		compose.NewComposeCmd(alauda),
		template.NewTemplateCmd(alauda),
		config.NewConfigCmd(alauda),
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
