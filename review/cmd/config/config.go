package config

import "github.com/zalgonoise/eljoth-go-code-review/coupon_service/internal/api"

type Option interface {
	Apply(c *Config)
}

type Config struct {
	Port int
}

func New(opts ...Option) *Config {
	return (&Config{Port: api.DefaultPort}).Apply(opts...)
}

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

func (c *Config) Apply(opts ...Option) *Config {
	for _, option := range opts {
		if option != nil {
			option.Apply(c)
		}
	}
	return c
}
