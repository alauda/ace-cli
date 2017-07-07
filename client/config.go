package client

// Config defines the response body for one config returned in the ListConfigs API.
type Config struct {
	Name        string       `json:"name"`
	CreatedBy   string       `json:"created_by"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	Description string       `json:"description"`
	Content     []ConfigItem `json:"content"`
}

// ConfigItem defines one configuration item (key-value pair) in a configuration.
type ConfigItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
