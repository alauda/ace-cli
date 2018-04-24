package util

import (
	"os"

	"github.com/alauda/alauda/client"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// SaveConfig writes all configuration back to the config file
// This is a temporary workaround until this PR gets merged.
// https://github.com/spf13/viper/pull/287
func SaveConfig() error {
	file, err := os.Create(viper.ConfigFileUsed())
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := yaml.Marshal(viper.AllSettings())
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// InitializeClient initializes the Alauda client from configuration.
func InitializeClient(client client.APIClient) {
	server := viper.GetString(SettingServer)

	account := viper.GetString(SettingAccount)
	if account == "" {
		account = viper.GetString(SettingUsername)
	}

	token := viper.GetString(SettingToken)

	client.Initialize(server, account, token)
}
