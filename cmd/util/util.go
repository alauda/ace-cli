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

// ConfigProject uses the project specified by the cmd, or the one in the config, as appropriate.
func ConfigProject(project string) (string, error) {
	if project != "" {
		viper.Set(SettingProject, project)

		err := SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(SettingProject)
	if result == "" {
		return "", errors.New("no project specified")
	}

	return result, nil
}

// ConfigNamespace uses the namespace specified by the cmd, or the one in the config, as appropriate.
func ConfigNamespace(namespace string) (string, error) {
	if namespace != "" {
		viper.Set(SettingNamespace, namespace)

		err := SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(SettingNamespace)
	if result == "" {
		return "", errors.New("no cluster namespace specified")
	}

	return result, nil
}

// ConfigSpace uses the space specified by the cmd, or the one in the config, as appropriate.
func ConfigSpace(space string) (string, error) {
	if space != "" {
		viper.Set(SettingSpace, space)

		err := SaveConfig()
		if err != nil {
			return "", err
		}
	}

	result := viper.GetString(SettingSpace)
	if result == "" {
		return "", errors.New("no space specified")
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

// ParseKeyValues parses KEY=VALUE into a map of key value pairs.
func ParseKeyValues(descs []string) (map[string]string, error) {
	keyValues := make(map[string]string)

	for _, desc := range descs {
		k, v, err := parseKeyValue(desc)
		if err != nil {
			return nil, err
		}
		keyValues[k] = v
	}

	return keyValues, nil
}

func parseKeyValue(desc string) (string, string, error) {
	result := strings.Split(desc, "=")

	if len(result) != 2 {
		return "", "", errors.New("invalid key-value descriptor")
	}

	return result[0], result[1], nil
}
