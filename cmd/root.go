package cmd

import (
	"fmt"
	"os"

	"github.com/alauda/alauda/client"
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alauda.yaml)")

	addCommands(rootCmd, alauda)

	return rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func addCommands(cmd *cobra.Command, alauda client.APIClient) {
	cmd.AddCommand(
		NewVersionCmd(alauda),
		NewLoginCmd(alauda),
		NewLogoutCmd(alauda),
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".alauda")
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
