package cmd

import (
	"fmt"
	"github.com/seed95/forward-proxy/internal"
	"github.com/seed95/forward-proxy/internal/handler"
	"github.com/seed95/forward-proxy/internal/repo/cache"
	"github.com/seed95/forward-proxy/internal/repo/statistical"
	"github.com/seed95/forward-proxy/internal/service"
	"github.com/seed95/forward-proxy/pkg/log"
	"github.com/seed95/forward-proxy/pkg/log/keyval"
	nativeLog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
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
	conf.StdLogLevel = internal.DefaultStdLogLevel
	conf.ExpirationTime = internal.DefaultRedisExpiration
	conf.PostgresConfig = internal.PostgresConfig{
		Address:      "localhost",
		Port:         "15432",
		Username:     "postgres",
		Password:     "pass",
		DatabaseName: "postgres",
	}

	rootCmd.Flags().StringVar(&conf.ProxyPort, "proxy-port", "", "Proxy port is used to received request from this port")
	_ = rootCmd.MarkFlagRequired("proxy-port")
	fmt.Println("proxy port", conf.ProxyPort)
	rootCmd.Flags().StringVar(&conf.RedisConfig.Address, "redis", "localhost", "Proxy port is used to received request from this port. Default is 2121")
}

func Execute() error {
	return rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	// Setup cache repo
	cacheRepo := cache.New(&conf.RedisConfig)

	// Setup statistical repo
	statsRepo := statistical.New(&conf.PostgresConfig)
	//if err != nil {
	//	// TODO handle error
	//	fmt.Println("service error", err)
	//}

	srv := service.New(cacheRepo, statsRepo)

	// Setup HTTP server
	httpHandler := handler.NewHttpHandler(srv)
	nativeLog.Printf(fmt.Sprintf("[INFO] Running HTTP server on port %s", conf.ProxyPort))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", conf.ProxyPort), httpHandler.Route()); err != nil {
		log.Panic(fmt.Sprintf("[ERROR] Failed to bind HTTP server on port %s", conf.ProxyPort), keyval.Error(err))
	}

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-killSig
	fmt.Println("sajad")
}
