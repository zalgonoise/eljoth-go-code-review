package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api"
)

func ParseFlags() *Config {
	httpPort := flag.Int("port", api.DefaultPort, "port to use for the HTTP server")
	// add new CLI flags as config expands

	flag.Parse()

	conf := New(
		Port(*httpPort),
		// configure new options as config expands
	)

	return conf.Merge(ParseOSEnv())
}

func ParseOSEnv() *Config {
	portStr := os.Getenv("COUPON_SERVICE_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// an invalid port as an OS environment variable will be ignored
		return nil
	}

	return &Config{
		Port: port,
	}
}
