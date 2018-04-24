package util

const (
	// DefaultAPIServer specifies the default Alauda API endpoint.
	DefaultAPIServer string = "https://api.alauda.cn/v1/"

	// ConfigFileName specifies the configuration filename with extension.
	ConfigFileName string = ".alauda.yml"

	// SettingServer is the setting for the API server.
	SettingServer string = "auth.server"

	// SettingToken is the setting for the authentication token.
	SettingToken string = "auth.token"

	// SettingAccount is the setting for the account.
	SettingAccount string = "auth.account"

	// SettingUsername is the setting for username.
	SettingUsername string = "auth.username"

	// SettingCluster is the setting for the current cluster.
	SettingCluster string = "context.cluster"

	// SettingNamespace is the setting for the current cluster namespace.
	SettingNamespace string = "context.namespace"

	// SettingSpace is the setting for the current space.
	SettingSpace string = "context.space"
)
