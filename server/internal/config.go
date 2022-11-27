package internal

import (
	"github.com/seed95/forward-proxy/pkg/log/zap"
	"time"
)

const (
	DefaultStdLogLevel      = zap.InfoLevel
	DefaultRedisExpiration  = 5 * time.Minute
	DefaultLimiterTps       = 10
	DefaultLimiterBurstSize = 10
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
	ProxyPort string // Proxy port is used to receive requests from this port to forward.
	RedisConfig
	PostgresConfig
}
