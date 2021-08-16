package config

type AccessLoggerConfig struct {
	Enabled     bool   `yaml:"Enabled"`
	Type        string `yaml:"ClientType"`
	Environment string `yaml:"Environment"`
	Filename    string `yaml:"Filename"`
	MaxSize     int    `yaml:"MaxSize"`
	MaxAge      int    `yaml:"MaxAge"`
	MaxBackups  int    `yaml:"MaxBackups"`
	LocalTime   bool   `yaml:"LocalTime"`
	Compress    bool   `yaml:"Compress"`
}
