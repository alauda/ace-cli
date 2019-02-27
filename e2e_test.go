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
	settingProject string = "test.project"
	settingSpace   string = "test.space"
	settingCluster string = "test.cluster"
	settingApp     string = "test.app"
)

var cliTests = []struct {
	args []string
}{
	{[]string{"alauda", "projects"}},
	{[]string{"alauda", "project", "ls"}},
	{[]string{"alauda", "project", "get"}},
	{[]string{"alauda", "project", "set", "%PROJECT%"}},
	{[]string{"alauda", "project", "inspect", "%PROJECT%"}},
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
	{[]string{"alauda", "namespaces"}},
	{[]string{"alauda", "namespace", "ls"}},
	{[]string{"alauda", "apps"}},
	{[]string{"alauda", "app", "ls"}},
	{[]string{"alauda", "app", "run", "%APP%", "index.alauda.cn/alauda/hello-world:latest"}},
	{[]string{"alauda", "app", "inspect", "%APP%"}},
	{[]string{"alauda", "app", "yaml", "%APP%"}},
	{[]string{"alauda", "app", "stop", "%APP%"}},
	{[]string{"alauda", "app", "start", "%APP%"}},
	{[]string{"alauda", "app", "rm", "%APP%"}},
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
		args[i] = strings.Replace(args[i], "%PROJECT%", viper.GetString(settingProject), -1)
		args[i] = strings.Replace(args[i], "%SPACE%", viper.GetString(settingSpace), -1)
		args[i] = strings.Replace(args[i], "%CLUSTER%", viper.GetString(settingCluster), -1)
		args[i] = strings.Replace(args[i], "%APP%", viper.GetString(settingApp), -1)
	}
}
