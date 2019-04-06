package configs

import "github.com/nabillarahmanizhafira/test_project/connections/cache"

type (
	// Configuration struct holds all configuration from config.yml
	Configuration struct {
		// Configs
		Server ServerConfiguration
		Redis  RedisConfiguration
	}

	// ServerConfiguration will hold server configs
	ServerConfiguration struct {
		Port string
	}

	// RedisConfiguration will hold redis configs
	RedisConfiguration struct {
		Host     string
		PoolSize string
	}

	// AppConfig will hold connection
	AppConfig struct {
		redis cache.Cache
	}
)

var (
	// VarsConfig holds env vars
	VarsConfig Configuration
	// GlobalConfig will be accessible to all
	GlobalConfig AppConfig
)
