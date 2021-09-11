package config

type AccessLoggerConfig struct {
	*LoggerConfig
	Enabled bool `yaml:"Enabled" json:"enabled"`
}


func (c *Config) AccessLogger() *AccessLoggerConfig {
	if c.options.AccessLoggerConfig == nil {
		c.options.AccessLoggerConfig = &AccessLoggerConfig{
			LoggerConfig: &LoggerConfig{
				Type:        "console",
				Environment: "",
				Filename:    "access.log",
				MaxSize:     500,
				MaxAge:      28,
				MaxBackups:  3,
				LocalTime:   false,
				Compress:    false,
			},
			Enabled: true,
		}
	}
	return c.options.AccessLoggerConfig
}
