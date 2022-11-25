package internal

import (
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	StlLogLevel int // Standard core log level
}

type RedisConfig struct {
	Address        string
	Password       string
	ExpirationTime time.Duration
}

func NewServerConfig(prefix string) *ServerConfig {
	// Viper Configuration
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()

	return &ServerConfig{
		StlLogLevel: v.GetInt("std_level"),
	}
}
