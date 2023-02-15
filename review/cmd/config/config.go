package config

import "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api"

// Option describes setter types for a Config
type Option interface {
	// Apply sets the option on the input Config `c`
	Apply(c *Config)
}

// Config structures the setup of the coupon_service app, according to the caller's needs
type Config struct {
	Port int
}

// New initializes a new config with default settings, and then iterates through
// all input ConfigOption `opts` applying them to the Config, which is returned
// to the caller
func New(opts ...Option) *Config {
	return (&Config{Port: api.DefaultPort}).Apply(opts...)
}

// Merge combines Configs `c` with `input`, returning a merged version
// of the two
//
// All set elements in `input` will be applied to `c`, and the unset elements
// will be ignored (keeps `c`'s data)
func (c *Config) Merge(input *Config) *Config {
	if input == nil {
		return c
	}
	if input.Port != 0 {
		c.Port = input.Port
	}
	// add similar checks as config expands
	return c
}

// Apply implements the Option interface
//
// It allows applying new options on top of an already existing config
func (c *Config) Apply(opts ...Option) *Config {
	for _, option := range opts {
		if option != nil {
			option.Apply(c)
		}
	}
	return c
}
