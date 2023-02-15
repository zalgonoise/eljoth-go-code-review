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

// DefaultPort defines the default port for the CouponsService if unset
const DefaultPort int = 8080

var (
	ErrNilServer = errors.New("server was not initialized properly and is nil")
)

// API describes the Coupons' HTTP server, configured with a discount.Service and
// HTTP handlers.
type API struct {
	srv *http.Server
	svc discount.Service
}

// New creates a new Coupons API, based off of the input port `port` and
// discount.Service `svc`. It initializes the HTTP server with its routes configured
// so it is ready to start.
func New(port int, svc discount.Service) *API {
	if port <= 0 {
		port = DefaultPort
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	api := &API{
		svc: svc,
	}
	r.POST("/api/apply", api.Apply())
	r.POST("/api/create", api.Create())
	r.GET("/api/coupons", api.List())
	r.GET("/api/coupon", api.Get())

	api.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	return api
}

// Start runs the server, returning an error. This is a blocking call.
func (a *API) Start() error {
	if a.srv == nil {
		return ErrNilServer
	}
	return a.srv.ListenAndServe()
}

// Close implements the io.Closer interface, gracefully shutting down the
// HTTP server; then the discount.Service. Returns an error if raised
func (a *API) Close() error {
	if a.srv == nil {
		return ErrNilServer
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		if serviceErr := a.srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("%w -- service shutdown error: %v", err, serviceErr)
		}
		return err
	}
	return a.svc.Close()
}
