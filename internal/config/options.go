package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"go.vixal.xyz/esp/pkg/fs"
)

// Database drivers (sql dialects).
const (
	MySQL     = "mysql"
	MariaDB   = "mariadb"
	SQLite    = "sqlite3"
	Postgres  = "postgres"
	SQLServer = "sqlserver"
)

// Options provides a struct in which application configuration is stored.
// Application code must use functions to get config options
type Options struct {
	Name                      string              `json:"-"`
	Version                   string              `json:"-"`
	Copyright                 string              `json:"-"`
	Debug                     bool                `yaml:"Debug" json:"Debug" flag:"debug"`
	Test                      bool                `yaml:"-" json:"Test,omitempty" flag:"test"`
	Demo                      bool                `yaml:"Demo" json:"-" flag:"demo"`
	Sponsor                   bool                `yaml:"-" json:"-" flag:"sponsor"`
	Public                    bool                `yaml:"Public" json:"-" flag:"public"`
	ReadOnly                  bool                `yaml:"ReadOnly" json:"ReadOnly" flag:"read-only"`
	Experimental              bool                `yaml:"Experimental" json:"Experimental" flag:"experimental"`
	ConfigPath                string              `yaml:"ConfigPath" json:"-" flag:"config-path"`
	ConfigFile                string              `json:"-"`
	AdminPassword             string              `yaml:"AdminPassword" json:"-" flag:"admin-password"`
	StoragePath               string              `yaml:"StoragePath" json:"-" flag:"storage-path"`
	TempPath                  string              `yaml:"TempPath" json:"-" flag:"temp-path"`
	BackupPath                string              `yaml:"BackupPath" json:"-" flag:"backup-path"`
	AssetsPath                string              `yaml:"AssetsPath" json:"-" flag:"assets-path"`
	CachePath                 string              `yaml:"CachePath" json:"-" flag:"cache-path"`
	Workers                   int                 `yaml:"Workers" json:"Workers" flag:"workers"`
	WakeupInterval            int                 `yaml:"WakeupInterval" json:"WakeupInterval" flag:"wakeup-interval"`
	AutoIndex                 int                 `yaml:"AutoIndex" json:"AutoIndex" flag:"auto-index"`
	AutoImport                int                 `yaml:"AutoImport" json:"AutoImport" flag:"auto-import"`
	DisableBackups            bool                `yaml:"DisableBackups" json:"DisableBackups" flag:"disable-backups"`
	DisableWebDAV             bool                `yaml:"DisableWebDAV" json:"DisableWebDAV" flag:"disable-webdav"`
	DisableSettings           bool                `yaml:"DisableSettings" json:"-" flag:"disable-settings"`
	DisablePlaces             bool                `yaml:"DisablePlaces" json:"DisablePlaces" flag:"disable-places"`
	LogLevel                  string              `yaml:"LogLevel" json:"-" flag:"log-level"`
	LogFilename               string              `yaml:"LogFilename" json:"-" flag:"log-filename"`
	PIDFilename               string              `yaml:"PIDFilename" json:"-" flag:"pid-filename"`
	CdnUrl                    string              `yaml:"CdnUrl" json:"CdnUrl" flag:"cdn-url"`
	SiteUrl                   string              `yaml:"SiteUrl" json:"SiteUrl" flag:"site-url"`
	SitePreview               string              `yaml:"SitePreview" json:"SitePreview" flag:"site-preview"`
	SiteTitle                 string              `yaml:"SiteTitle" json:"SiteTitle" flag:"site-title"`
	SiteCaption               string              `yaml:"SiteCaption" json:"SiteCaption" flag:"site-caption"`
	SiteDescription           string              `yaml:"SiteDescription" json:"SiteDescription" flag:"site-description"`
	SiteAuthor                string              `yaml:"SiteAuthor" json:"SiteAuthor" flag:"site-author"`
	DatabaseDriver            string              `yaml:"DatabaseDriver" json:"-" flag:"database-driver"`
	DatabaseDsn               string              `yaml:"DatabaseDsn" json:"-" flag:"database-dsn"`
	DatabaseServer            string              `yaml:"DatabaseServer" json:"-" flag:"database-server"`
	DatabaseName              string              `yaml:"DatabaseName" json:"-" flag:"database-name"`
	DatabaseUser              string              `yaml:"DatabaseUser" json:"-" flag:"database-user"`
	DatabasePassword          string              `yaml:"DatabasePassword" json:"-" flag:"database-password"`
	DatabaseConns             int                 `yaml:"DatabaseConns" json:"-" flag:"database-conns"`
	DatabaseConnsIdle         int                 `yaml:"DatabaseConnsIdle" json:"-" flag:"database-conns-idle"`
	HttpHost                  string              `yaml:"HttpHost" json:"-" flag:"http-host"`
	HttpPort                  int                 `yaml:"HttpPort" json:"-" flag:"http-port"`
	HttpMode                  string              `yaml:"HttpMode" json:"-" flag:"http-mode"`
	HttpCompression           string              `yaml:"HttpCompression" json:"-" flag:"http-compression"`
	DetachServer              bool                `yaml:"DetachServer" json:"-" flag:"detach-server"`
	DownloadToken             string              `yaml:"DownloadToken" json:"-" flag:"download-token"`
	PreviewToken              string              `yaml:"PreviewToken" json:"-" flag:"preview-token"`
	SessionConfig             *SessionConfig      `yaml:"Session"`
	FiberViews                *FiberViewConfig    `yaml:"FiberViews"`
	HasherDriver              *HasherConfig       `yaml:"HasherDriver" json:"hasher_driver"`
	AccessLoggerConfig        *AccessLoggerConfig `yaml:"AccessLogger"`
	ForceHTTPS                bool
	ForceTrailingSlash        bool
	EnableHSTS                bool
	HSTSMaxAge                int
	HSTSIncludeSubdomains     bool
	HSTSPreload               bool
	SuppressWWW               bool
	FiberRecoverDisabled      bool
	FiberCacheEnabled         bool
	FiberCacheExpiration      time.Duration
	FiberCacheCacheControl    bool
	FiberCompressEnabled      bool
	FiberCompressLevel        int
	FiberCORSEnabled          bool
	FiberCORSAllowOrigins     string
	FiberCORSAllowMethods     string
	FiberCORSAllowHeaders     string
	FiberCORSAllowCredentials bool
	FiberCORSExposeHeaders    string
	FiberCORSMaxAge           int
	FiberCSRFEnabled          bool
	FiberCSRFKeyLookup        string
	FiberCSRFCookieName       string
	FiberCSRFCookieSameSite   string
	FiberCSRFExpiration       time.Duration
	FiberCSRFContextKey       string
	FiberETagEnabled          bool
	FiberETagWeak             bool
	FiberExpvarEnabled        bool
	FiberFaviconEnabled       bool
	FiberFaviconFile          string
	FiberFaviconCacheControl  string
	FiberLimiterEnabled       bool
	FiberLimiterMax           int
	FiberLimiterExpiration    time.Duration
	FiberMonitorEnabled       bool
	FiberPprofEnabled         bool
	FiberRequestIDEnabled     bool
	FiberRequestIDHeader      string
	FiberRequestIDContextKey  string
	FiberConfig               *FiberConfig `yaml:"Fiber"`
}

type FiberConfig struct {
	Prefork                   bool          `yaml:"Prefork" json:"prefork"`
	ServerHeader              string        `yaml:"ServerHeader" json:"server_header"`
	StrictRouting             bool          `yaml:"StrictRouting" json:"strict_routing"`
	CaseSensitive             bool          `yaml:"case_sensitive" json:"case_sensitive"`
	Immutable                 bool          `yaml:"immutable" json:"immutable"`
	UnescapePath              bool          `yaml:"UnescapePath" json:"unescape_path"`
	ETag                      bool          `yaml:"ETag" json:"etag"`
	BodyLimit                 int           `yaml:"BodyLimit" json:"body_limit"`
	Concurrency               int           `yaml:"Concurrency" json:"concurrency"`
	ReadTimeout               time.Duration `yaml:"ReadTimeout" json:"read_timeout"`
	WriteTimeout              time.Duration `yaml:"WriteTimeout" json:"write_timeout"`
	IdleTimeout               time.Duration `yaml:"IdleTimeout" json:"idle_timeout"`
	ReadBufferSize            int           `yaml:"ReadBufferSize" json:"read_buffer_size"`
	WriteBufferSize           int           `yaml:"WriteBufferSize" json:"write_buffer_size"`
	CompressedFileSuffix      string        `yaml:"CompressedFileSuffix" json:"compressed_file_suffix"`
	ProxyHeader               string        `yaml:"ProxyHeader" json:"proxy_header"`
	GETOnly                   bool          `yaml:"GETOnly" json:"get_only"`
	DisableKeepalive          bool          `yaml:"DisableKeepalive" json:"disable_keepalive"`
	DisableDefaultDate        bool          `yaml:"DisableDefaultDate" json:"disable_default_date"`
	DisableDefaultContentType bool          `yaml:"DisableDefaultContentType" json:"disable_default_content_type"`
	DisableHeaderNormalizing  bool          `yaml:"DisableHeaderNormalizing" json:"disable_header_normalizing"`
	DisableStartupMessage     bool          `yaml:"DisableStartupMessage" json:"disable_startup_message"`
	ReduceMemoryUsage         bool          `yaml:"ReduceMemoryUsage" json:"reduce_memory_usage"`
	ViewsLayout               string        `yaml:"ViewsLayout"`
	EnableTrustedProxyCheck   bool          `yaml:"EnableTrustedProxyCheck" json:"enable_trusted_proxy_check"`
	TrustedProxies            []string      `yaml:"TrustedProxies" json:"trusted_proxies"`
}

type FiberViewConfig struct {
	Engine    string `yaml:"Engine"`
	Layout    string `yaml:"Layout"`
	Directory string `yaml:"Directory"`
	DelimsL   string `yaml:"DelimsL"`
	DelimsR   string `yaml:"DelimsR"`
	Extension string `yaml:"Extension"`
	Debug     bool   `yaml:"Debug"`
	Reload    bool   `yaml:"Reload"`
}

type HasherConfig struct {
	Driver      string `yaml:"Driver", json:"driver"`
	Rounds      int    `yaml:"Rounds"`
	Memory      uint32 `yaml:"Memory"`
	SaltLength  uint32 `yaml:"SaltLength"`
	KeyLength   uint32 `yaml:"KeyLength"`
	Iterations  uint32 `yaml:"Iterations"`
	Parallelism uint8  `yaml:"Parallelism"`
}

type SessionConfig struct {
	Provider       string        `yaml:"Provider"`
	KeyLookup      string        `yaml:"KeyLookup" json:"key_lookup"`
	Database       string        `yaml:"Database"`
	Table          string        `yaml:"Table"`
	Host           string        `yaml:"Host"`
	Port           int           `yaml:"Port"`
	Username       string        `yaml:"Username"`
	Password       string        `yaml:"Password"`
	Expiration     time.Duration `yaml:"Expiration"`
	GCInterval     time.Duration `yaml:"GCInterval"`
	KeyPrefix      string        `yaml:"KeyPrefix"`
	CookieDomain   string        `yaml:"CookieDomain"`
	CookieSameSite string        `yaml:"CookieSameSite"`
	CookieSecure   bool          `yaml:"CookieSecure"`
	CookiePath     string        `yaml:"CookiePath"`
}

// NewOptions creates a new configuration entity by using two methods:
//
// 1. Load: This will initialize options from a yaml config file.
//
// 2. SetContext: Which comes after Load and overrides
//    any previous options giving an option two override file configs through the CLI.
func NewOptions(ctx *cli.Context) *Options {
	c := &Options{}

	if ctx == nil {
		return c
	}

	c.Name = ctx.App.Name
	c.Copyright = ctx.App.Copyright
	c.Version = ctx.App.Version
	c.ConfigFile = fs.Abs(ctx.String("config-file"))

	if err := c.Load(c.ConfigFile); err != nil {
		log.Debug(err)
	}

	if err := c.SetContext(ctx); err != nil {
		log.Error(err)
	}

	return c
}

// expandFilenames converts path in config to absolute path
func (c *Options) expandFilenames() {
	c.ConfigPath = fs.Abs(c.ConfigPath)
	c.BackupPath = fs.Abs(c.BackupPath)
	c.AssetsPath = fs.Abs(c.AssetsPath)
	c.CachePath = fs.Abs(c.CachePath)
	c.TempPath = fs.Abs(c.TempPath)
	c.PIDFilename = fs.Abs(c.PIDFilename)
	c.LogFilename = fs.Abs(c.LogFilename)
}

// Load uses a yaml config file to initiate the configuration entity.
func (c *Options) Load(fileName string) error {
	if fileName == "" {
		return nil
	}

	if !fs.FileExists(fileName) {
		return errors.New(fmt.Sprintf("config: %s not found", fileName))
	}

	yamlConfig, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlConfig, c)
}

// SetContext uses options from the CLI to setup configuration overrides
// for the entity.
func (c *Options) SetContext(ctx *cli.Context) error {
	v := reflect.ValueOf(c).Elem()

	// Iterate through all config fields.
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)

		tagValue := v.Type().Field(i).Tag.Get("flag")

		// Automatically assign options to field with "flag" tag.
		if tagValue != "" {
			switch t := fieldValue.Interface().(type) {
			case int, int64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Int64(tagValue)
					fieldValue.SetInt(f)
				} else if ctx.IsSet(tagValue) || fieldValue.Int() == 0 {
					f := ctx.Int64(tagValue)
					fieldValue.SetInt(f)
				}
			case uint, uint64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Uint64(tagValue)
					fieldValue.SetUint(f)
				} else if ctx.IsSet(tagValue) || fieldValue.Uint() == 0 {
					f := ctx.Uint64(tagValue)
					fieldValue.SetUint(f)
				}
			case string:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				} else if ctx.IsSet(tagValue) || fieldValue.String() == "" {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				}
			case bool:
				if ctx.IsSet(tagValue) {
					f := ctx.Bool(tagValue)
					fieldValue.SetBool(f)
				} else if ctx.IsSet(tagValue) {
					f := ctx.Bool(tagValue)
					fieldValue.SetBool(f)
				}
			default:
				log.Warnf("can't assign value of type %s from cli flag %s", t, tagValue)
			}
		}
	}

	return nil
}
