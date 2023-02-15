package main

import (
	"fmt"
	"time"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/config"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/repository/memdb"
	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/service"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	svc := service.New(repo)
	本 := api.New(cfg.API, svc)
	本.Start()
	fmt.Println("Starting Coupon service server")
	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Coupon service server alive for a year, closing")
	本.Close()
}
