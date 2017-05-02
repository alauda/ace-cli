package util

import (
	"errors"
	"strings"
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
