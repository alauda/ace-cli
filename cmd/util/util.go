package util

import (
	"errors"
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
