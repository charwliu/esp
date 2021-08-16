package config



func (c *Config) FiberViews() *FiberViewConfig {
	if c.options.FiberViews == nil {
		c.options.FiberViews = &FiberViewConfig{
			Engine:    "html",
			Layout:    "embed",
			Directory: "resources/views",
			DelimsL:   "{{",
			DelimsR:   "}}",
			Extension: "html",
			Debug:     false,
			Reload:    true,
		}
	}
	return c.options.FiberViews
}

