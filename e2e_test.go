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
	settingCluster         string = "test.cluster"
	settingSpace           string = "test.space"
	settingService         string = "test.service"
	settingApp             string = "test.app"
	settingImage           string = "test.image"
	settingLB              string = "test.lb"
	settingVolume          string = "test.volume"
	settingTemplate        string = "test.template"
	settingConfig          string = "test.config"
	settingRegistry        string = "test.registry"
	settingRegistryProject string = "test.registryProject"
	settingRepo            string = "test.repo"
	settingNode            string = "test.node"
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
	{[]string{"alauda", "node", "ls"}},
	{[]string{"alauda", "nodes"}},
	{[]string{"alauda", "node", "inspect", "%NODE%"}},
	{[]string{"alauda", "node", "labels", "%NODE%"}},
	{[]string{"alauda", "node", "labels", "set", "%NODE%", "-l", "type=GPU"}},
	{[]string{"alauda", "node", "cordon", "%NODE%"}},
	{[]string{"alauda", "node", "uncordon", "%NODE%"}},
	{[]string{"alauda", "registries"}},
	{[]string{"alauda", "registry", "ls"}},
	{[]string{"alauda", "registry", "projects", "%REGISTRY%"}},
	{[]string{"alauda", "image", "ls", "-r", "%REGISTRY%", "-p", "%REGISTRYPROJECT%"}},
	{[]string{"alauda", "images"}},
	{[]string{"alauda", "image", "tags", "%REPO%"}},
	{[]string{"alauda", "lbs"}},
	{[]string{"alauda", "lb", "ls"}},
	{[]string{"alauda", "lb", "inspect", "%LB%"}},
	{[]string{"alauda", "volumes"}},
	{[]string{"alauda", "volume", "ls"}},
	{[]string{"alauda", "volume", "create", "%VOLUME%"}},
	{[]string{"alauda", "volume", "inspect", "%VOLUME%"}},
	{[]string{"alauda", "templates"}},
	{[]string{"alauda", "template", "ls"}},
	{[]string{"alauda", "template", "create", "%TEMPLATE%", "-f", "examples/alauda-compose.yml"}},
	{[]string{"alauda", "template", "inspect", "%TEMPLATE%"}},
	{[]string{"alauda", "template", "update", "%TEMPLATE%", "-f", "examples/alauda-compose.yml"}},
	{[]string{"alauda", "apps"}},
	{[]string{"alauda", "compose", "ls"}},
	{[]string{"alauda", "compose", "up", "%APP%", "-t", "%TEMPLATE%", "-s", "--timeout", "120"}},
	{[]string{"alauda", "compose", "inspect", "%APP%"}},
	{[]string{"alauda", "compose", "ps", "%APP%"}},
	{[]string{"alauda", "compose", "rm", "%APP%"}},
	{[]string{"alauda", "template", "rm", "%TEMPLATE%"}},
	{[]string{"alauda", "configs"}},
	{[]string{"alauda", "config", "ls"}},
	{[]string{"alauda", "config", "create", "%CONFIG%", "-d", "test config", "--item", "k1=v1", "-i", "k2=v2"}},
	{[]string{"alauda", "config", "inspect", "%CONFIG%"}},
	{[]string{"alauda", "config", "items", "%CONFIG%"}},
	{[]string{"alauda", "service", "ps"}},
	{[]string{"alauda", "service", "run", "%SERVICE%", "%IMAGE%",
		"-c", "%CLUSTER%", "-s", "%SPACE%",
		"--expose", "80", "--expose", "81",
		"--publish", "10080", "-p", "%LB%:10081:10081", "-p", "%LB%:10082:10082/http",
		"-v", "%VOLUME%:/tempdata", "--volume", "/var:/tempdata2",
		"--cpu", "0.256", "--memory", "256",
		"-n", "2",
		"--env", "FOO=foo", "-e", "BAR=bar",
		"-r", "do this", "--entrypoint", "and that",
		"--config", "%CONFIG%:k1:/config.txt"}},
	{[]string{"alauda", "service", "inspect", "%SERVICE%"}},
	{[]string{"alauda", "lb", "bind", "%LB%", "--listener", "%SERVICE%:80", "-l", "%SERVICE%:81/http", "-l", "%SERVICE%:21234:1234"}},
	{[]string{"alauda", "lb", "unbind", "%LB%", "--listener", "%SERVICE%:80:80", "-l", "%SERVICE%:21234:1234"}},
	{[]string{"alauda", "volume", "rm", "%VOLUME%"}},
	{[]string{"alauda", "config", "rm", "%CONFIG%"}},
	{[]string{"alauda", "service", "rm", "%SERVICE%"}},
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
		args[i] = strings.Replace(args[i], "%SERVICE%", viper.GetString(settingService), -1)
		args[i] = strings.Replace(args[i], "%APP%", viper.GetString(settingApp), -1)
		args[i] = strings.Replace(args[i], "%IMAGE%", viper.GetString(settingImage), -1)
		args[i] = strings.Replace(args[i], "%CLUSTER%", viper.GetString(settingCluster), -1)
		args[i] = strings.Replace(args[i], "%SPACE%", viper.GetString(settingSpace), -1)
		args[i] = strings.Replace(args[i], "%LB%", viper.GetString(settingLB), -1)
		args[i] = strings.Replace(args[i], "%VOLUME%", viper.GetString(settingVolume), -1)
		args[i] = strings.Replace(args[i], "%TEMPLATE%", viper.GetString(settingTemplate), -1)
		args[i] = strings.Replace(args[i], "%CONFIG%", viper.GetString(settingConfig), -1)
		args[i] = strings.Replace(args[i], "%REGISTRY%", viper.GetString(settingRegistry), -1)
		args[i] = strings.Replace(args[i], "%REGISTRYPROJECT%", viper.GetString(settingRegistryProject), -1)
		args[i] = strings.Replace(args[i], "%REPO%", viper.GetString(settingRepo), -1)
		args[i] = strings.Replace(args[i], "%NODE%", viper.GetString(settingNode), -1)
	}
}
