package cmd

import (
	"fmt"
	nativeLog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/seed95/forward-proxy/internal"
	"github.com/seed95/forward-proxy/internal/handler"
	"github.com/seed95/forward-proxy/internal/helper"
	"github.com/seed95/forward-proxy/internal/repo/cache"
	"github.com/seed95/forward-proxy/internal/repo/statistical"
	"github.com/seed95/forward-proxy/internal/service"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"

	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
)

var (
	rootCmd = &cobra.Command{
		Use:   "forward proxy",
		Short: "forward proxy is a simplify forward received request to target",
		Run:   run,
	}

	conf internal.Config
)

func init() {
	conf.ExpirationTime = internal.DefaultRedisExpiration

	log.Init(internal.DefaultStdLogLevel)

	rootCmd.Flags().StringVar(&conf.ProxyPort, "proxy-port", "", "Proxy port is used to received request from this port")
	rootCmd.Flags().StringVar(&conf.RedisConfig.Address, "redis-address", "", "Redis address.")
	rootCmd.Flags().StringVar(&conf.PostgresConfig.Address, "postgres-address", "", "Postgres address.")
	rootCmd.Flags().StringVar(&conf.PostgresConfig.Port, "postgres-port", "", "Postgres port.")
	rootCmd.Flags().StringVar(&conf.PostgresConfig.Username, "postgres-username", "", "Postgres username.")
	rootCmd.Flags().StringVar(&conf.PostgresConfig.Password, "postgres-password", "", "Postgres password.")
	rootCmd.Flags().StringVar(&conf.PostgresConfig.DatabaseName, "postgres-dbname", "", "Postgres database name.")

	_ = rootCmd.MarkFlagRequired("proxy-port")
	_ = rootCmd.MarkFlagRequired("redis-address")
	_ = rootCmd.MarkFlagRequired("postgres-address")
	_ = rootCmd.MarkFlagRequired("postgres-port")
	_ = rootCmd.MarkFlagRequired("postgres-username")
	_ = rootCmd.MarkFlagRequired("postgres-password")
	_ = rootCmd.MarkFlagRequired("postgres-dbname")
}

func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	// Setup redis database
	rdb, err := helper.NewRedisDatabase(conf.RedisConfig)
	if err != nil {
		log.Panic("redis connection", keyval.Error(err))
	}

	// Setup postgres gorm database
	db, err := helper.NewPostgresGormDB(conf.PostgresConfig)
	if err != nil {
		log.Panic("postgres connection", keyval.Error(err))
	}

	// Setup cache repo
	cacheRepo := cache.New(rdb, conf.ExpirationTime)

	// Setup statistical repo
	statsRepo := statistical.New(db)

	// Setup service
	srv := service.New(cacheRepo, statsRepo)

	// Setup http limiter
	limiter := rate.NewLimiter(rate.Limit(internal.DefaultLimiterTps), internal.DefaultLimiterBurstSize)

	// Setup HTTP server
	httpHandler := handler.NewHttpHandler(srv, limiter)
	s := http.Server{
		Addr:    fmt.Sprintf(":%s", conf.ProxyPort),
		Handler: handler.RecoverMiddleware(handler.LogMiddleware(httpHandler.Serve)),
	}
	nativeLog.Printf(fmt.Sprintf("[INFO] Running HTTP server on port %s", conf.ProxyPort))
	if err := s.ListenAndServe(); err != nil {
		log.Panic(fmt.Sprintf("[ERROR] Failed to bind HTTP server on port %s", conf.ProxyPort), keyval.Error(err))
	}

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-killSig
}
