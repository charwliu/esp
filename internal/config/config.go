package config

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memcache"
	"github.com/gofiber/storage/mysql"
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/storage/sqlite3"
	"github.com/gofiber/template/ace"
	"github.com/gofiber/template/amber"
	"github.com/gofiber/template/django"
	"github.com/gofiber/template/handlebars"
	"github.com/gofiber/template/html"
	"github.com/gofiber/template/jet"
	"github.com/gofiber/template/mustache"
	"github.com/gofiber/template/pug"
	"github.com/jameskeane/bcrypt"
	"github.com/klauspost/cpuid/v2"

	"github.com/alexedwards/argon2id"
	hashing "github.com/thomasvvugt/fiber-hashing"
	argon_driver "github.com/thomasvvugt/fiber-hashing/driver/argon2id"
	bcrypt_driver "github.com/thomasvvugt/fiber-hashing/driver/bcrypt"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/internal/mutex"
	"go.vixal.xyz/esp/pkg/fs"
	"go.vixal.xyz/esp/pkg/rnd"
	"go.vixal.xyz/esp/pkg/storage/postgres"
	"go.vixal.xyz/esp/pkg/txt"
)

var log = event.Log.Sugar()

var once sync.Once

const ApiUri = "/api/v1"
const StaticUri = "/static"

// Config holds database, cache and all parameters of esp
type Config struct {
	once         sync.Once
	db           *gorm.DB
	options      *Options
	settings     *Settings
	fiber        *fiber.Config
	token        string
	serial       string
	logger       *LoggerConfig
	errorHandler fiber.ErrorHandler
}

type LoggerConfig struct {
	// Type determines whether zap will be initialized as a file logger or,
	// by default, as a console logger.
	Type string `json:"type"`

	// Environment determines whether zap will be initialized using a production
	// or a development logger.
	Environment string `json:"environment"`

	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.
	Filename string `json:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated.
	MaxSize int `json:"max_size"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"max_age"`

	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int `json:"max_backups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.
	LocalTime bool `json:"local_time"`

	// Compress determines if the rotated log files should be compressed
	// using gzip.
	Compress bool `json:"compress"`

	Verbose bool `json:"verbose"`
}

//
//func New() *Config {
//	config := &Config{
//		Viper: viper.New(),
//	}
//
//	// Set default configurations
//	config.setDefaults()
//
//	// Select the .env file
//	config.SetConfigName(".env")
//	config.SetConfigType("dotenv")
//	config.AddConfigPath(".")
//
//	// Automatically refresh environment variables
//	config.AutomaticEnv()
//
//	// Read configuration
//	if err := config.ReadInConfig(); err != nil {
//		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
//			fmt.Println("failed to read configuration:", err.Error())
//			os.Exit(1)
//		}
//	}
//
//	config.SetErrorHandler(defaultErrorHandler)
//
//	// TODO: Logger (Maybe a different zap object)
//	config.setLoggerConfig()
//
//	// TODO: Add APP_KEY generation
//
//	// TODO: Write changes to configuration file
//
//	// Set Fiber configurations
//	config.initFiber()
//
//	return config
//}
//

func (c *Config) SetErrorHandler(errorHandler fiber.ErrorHandler) {
	c.errorHandler = errorHandler
}

func (c *Config) setDefaults() {
	//// Set default App configuration
	//c.SetDefault("APP_ADDR", ":8080")
	//c.SetDefault("APP_ENV", "local")
	//execPath, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//workDir := filepath.Dir(execPath)
	//if workDir, err := filepath.EvalSymlinks(workDir); err == nil {
	//	c.SetDefault("APP_WORKDIR", workDir)
	//} else {
	//	c.SetDefault("APP_WORKDIR", "./")
	//}
	//
	//// Set default database configuration
	//c.SetDefault("DB_DRIVER", "mysql")
	//c.SetDefault("DB_HOST", "localhost")
	//c.SetDefault("DB_USERNAME", "esp")
	//c.SetDefault("DB_PASSWORD", "password")
	//c.SetDefault("DB_PORT", 3306)
	//c.SetDefault("DB_DATABASE", "esp")
	//
	//// Set default hasher configuration
	//c.SetDefault("HASHER_DRIVER", "argon2id")
	//c.SetDefault("HASHER_MEMORY", 131072)
	//c.SetDefault("HASHER_ITERATIONS", 4)
	//c.SetDefault("HASHER_PARALLELISM", 4)
	//c.SetDefault("HASHER_SALTLENGTH", 16)
	//c.SetDefault("HASHER_KEYLENGTH", 32)
	//c.SetDefault("HASHER_ROUNDS", bcrypt.DefaultRounds)
	//
	//// Set default session configuration
	//c.SetDefault("SESSION_PROVIDER", "mysql")
	//c.SetDefault("SESSION_KEYPREFIX", "session")
	//c.SetDefault("SESSION_HOST", "localhost")
	//c.SetDefault("SESSION_PORT", 3306)
	//c.SetDefault("SESSION_USERNAME", "fiber")
	//c.SetDefault("SESSION_PASSWORD", "secret")
	//c.SetDefault("SESSION_DATABASE", "esp")
	//c.SetDefault("SESSION_TABLENAME", "sessions")
	//c.SetDefault("SESSION_LOOKUP", "cookie:session_id")
	//c.SetDefault("SESSION_DOMAIN", "")
	//c.SetDefault("SESSION_SAMESITE", "Lax")
	//c.SetDefault("SESSION_EXPIRATION", "12h")
	//c.SetDefault("SESSION_SECURE", false)
	//c.SetDefault("SESSION_GCINTERVAL", "1m")
	//
	//// Set default Fiber configuration
	//c.SetDefault("FIBER_PREFORK", false)
	//c.SetDefault("FIBER_SERVERHEADER", "")
	//c.SetDefault("FIBER_STRICTROUTING", false)
	//c.SetDefault("FIBER_CASESENSITIVE", false)
	//c.SetDefault("FIBER_IMMUTABLE", false)
	//c.SetDefault("FIBER_UNESCAPEPATH", false)
	//c.SetDefault("FIBER_ETAG", false)
	//c.SetDefault("FIBER_BODYLIMIT", 4194304)
	//c.SetDefault("FIBER_CONCURRENCY", 262144)
	//c.SetDefault("FIBER_VIEWS", "html")
	//c.SetDefault("FIBER_VIEWS_DIRECTORY", "resources/views")
	//c.SetDefault("FIBER_VIEWS_RELOAD", false)
	//c.SetDefault("FIBER_VIEWS_DEBUG", false)
	//c.SetDefault("FIBER_VIEWS_LAYOUT", "embed")
	//c.SetDefault("FIBER_VIEWS_DELIMS_L", "{{")
	//c.SetDefault("FIBER_VIEWS_DELIMS_R", "}}")
	//c.SetDefault("FIBER_READTIMEOUT", 0)
	//c.SetDefault("FIBER_WRITETIMEOUT", 0)
	//c.SetDefault("FIBER_IDLETIMEOUT", 0)
	//c.SetDefault("FIBER_READBUFFERSIZE", 4096)
	//c.SetDefault("FIBER_WRITEBUFFERSIZE", 4096)
	//c.SetDefault("FIBER_COMPRESSEDFILESUFFIX", ".fiber.gz")
	//c.SetDefault("FIBER_PROXYHEADER", "")
	//c.SetDefault("FIBER_GETONLY", false)
	//c.SetDefault("FIBER_DISABLEKEEPALIVE", false)
	//c.SetDefault("FIBER_DISABLEDEFAULTDATE", false)
	//c.SetDefault("FIBER_DISABLEDEFAULTCONTENTTYPE", false)
	//c.SetDefault("FIBER_DISABLEHEADERNORMALIZING", false)
	//c.SetDefault("FIBER_DISABLESTARTUPMESSAGE", false)
	//c.SetDefault("FIBER_REDUCEMEMORYUSAGE", false)
	//
	//// Set default Custom Access Logger middleware configuration
	//c.SetDefault("MW_ACCESS_LOGGER_ENABLED", true)
	//c.SetDefault("MW_ACCESS_LOGGER_TYPE", "console")
	//c.SetDefault("MW_ACCESS_LOGGER_FILENAME", "access.log")
	//c.SetDefault("MW_ACCESS_LOGGER_MAXSIZE", 500)
	//c.SetDefault("MW_ACCESS_LOGGER_MAXAGE", 28)
	//c.SetDefault("MW_ACCESS_LOGGER_MAXBACKUPS", 3)
	//c.SetDefault("MW_ACCESS_LOGGER_LOCALTIME", false)
	//c.SetDefault("MW_ACCESS_LOGGER_COMPRESS", false)
	//
	//// Set default Force HTTPS middleware configuration
	//c.SetDefault("MW_FORCE_HTTPS_ENABLED", false)
	//
	//// Set default Force trailing slash middleware configuration
	//c.SetDefault("MW_FORCE_TRAILING_SLASH_ENABLED", false)
	//
	//// Set default HSTS middleware configuration
	//c.SetDefault("MW_HSTS_ENABLED", false)
	//c.SetDefault("MW_HSTS_MAXAGE", 31536000)
	//c.SetDefault("MW_HSTS_INCLUDESUBDOMAINS", true)
	//c.SetDefault("MW_HSTS_PRELOAD", false)
	//
	//// Set default Suppress WWW middleware configuration
	//c.SetDefault("MW_SUPPRESS_WWW_ENABLED", true)
	//
	//// Set default Fiber Cache middleware configuration
	//c.SetDefault("MW_FIBER_CACHE_ENABLED", false)
	//c.SetDefault("MW_FIBER_CACHE_EXPIRATION", "1m")
	//c.SetDefault("MW_FIBER_CACHE_CACHECONTROL", false)
	//
	//// Set default Fiber Compress middleware configuration
	//c.SetDefault("MW_FIBER_COMPRESS_ENABLED", false)
	//c.SetDefault("MW_FIBER_COMPRESS_LEVEL", 0)
	//
	//// Set default Fiber CORS middleware configuration
	//c.SetDefault("MW_FIBER_CORS_ENABLED", false)
	//c.SetDefault("MW_FIBER_CORS_ALLOWORIGINS", "*")
	//c.SetDefault("MW_FIBER_CORS_ALLOWMETHODS", "GET,POST,HEAD,PUT,DELETE,PATCH")
	//c.SetDefault("MW_FIBER_CORS_ALLOWHEADERS", "")
	//c.SetDefault("MW_FIBER_CORS_ALLOWCREDENTIALS", false)
	//c.SetDefault("MW_FIBER_CORS_EXPOSEHEADERS", "")
	//c.SetDefault("MW_FIBER_CORS_MAXAGE", 0)
	//
	//// Set default Fiber CSRF middleware configuration
	//c.SetDefault("MW_FIBER_CSRF_ENABLED", false)
	//c.SetDefault("MW_FIBER_CSRF_KEYLOOKUP", "header:X-CSRF-Token")
	//c.SetDefault("MW_FIBER_CSRF_COOKIE_NAME", "_csrf")
	//c.SetDefault("MW_FIBER_CSRF_COOKIE_SAMESITE", "Strict")
	//c.SetDefault("MW_FIBER_CSRF_COOKIE_EXPIRES", "24h")
	//c.SetDefault("MW_FIBER_CSRF_CONTEXTKEY", "csrf")
	//
	//// Set default Fiber ETag middleware configuration
	//c.SetDefault("MW_FIBER_ETAG_ENABLED", false)
	//c.SetDefault("MW_FIBER_ETAG_WEAK", false)
	//
	//// Set default Fiber Expvar middleware configuration
	//c.SetDefault("MW_FIBER_EXPVAR_ENABLED", false)
	//
	//// Set default Fiber Favicon middleware configuration
	//c.SetDefault("MW_FIBER_FAVICON_ENABLED", false)
	//c.SetDefault("MW_FIBER_FAVICON_FILE", "")
	//c.SetDefault("MW_FIBER_FAVICON_CACHECONTROL", "public, max-age=31536000")
	//
	//// Set default Fiber Limiter middleware configuration
	//c.SetDefault("MW_FIBER_LIMITER_ENABLED", false)
	//c.SetDefault("MW_FIBER_LIMITER_MAX", 5)
	//c.SetDefault("MW_FIBER_LIMITER_EXPIRATION", "1m")
	//
	//// Set default Fiber Monitor middleware configuration
	//c.SetDefault("MW_FIBER_MONITOR_ENABLED", false)
	//
	//// Set default Fiber Pprof middleware configuration
	//c.SetDefault("MW_FIBER_PPROF_ENABLED", false)
	//
	//// Set default Fiber Recover middleware configuration
	//c.SetDefault("MW_FIBER_RECOVER_ENABLED", true)
	//
	//// Set default Fiber RequestID middleware configuration
	//c.SetDefault("MW_FIBER_REQUESTID_ENABLED", false)
	//c.SetDefault("MW_FIBER_REQUESTID_HEADER", "X-Request-ID")
	//c.SetDefault("MW_FIBER_REQUESTID_CONTEXTKEY", "requestid")
}

func (c *Config) getFiberViewsEngine() fiber.Views {
	var viewsEngine fiber.Views
	views := c.FiberViews()

	switch strings.ToLower(views.Engine) {
	case "ace":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".ace"
		}
		engine := ace.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	case "amber":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".amber"
		}
		engine := amber.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	case "django":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".django"
		}
		engine := django.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout)
		viewsEngine = engine
		break
	case "handlebars":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".hbs"
		}
		engine := handlebars.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	case "jet":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".jet"
		}
		engine := jet.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	case "mustache":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".mustache"
		}
		engine := mustache.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	case "pug":
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".pug"
		}
		engine := pug.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	// Use the official html/template package by default
	default:
		// Set file extension dynamically to FIBER_VIEWS
		if views.Extension == "" {
			views.Extension = ".html"
		}
		engine := html.New(views.Directory, views.Extension)
		engine.Reload(views.Reload).
			Debug(views.Debug).
			Layout(views.Layout).
			Delims(views.DelimsL, views.DelimsR)
		viewsEngine = engine
		break
	}

	return viewsEngine
}

func (c *Config) FiberConfig() *fiber.Config {
	if c.options.FiberConfig == nil {
		c.options.FiberConfig = &FiberConfig{
			Prefork:                   false,
			ServerHeader:              "",
			StrictRouting:             false,
			CaseSensitive:             false,
			Immutable:                 false,
			UnescapePath:              false,
			ETag:                      false,
			BodyLimit:                 4194303,
			Concurrency:               262144,
			ReadTimeout:               0,
			WriteTimeout:              0,
			IdleTimeout:               0,
			ReadBufferSize:            4096,
			WriteBufferSize:           4096,
			CompressedFileSuffix:      ".fiber.gz",
			ProxyHeader:               "",
			GETOnly:                   false,
			DisableKeepalive:          false,
			DisableDefaultDate:        false,
			DisableDefaultContentType: false,
			DisableHeaderNormalizing:  false,
			DisableStartupMessage:     false,
			ReduceMemoryUsage:         false,
		}
	}
	cfg := c.options.FiberConfig
	c.fiber = &fiber.Config{
		Prefork:                   cfg.Prefork,
		ServerHeader:              cfg.ServerHeader,
		StrictRouting:             cfg.StrictRouting,
		CaseSensitive:             cfg.CaseSensitive,
		Immutable:                 cfg.Immutable,
		UnescapePath:              cfg.UnescapePath,
		ETag:                      cfg.ETag,
		BodyLimit:                 cfg.BodyLimit,
		Concurrency:               cfg.Concurrency,
		Views:                     c.getFiberViewsEngine(),
		ViewsLayout:               cfg.ViewsLayout,
		ReadTimeout:               cfg.ReadTimeout,
		WriteTimeout:              cfg.WriteTimeout,
		IdleTimeout:               cfg.IdleTimeout,
		ReadBufferSize:            cfg.ReadBufferSize,
		WriteBufferSize:           cfg.WriteBufferSize,
		CompressedFileSuffix:      cfg.CompressedFileSuffix,
		ProxyHeader:               cfg.ProxyHeader,
		GETOnly:                   cfg.GETOnly,
		ErrorHandler:              defaultErrorHandler,
		DisableKeepalive:          cfg.DisableKeepalive,
		DisableDefaultDate:        cfg.DisableDefaultDate,
		DisableDefaultContentType: cfg.DisableDefaultContentType,
		DisableHeaderNormalizing:  cfg.DisableHeaderNormalizing,
		DisableStartupMessage:     cfg.DisableStartupMessage,
		AppName:                   c.options.Name,
		ReduceMemoryUsage:         cfg.ReduceMemoryUsage,
		EnableTrustedProxyCheck:   cfg.EnableTrustedProxyCheck,
		TrustedProxies:            cfg.TrustedProxies,
		Network:                   cfg.Network,
	}
	return c.fiber
}

//func (config *Config) setLoggerConfig() {
//	config.logger = &LoggerConfig{
//		ClientType:        config.GetString("LOG_TYPE"),
//		Environment: config.GetString("LOG_ENV"),
//		Filename:    config.GetString("LOG_FILENAME"),
//		MaxSize:     config.GetInt("LOG_MAXSIZE"),
//		MaxAge:      config.GetInt("LOG_MAXAGE"),
//		MaxBackups:  config.GetInt("LOG_MAXBACKUPS"),
//		LocalTime:   config.GetBool("LOG_LOCALTIME"),
//		Compress:    config.GetBool("LOG_COMPRESS"),
//		Verbose:     config.GetBool("LOG_VERBOSE"),
//	}
//}

func (c *Config) LoggerConfig() *LoggerConfig {
	return c.logger
}

func (c *Config) HasherConfig() hashing.Config {
	driver := c.HasherDriver()
	if strings.ToLower(driver.Driver) == "bcrypt" {
		return hashing.Config{
			Driver: bcrypt_driver.New(bcrypt_driver.Config{
				Complexity: driver.Rounds,
			})}
	} else {
		return hashing.Config{
			Driver: argon_driver.New(argon_driver.Config{
				Params: &argon2id.Params{
					Memory:      driver.Memory,
					Iterations:  driver.Iterations,
					Parallelism: driver.Parallelism,
					SaltLength:  driver.SaltLength,
					KeyLength:   driver.KeyLength,
				}})}
	}
}

func (c *Config) SessionConfig() session.Config {
	var storage fiber.Storage
	cfg := c.Session()

	switch strings.ToLower(cfg.Provider) {
	case "memcache":
		storage = memcache.New(memcache.Config{
			Servers: cfg.Host + ":" + strconv.Itoa(cfg.Port),
		})

		break
	case "mysql":
		storage = mysql.New(mysql.Config{
			Host:       cfg.Host,
			Port:       cfg.Port,
			Username:   cfg.Username,
			Password:   cfg.Password,
			Database:   cfg.Database,
			Table:      cfg.Table,
			GCInterval: cfg.GCInterval,
		})
		break
	case "postgresql", "postgres":
		storage = postgres.New(postgres.Config{
			Host:       cfg.Host,
			Port:       cfg.Port,
			Username:   cfg.Username,
			Password:   cfg.Password,
			Database:   cfg.Database,
			Table:      cfg.Table,
			GCInterval: cfg.GCInterval,
		})
		break
	case "redis":
		db, _ := strconv.Atoi(cfg.Database)
		sessionProvider := redis.New(redis.Config{
			Host:     cfg.Host,
			Port:     cfg.Port,
			Username: cfg.Username,
			Password: cfg.Password,
			Database: db,
		})
		storage = sessionProvider
		break
	case "sqlite3":
		sessionProvider := sqlite3.New(sqlite3.Config{
			Database:   cfg.Database,
			Table:      cfg.Table,
			GCInterval: cfg.GCInterval,
		})
		storage = sessionProvider
		break
	}

	return session.Config{
		KeyLookup:      cfg.KeyLookup,
		CookieSecure:   cfg.CookieSecure,
		CookieDomain:   cfg.CookieDomain,
		CookieSameSite: cfg.CookieSameSite,
		CookiePath:     cfg.CookiePath,
		Expiration:     cfg.Expiration,
		Storage:        storage,
	}
}

func initLogger(debug bool) {
	once.Do(func() {
		if debug {

		} else {

		}
	})
}

// NewConfig initialises a new configuration file
func NewConfig(ctx *cli.Context) *Config {
	initLogger(ctx.Bool("debug"))

	c := &Config{
		options:      NewOptions(ctx),
		token:        rnd.Token(8),
		errorHandler: defaultErrorHandler,
	}

	if configFile := c.ConfigFile(); c.options.ConfigFile == "" && fs.FileExists(configFile) {
		if err := c.options.Load(configFile); err != nil {
			log.Warnf("config: %s", err)
		} else {
			log.Debugf("config: options loaded from %s", txt.Quote(configFile))
		}
	}

	return c
}

// Options returns the raw config options.
func (c *Config) Options() *Options {
	if c.options == nil {
		log.Warnf("config: options should not be nil - bug?")
		c.options = NewOptions(nil)
	}

	return c.options
}

// Propagate updates config options in other packages as needed.
func (c *Config) Propagate() {
	//log.SetLevel(c.LogLevel())
	//
	//places.UserAgent = c.UserAgent()
	//entity.GeoApi = c.GeoApi()

	c.Settings().Propagate()
}

// Init creates directories, parses additional config files, opens a database
// connection and initializes dependencies.
func (c *Config) Init() error {
	if err := c.CreateDirectories(); err != nil {
		return err
	}

	if err := c.initStorage(); err != nil {
		return err
	}

	if insensitive, err := c.CaseInsensitive(); err != nil {
		return err
	} else if insensitive {
		log.Infof("config: case-insensitive file system detected")
		fs.IgnoreCase()
	}

	if cpuName := cpuid.CPU.BrandName; cpuName != "" {
		log.Debugf("config: running on %s", txt.Quote(cpuid.CPU.BrandName))
	}

	c.initSettings()

	c.Propagate()

	return c.connectDB()
}

// initStorage initializes storage directories with a random serial.
func (c *Config) initStorage() error {
	if c.serial != "" {
		return nil
	}
	const serialName = "serial"
	c.serial = rnd.PPID('z')

	storageName := filepath.Join(c.StoragePath(), serialName)
	backupName := filepath.Join(c.BackupPath(), serialName)

	if data, err := os.ReadFile(storageName); err == nil {
		c.serial = string(data)
	} else if data, err := os.ReadFile(backupName); err == nil {
		c.serial = string(data)
	} else if err := os.WriteFile(storageName, []byte(c.serial), os.ModePerm); err != nil {
		return fmt.Errorf("failed creating %s: %s", storageName, err)
	} else if err := os.WriteFile(backupName, []byte(c.serial), os.ModePerm); err != nil {
		return fmt.Errorf("failed creating %s: %s", backupName, err)
	}

	return nil
}

// Serial returns the random storage serial.
func (c *Config) Serial() string {
	if err := c.initStorage(); err != nil {
		log.Errorf("config: %s", err)
	}

	return c.serial
}

// SerialChecksum returns the CRC32 checksum of the storage serial.
func (c *Config) SerialChecksum() string {
	var result []byte

	hash := crc32.New(crc32.MakeTable(crc32.Castagnoli))

	if _, err := hash.Write([]byte(c.Serial())); err != nil {
		log.Warnf("config: %s", err)
	}

	return hex.EncodeToString(hash.Sum(result))
}

// Name returns the application name ("Gosh ESP").
func (c *Config) Name() string {
	return c.options.Name
}

// Version returns the application version.
func (c *Config) Version() string {
	return c.options.Version
}

// UserAgent returns a HTTP user agent string based on app name & version.
func (c *Config) UserAgent() string {
	return fmt.Sprintf("%s/%s", c.Name(), c.Version())
}

// Copyright returns the application copyright.
func (c *Config) Copyright() string {
	return c.options.Copyright
}

// BaseUri returns the site base URI for a given resource.
func (c *Config) BaseUri(res string) string {
	if c.SiteUrl() == "" {
		return res
	}

	u, err := url.Parse(c.SiteUrl())

	if err != nil {
		return res
	}

	return strings.TrimRight(u.Path, "/") + res
}

// ApiUri returns the api URI.
func (c *Config) ApiUri() string {
	return c.BaseUri(ApiUri)
}

// CdnUrl returns the optional content delivery network URI without trailing slash.
func (c *Config) CdnUrl(res string) string {
	return strings.TrimRight(c.options.CdnUrl, "/") + res
}

// ContentUri returns the content delivery URI.
func (c *Config) ContentUri() string {
	return c.CdnUrl(c.ApiUri())
}

// StaticUri returns the static content URI.
func (c *Config) StaticUri() string {
	return c.CdnUrl(c.BaseUri(StaticUri))
}

// SiteUrl returns the public server URL (default is "http://localhost:2345/").
func (c *Config) SiteUrl() string {
	if c.options.SiteUrl == "" {
		return "http://localhost:2345/"
	}

	return strings.TrimRight(c.options.SiteUrl, "/") + "/"
}

// SitePreview returns the site preview image URL for sharing.
func (c *Config) SitePreview() string {
	if c.options.SitePreview == "" {
		return c.SiteUrl() + "static/img/preview.jpg"
	}

	if !strings.HasPrefix(c.options.SitePreview, "http") {
		return c.SiteUrl() + c.options.SitePreview
	}

	return c.options.SitePreview
}

// SiteTitle returns the main site title (default is application name).
func (c *Config) SiteTitle() string {
	if c.options.SiteTitle == "" {
		return c.Name()
	}

	return c.options.SiteTitle
}

// SiteCaption returns a short site caption.
func (c *Config) SiteCaption() string {
	return c.options.SiteCaption
}

// SiteDescription returns a long site description.
func (c *Config) SiteDescription() string {
	return c.options.SiteDescription
}

// SiteAuthor returns the site author / copyright.
func (c *Config) SiteAuthor() string {
	return c.options.SiteAuthor
}

// Debug tests if debug mode is enabled.
func (c *Config) Debug() bool {
	return c.options.Debug
}

// Test tests if test mode is enabled.
func (c *Config) Test() bool {
	return c.options.Test
}

// Demo tests if demo mode is enabled.
func (c *Config) Demo() bool {
	return c.options.Demo
}

// Sponsor reports if your continuous support helps to pay for development and operating expenses.
func (c *Config) Sponsor() bool {
	return c.options.Sponsor || c.Test()
}

// Public tests if app runs in public mode and requires no authentication.
func (c *Config) Public() bool {
	if c.Demo() {
		return true
	}

	return c.options.Public
}

// Experimental tests if experimental features should be enabled.
func (c *Config) Experimental() bool {
	return c.options.Experimental
}

// ReadOnly tests if photo directories are write-protected.
func (c *Config) ReadOnly() bool {
	return c.options.ReadOnly
}

// AdminPassword returns the initial admin password.
func (c *Config) AdminPassword() string {
	return c.options.AdminPassword
}

// LogLevel returns the zap log level.
func (c *Config) LogLevel() zapcore.Level {
	if c.Debug() {
		c.options.LogLevel = "debug"
	}
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(c.options.LogLevel)); err != nil {
		return zapcore.InfoLevel
	}
	return logLevel
}

// Shutdown services and workers.
func (c *Config) Shutdown() {
	mutex.MainWorker.Cancel()
	mutex.ShareWorker.Cancel()
	mutex.SyncWorker.Cancel()
	mutex.MetaWorker.Cancel()

	if err := c.CloseDB(); err != nil {
		log.Errorf("could not close database connection: %s", err)
	} else {
		log.Info("closed database connection")
	}
}

// WakeupInterval returns the background worker wakeup interval duration.
func (c *Config) WakeupInterval() time.Duration {
	if c.options.WakeupInterval <= 0 || c.options.WakeupInterval > 86400 {
		return 15 * time.Minute
	}

	return time.Duration(c.options.WakeupInterval) * time.Second
}

// Workers returns the number of workers e.g. for indexing files.
func (c *Config) Workers() int {
	// NumCPU returns the number of logical CPU cores.
	cores := runtime.NumCPU()

	// Limit to physical cores to avoid high load on HT capable CPUs.
	if cores > cpuid.CPU.PhysicalCores {
		cores = cpuid.CPU.PhysicalCores
	}

	// Limit number of workers when using SQLite to avoid database locking issues.
	if c.DatabaseDriver() == SQLite && (cores >= 8 && c.options.Workers <= 0 || c.options.Workers > 4) {
		return 4
	}

	// Return explicit value if set and not too large.
	if c.options.Workers > runtime.NumCPU() {
		return runtime.NumCPU()
	} else if c.options.Workers > 0 {
		return c.options.Workers
	}

	// Use half the available cores by default.
	if cores > 1 {
		return cores / 2
	}

	return 1
}

// AutoIndex returns the auto indexing delay duration.
func (c *Config) AutoIndex() time.Duration {
	if c.options.AutoIndex < 0 {
		return time.Duration(0)
	} else if c.options.AutoIndex == 0 || c.options.AutoIndex > 86400 {
		return 5 * time.Minute
	}

	return time.Duration(c.options.AutoIndex) * time.Second
}

// AutoImport returns the auto importing delay duration.
func (c *Config) AutoImport() time.Duration {
	if c.options.AutoImport < 0 || c.ReadOnly() {
		return time.Duration(0)
	} else if c.options.AutoImport == 0 || c.options.AutoImport > 86400 {
		return 3 * time.Minute
	}

	return time.Duration(c.options.AutoImport) * time.Second
}

// GeoApi returns the preferred geocoding api (none or places).
func (c *Config) GeoApi() string {
	if c.options.DisablePlaces {
		return ""
	}

	return "places"
}

func (c *Config) AccessLogger() *AccessLoggerConfig {
	if c.options.AccessLoggerConfig == nil {
		c.options.AccessLoggerConfig = &AccessLoggerConfig{
			Enabled:     true,
			Type:        "console",
			Environment: "",
			Filename:    "access.log",
			MaxSize:     500,
			MaxAge:      28,
			MaxBackups:  3,
			LocalTime:   false,
			Compress:    false,
		}
	}
	return c.options.AccessLoggerConfig
}

func (c *Config) Session() *SessionConfig {
	if c.options.SessionConfig == nil {
		c.options.SessionConfig = &SessionConfig{
			Provider:       "sqlite3",
			KeyLookup:      "cookie:session_id",
			Database:       "esp_db",
			Table:          "sessions",
			Host:           "localhost",
			Port:           5432,
			Username:       "esp",
			Password:       "Pentinum#1",
			CookieDomain:   "",
			CookieSameSite: "Lax",
			Expiration:     12 * time.Hour,
			GCInterval:     1 * time.Minute,
			KeyPrefix:      "sessions",
			CookieSecure:   false,
		}
	}
	return c.options.SessionConfig
}

func (c *Config) ForceHTTPS() bool {
	return c.options.ForceHTTPS
}

func (c *Config) ForceTrailingSlash() bool {
	return c.options.ForceTrailingSlash
}

func (c *Config) HSTSEnabled() bool {
	return c.options.EnableHSTS
}

func (c *Config) HSTSMaxAge() int {
	return c.options.HSTSMaxAge
}

func (c *Config) HSTSIncludeSubdomains() bool {
	return c.options.HSTSIncludeSubdomains
}

func (c *Config) HSTSPreload() bool {
	return c.options.HSTSPreload
}

func (c *Config) SuppressWWW() bool {
	return c.options.SuppressWWW
}

func (c *Config) FiberRecoverDisabled() bool {
	return c.options.FiberRecoverDisabled
}

func (c *Config) FiberCacheEnabled() bool {
	return c.options.FiberCacheEnabled
}

func (c *Config) FiberCacheExpiration() time.Duration {
	return c.options.FiberCacheExpiration
}

func (c *Config) FiberCacheCacheControl() bool {
	return c.options.FiberCacheCacheControl
}

func (c *Config) FiberCompressEnabled() bool {
	return c.options.FiberCompressEnabled
}

func (c *Config) FiberCompressLevel() int {
	return c.options.FiberCompressLevel
}

func (c *Config) FiberCORSEnabled() bool {
	return c.options.FiberCORSEnabled
}

func (c *Config) FiberCORSAllowOrigins() string {
	return c.options.FiberCORSAllowOrigins
}

func (c *Config) FiberCORSAllowMethods() string {
	return c.options.FiberCORSAllowMethods
}

func (c *Config) FiberCORSAllowHeaders() string {
	return c.options.FiberCORSAllowHeaders
}

func (c *Config) FiberCORSAllowCredentials() bool {
	return c.options.FiberCORSAllowCredentials
}

func (c *Config) FiberCORSExposeHeaders() string {
	return c.options.FiberCORSExposeHeaders
}

func (c *Config) FiberCORSMaxAge() int {
	return c.options.FiberCORSMaxAge
}

func (c *Config) FiberCSRFEnabled() bool {
	return c.options.FiberCSRFEnabled
}

func (c *Config) FiberCSRFKeyLookup() string {
	return c.options.FiberCSRFKeyLookup
}

func (c *Config) FiberCSRFCookieName() string {
	return c.options.FiberCSRFCookieName
}

func (c *Config) FiberCSRFCookieSameSite() string {
	return c.options.FiberCSRFCookieSameSite
}

func (c *Config) FiberCSRFExpiration() time.Duration {
	return c.options.FiberCSRFExpiration
}

func (c *Config) FiberCSRFContextKey() string {
	return c.options.FiberCSRFContextKey
}

func (c *Config) FiberETagEnabled() bool {
	return c.options.FiberETagEnabled
}

func (c *Config) FiberETagWeak() bool {
	return c.options.FiberETagWeak
}

func (c *Config) FiberExpvarEnabled() bool {
	return c.options.FiberExpvarEnabled
}

func (c *Config) FiberFaviconEnabled() bool {
	return c.options.FiberFaviconEnabled
}

func (c *Config) FiberFaviconFile() string {
	return c.options.FiberFaviconFile
}

func (c *Config) FiberFaviconCacheControl() string {
	return c.options.FiberFaviconCacheControl
}

func (c *Config) FiberLimiterEnabled() bool {
	return c.options.FiberLimiterEnabled
}

func (c *Config) FiberLimiterMax() int {
	return c.options.FiberLimiterMax
}

func (c *Config) FiberLimiterExpiration() time.Duration {
	return c.options.FiberLimiterExpiration
}

func (c *Config) FiberMonitorEnabled() bool {
	return c.options.FiberMonitorEnabled
}

func (c *Config) FiberPprofEnabled() bool {
	return c.options.FiberPprofEnabled
}

func (c *Config) FiberRequestIDEnabled() bool {
	return c.options.FiberRequestIDEnabled
}

func (c *Config) FiberRequestIDHeader() string {
	return c.options.FiberRequestIDHeader
}

func (c *Config) FiberRequestIDContextKey() string {
	return c.options.FiberRequestIDContextKey
}

func (c *Config) HasherDriver() *HasherConfig {
	if c.options.HasherDriver == nil {
		c.options.HasherDriver = &HasherConfig{
			Driver:      "argon2id",
			Rounds:      bcrypt.DefaultRounds,
			Memory:      131072,
			SaltLength:  16,
			KeyLength:   32,
			Iterations:  4,
			Parallelism: 4,
		}
	}
	return c.options.HasherDriver
}

var defaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Set error message
	message := err.Error()
	//
	//// Check if it's a fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// TODO: Check return type for the client, JSON, HTML, YAML or any other (API vs web)

	return c.Status(code).JSON(fiber.Map{"message": message})

	// Render default error view
	//err = c.Render("errors/"+strconv.Itoa(code), fiber.Map{"message": message})
	//if err != nil {
	//	return c.SendString(message)
	//}
	//return err
}
