package internal

import (
	"time"
)

const (
	DefaultStdLogLevel     = 1
	DefaultRedisExpiration = 5 * time.Minute
)

type RedisConfig struct {
	Address        string
	Password       string
	ExpirationTime time.Duration
}

type PostgresConfig struct {
	Address      string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

type Config struct {
	StdLogLevel int    // Standard core log level
	ProxyPort   string // Proxy port is used to receive requests from this port to forward.
	RedisConfig
	PostgresConfig
}
