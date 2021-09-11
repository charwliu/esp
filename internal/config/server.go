package config

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/pkg/fs"
)

// DetachServer tests if server should detach from console (daemon mode).
func (c *Config) DetachServer() bool {
	return c.options.DetachServer
}

// HttpHost returns the built-in HTTP server host name or IP address (empty for all interfaces).
func (c *Config) HttpHost() string {
	if c.options.HttpHost == "" {
		if c.fiber.Network == fiber.NetworkTCP6 {
			return "[::1]"
		} else if c.fiber.Network == fiber.NetworkTCP4 {
			return "0.0.0.0"
		} else {
			return ""
		}
	}

	return c.options.HttpHost
}

func (c *Config) CertFile() string  {
	if c.options.CertFile == "" {
		c.options.CertFile = "server.pem"
	}
	return c.options.CertFile
}

func (c *Config) KeyFile() string  {
	if c.options.KeyFile == "" {
		c.options.KeyFile = "server.key"
	}
	return c.options.KeyFile
}




// HttpPort returns the built-in HTTP server port.
func (c *Config) HttpPort() int {
	if c.options.HttpPort == 0 {
		return 2342
	}

	return c.options.HttpPort
}

// HttpMode returns the server mode.
func (c *Config) HttpMode() string {
	if c.options.HttpMode == "" {
		if c.Debug() {
			return "debug"
		}

		return "release"
	}

	return c.options.HttpMode
}

// HttpCompression returns the http compression method (none or gzip).
func (c *Config) HttpCompression() string {
	return strings.ToLower(strings.TrimSpace(c.options.HttpCompression))
}

// TemplatesPath returns the server templates path.
func (c *Config) TemplatesPath() string {
	return filepath.Join(c.AssetsPath(), "templates")
}

// TemplateExists tests if a template with the given name exists (e.g. index.tmpl).
func (c *Config) TemplateExists(name string) bool {
	return fs.FileExists(filepath.Join(c.TemplatesPath(), name))
}

// TemplateName returns the name of the default template (e.g. index.tmpl).
func (c *Config) TemplateName() string {
	if s := c.Settings(); s != nil {
		if c.TemplateExists(s.Templates.Default) {
			return s.Templates.Default
		}
	}

	return "index.html"
}

// StaticPath returns the static asset's path.
func (c *Config) StaticPath() string {
	return filepath.Join(c.AssetsPath(), "static")
}

// BuildPath returns the static build path.
func (c *Config) BuildPath() string {
	return filepath.Join(c.AssetsPath(), "build")
}

// ImgPath returns the static image path.
func (c *Config) ImgPath() string {
	return filepath.Join(c.AssetsPath(), "img")
}
