package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/discount"

	"github.com/gin-gonic/gin"
)

const defaultPort int = 8080

var (
	ErrNilServer = errors.New("server was not initialized properly and is nil")
)

type Config struct {
	Host string
	Port int
}

type API struct {
	srv *http.Server
	svc discount.Service
}

func New(port int, svc discount.Service) *API {
	if port <= 0 {
		port = defaultPort
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	api := &API{
		svc: svc,
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
		},
	}

	apiGroup := r.Group("/api")
	apiGroup.POST("/apply", api.Apply)
	apiGroup.POST("/create", api.Create)
	apiGroup.GET("/coupons", api.Get)

	return api
}

func (a *API) Start() error {
	if a.srv == nil {
		return ErrNilServer
	}
	return a.srv.ListenAndServe()
}

func (a *API) Close() error {
	if a.srv == nil {
		return ErrNilServer
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: implement a graceful shutdown for the CouponService
	// if err := a.svc.Close(); err != nil { return err }
	return a.srv.Shutdown(ctx)
}
