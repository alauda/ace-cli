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
)

func TestSpaceLs(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	os.Args = []string{"alauda", "space", "ls"}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestSpaceInspect(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	space := viper.GetString(settingSpace)

	os.Args = []string{"alauda", "space", "inspect", space}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestClusterLs(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	os.Args = []string{"alauda", "cluster", "ls"}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestClusterInspect(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	cluster := viper.GetString(settingCluster)

	os.Args = []string{"alauda", "cluster", "inspect", cluster}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestLBLs(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	os.Args = []string{"alauda", "lb", "ls"}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestServicePs(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	os.Args = []string{"alauda", "service", "ps"}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestServiceRun(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	name := viper.GetString(settingName)
	image := viper.GetString(settingImage)
	cluster := viper.GetString(settingCluster)
	space := viper.GetString(settingSpace)

	os.Args = []string{"alauda", "service", "run", name, image,
		"-c", cluster, "-s", space,
		"--expose", "80", "--expose", "81",
		"--cpu", "0.256", "--memory", "256",
		"-n", "2",
		"--env", "FOO=foo", "-e", "BAR=bar",
		"-r", "do this", "--entrypoint", "and that"}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestServiceInspect(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	name := viper.GetString(settingName)

	os.Args = []string{"alauda", "service", "inspect", name}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}

func TestServiceRm(t *testing.T) {
	alauda, err := client.NewClient()
	if err != nil {
		t.Error(err)
	}

	rootCmd := cmd.NewRootCmd(alauda)

	name := viper.GetString(settingName)

	os.Args = []string{"alauda", "service", "rm", name}

	err = rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}
}
