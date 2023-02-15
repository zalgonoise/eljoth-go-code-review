package config

type portSetting int

func (p portSetting) Apply(c *Config) {
	c.Port = (int)(p)
}

func Port(port int) Option {
	return (portSetting)(port)
}
