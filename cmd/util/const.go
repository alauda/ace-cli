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

	// SettingNamespace is the setting for the namespace.
	SettingNamespace string = "auth.namespace"

	// SettingUsername is the setting for username.
	SettingUsername string = "auth.username"

	// SettingCluster is the setting for the current cluster.
	SettingCluster string = "context.cluster"
)
