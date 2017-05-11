package main

import (
	"os"
	"testing"

	"github.com/alauda/alauda/client"
	"github.com/alauda/alauda/cmd"
	"github.com/spf13/viper"
)

const (
	settingCluster string = "test.cluster"
	settingSpace   string = "test.space"
	settingName    string = "test.name"
	settingImage   string = "test.image"
	settingLB      string = "test.lb"
)

var cliTests = []struct {
	args []string
}{
	{[]string{"alauda", "space", "ls"}},
	{[]string{"alauda", "space", "inspect", "%SPACE%"}},
	{[]string{"alauda", "cluster", "ls"}},
	{[]string{"alauda", "cluster", "inspect", "%CLUSTER%"}},
	{[]string{"alauda", "lb", "ls"}},
	{[]string{"alauda", "lb", "inspect", "%LB%"}},
	{[]string{"alauda", "service", "ps"}},
	{[]string{"alauda", "service", "run", "%NAME%", "%IMAGE%",
		"-c", "%CLUSTER%", "-s", "%SPACE%",
		"--expose", "80", "--expose", "81",
		"--cpu", "0.256", "--memory", "256",
		"-n", "2",
		"--env", "FOO=foo", "-e", "BAR=bar",
		"-r", "do this", "--entrypoint", "and that"}},
	{[]string{"alauda", "service", "inspect", "%NAME%"}},
	{[]string{"alauda", "service", "rm", "%NAME%"}},
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

		err = rootCmd.Execute()
		if err != nil {
			t.Error(err)
		}
	}
}

func bind(args []string) {
	for i, s := range args {
		switch s {
		case "%NAME%":
			args[i] = viper.GetString(settingName)
		case "%IMAGE%":
			args[i] = viper.GetString(settingImage)
		case "%CLUSTER%":
			args[i] = viper.GetString(settingCluster)
		case "%SPACE%":
			args[i] = viper.GetString(settingSpace)
		case "%LB%":
			args[i] = viper.GetString(settingLB)
		}
	}
}
