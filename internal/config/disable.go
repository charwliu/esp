package config

// DisableBackups tests if photo and album metadata backups should be disabled.
func (c *Config) DisableBackups() bool {
	return c.options.DisableBackups
}

// DisableWebDAV tests if the built-in WebDAV server should be disabled.
func (c *Config) DisableWebDAV() bool {
	if c.ReadOnly() || c.Demo() {
		return true
	}

	return c.options.DisableWebDAV
}

// DisableSettings tests if users should not be allowed to change settings.
func (c *Config) DisableSettings() bool {
	return c.options.DisableSettings
}

// DisablePlaces tests if geocoding and maps should be disabled.
func (c *Config) DisablePlaces() bool {
	return c.options.DisablePlaces
}
