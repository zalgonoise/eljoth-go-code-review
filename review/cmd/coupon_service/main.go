package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/cmd/config"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository/memdb"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/service"
)

func main() {
	closeSignal := make(chan os.Signal, 1)
	signal.Notify(closeSignal, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

	// use a simple standard library logger
	logger := log.New(os.Stderr, "coupon-service", log.Default().Flags())

	cfg := config.ParseFlags()
	svc := service.New(memdb.New())
	server := api.New(cfg.Port, svc)

	logger.Println("starting Coupon service server")
	if err := server.Start(); err != nil {
		logger.Fatalf("failed to start HTTP server: %v\n", err)
		os.Exit(1)
	}
	<-closeSignal
	logger.Println("received an interrupt / terminate signal; exiting")
	if err := server.Close(); err != nil {
		logger.Fatalf("failed to stop HTTP server: %v\n", err)
		os.Exit(1)
	}
}
