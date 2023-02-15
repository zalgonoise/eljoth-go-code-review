package config

type portSetting int

// Apply sets the option on the input Config `c`
func (p portSetting) Apply(c *Config) {
	c.Port = (int)(p)
}

// Port defines the HTTP port for the server
func Port(port int) Option {
	return (portSetting)(port)
}
