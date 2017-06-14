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
	settingName    string = "test.name"
	settingImage   string = "test.image"
	settingLB      string = "test.lb"
	settingVolume  string = "test.volume"
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
	{[]string{"alauda", "volume", "ls"}},
	{[]string{"alauda", "volume", "create", "%VOLUME%"}},
	{[]string{"alauda", "volume", "inspect", "%VOLUME%"}},
	{[]string{"alauda", "compose", "ls"}},
	{[]string{"alauda", "service", "ps"}},
	{[]string{"alauda", "service", "run", "%NAME%", "%IMAGE%",
		"-c", "%CLUSTER%", "-s", "%SPACE%",
		"--expose", "80", "--expose", "81",
		"--publish", "10080", "-p", "%LB%:10081:10081", "-p", "%LB%:10082:10082/http",
		"-v", "%VOLUME%:/tempdata", "--volume", "/var:/tempdata2",
		"--cpu", "0.256", "--memory", "256",
		"-n", "2",
		"--env", "FOO=foo", "-e", "BAR=bar",
		"-r", "do this", "--entrypoint", "and that"}},
	{[]string{"alauda", "service", "inspect", "%NAME%"}},
	{[]string{"alauda", "lb", "bind", "%LB%", "--listener", "%NAME%:80", "-l", "%NAME%:81/http", "-l", "%NAME%:21234:1234"}},
	{[]string{"alauda", "lb", "unbind", "%LB%", "--listener", "%NAME%:80:80", "-l", "%NAME%:21234:1234"}},
	{[]string{"alauda", "service", "rm", "%NAME%"}},
	{[]string{"alauda", "volume", "rm", "%VOLUME%"}},
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
		args[i] = strings.Replace(args[i], "%NAME%", viper.GetString(settingName), -1)
		args[i] = strings.Replace(args[i], "%IMAGE%", viper.GetString(settingImage), -1)
		args[i] = strings.Replace(args[i], "%CLUSTER%", viper.GetString(settingCluster), -1)
		args[i] = strings.Replace(args[i], "%SPACE%", viper.GetString(settingSpace), -1)
		args[i] = strings.Replace(args[i], "%LB%", viper.GetString(settingLB), -1)
		args[i] = strings.Replace(args[i], "%VOLUME%", viper.GetString(settingVolume), -1)
	}
}
