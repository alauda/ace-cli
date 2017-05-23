package util

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// ParseImageNameTag extracts the image name and image tag.
func ParseImageNameTag(image string) (string, string, error) {
	result := strings.Split(image, ":")

	switch len(result) {
	case 1:
		return image, "latest", nil
	case 2:
		return result[0], result[1], nil
	}
	return "", "", errors.New("invalid image name format")
}

// ConfigCluster uses the cluster specified by the cmd, or the one in the config, as appropriate.
func ConfigCluster(cluster string) (string, error) {
	if cluster != "" {
		viper.Set(SettingCluster, cluster)

		err := SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(SettingCluster)
	if result == "" {
		return "", errors.New("no cluster specified")
	}

	return result, nil
}

// ParseListener parses the listener information specified in the form name:listenerPort:containerPort/protocol
func ParseListener(desc string) (string, int, int, string, error) {
	var name string
	var listenerPort int
	var containerPort int
	var protocol string
	var err error

	result := strings.Split(desc, "/")

	if len(result) > 2 {
		return "", 0, 0, "", errors.New("invalid listener descriptor, expecting [name:][listenerPort:]containerPort[/protocol]")
	}

	if len(result) == 2 {
		desc = result[0]
		protocol = result[1]

		if protocol != "http" && protocol != "tcp" {
			return "", 0, 0, "", errors.New("invalid protocol specified, supported protocols are [tcp, http]")
		}
	}

	result = strings.Split(desc, ":")

	if len(result) > 3 {
		return "", 0, 0, "", errors.New("invalid listener descriptor, expecting [name:][listenerPort:]containerPort")
	}

	switch len(result) {
	case 1:
		// containerPort
		containerPort, err = strconv.Atoi(result[0])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, containerPort should be int")
		}
	case 2:
		// name:containerPort or listenerPort:containerPort
		containerPort, err = strconv.Atoi(result[1])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, expecting name:containerPort or listenerPort:containerPort")
		}

		listenerPort, err = strconv.Atoi(result[0])
		if err != nil {
			name = result[0]
		}
	case 3:
		// name:listenerPort:containerPort
		name = result[0]

		listenerPort, err = strconv.Atoi(result[1])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, listenerPort is not int, in name:listenerPort:containerPort")
		}

		containerPort, err = strconv.Atoi(result[2])
		if err != nil {
			return "", 0, 0, "", errors.New("invalid listener descriptor, containerPort is not int, in name:listenerPort:containerPort")
		}
	}

	return name, listenerPort, containerPort, protocol, nil
}
