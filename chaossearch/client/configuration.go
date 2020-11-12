package client

// Configuration stores the configuration for a Client
type Configuration struct {
	URL string
}

// NewConfiguration creates a default Configuration struct
func NewConfiguration() *Configuration {
	cfg := &Configuration{}

	return cfg
}
