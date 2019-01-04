package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd"
	"github.com/spf13/viper"
)

const (
	settingCluster string = "test.cluster"
	settingSpace   string = "test.space"
	settingApp     string = "test.app"
)

var cliTests = []struct {
	args []string
}{
	{[]string{"alauda", "spaces"}},
	{[]string{"alauda", "space", "ls"}},
	{[]string{"alauda", "space", "get"}},
	{[]string{"alauda", "space", "set", "%SPACE%"}},
	{[]string{"alauda", "space", "inspect", "%SPACE%"}},
	{[]string{"alauda", "clusters"}},
	{[]string{"alauda", "cluster", "ls"}},
	{[]string{"alauda", "cluster", "get"}},
	{[]string{"alauda", "cluster", "set", "%CLUSTER%"}},
	{[]string{"alauda", "cluster", "inspect", "%CLUSTER%"}},
	{[]string{"alauda", "app", "ls"}},
	{[]string{"alauda", "apps"}},
}

func TestCli(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	for _, tt := range cliTests {
		bind(tt.args)
		os.Args = tt.args
		fmt.Println(os.Args)

		err = rootCmd.Execute()
		if err != nil {
			t.Error(err)
		}
	}
}

func bind(args []string) {
	for i := range args {
		args[i] = strings.Replace(args[i], "%APP%", viper.GetString(settingApp), -1)
		args[i] = strings.Replace(args[i], "%CLUSTER%", viper.GetString(settingCluster), -1)
		args[i] = strings.Replace(args[i], "%SPACE%", viper.GetString(settingSpace), -1)
	}
}
