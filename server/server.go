package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/expvar"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/session/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/app/models"
	configuration "go.vixal.xyz/esp/config"
	"go.vixal.xyz/esp/pkg/middleware"
	"go.vixal.xyz/esp/pkg/routes"
	"go.vixal.xyz/esp/platform/database"
)

type App struct {
	*fiber.App

	DB      *database.Database
	Hasher  hashing.Driver
	Session *session.Session
}

func Create(config *configuration.Config) (*App, error) {
	app := &App{
		App:     fiber.New(*config.GetFiberConfig()),
		Hasher:  hashing.New(config.GetHasherConfig()),
		Session: session.New(config.GetSessionConfig()),
	}

	app.registerMiddlewares(config)

	// Initialize database
	db, err := database.New(&database.Config{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	// Auto-migrate database models
	if err != nil {
		zap.S().Info("failed to connect to database: ", err.Error())
	} else {
		if db == nil {
			zap.S().Info("failed to connect to database: db variable is nil")
		} else {
			app.DB = db
			err := app.DB.AutoMigrate(&models.Role{})
			if err != nil {
				zap.S().Info("failed to auto-migrate role model:", err.Error())
				return app, err
			}
			err = app.DB.AutoMigrate(&models.User{})
			if err != nil {
				zap.S().Info("failed to auto-migrate user model:", err.Error())
				return app, err
			}
		}
	}
	// Register web routes
	web := app.Group("")
	routes.RegisterWeb(web, app.Session, config.GetString("SESSION_LOOKUP"), app.DB, app.Hasher)

	return app, nil
}

func (app *App) registerMiddlewares(config *configuration.Config) {
	// Middleware - Custom Access Logger based on zap
	if config.GetBool("MW_ACCESS_LOGGER_ENABLED") {
		app.Use(middleware.AccessLogger(&middleware.AccessLoggerConfig{
			Type:        config.GetString("MW_ACCESS_LOGGER_TYPE"),
			Environment: config.GetString("APP_ENV"),
			Filename:    config.GetString("MW_ACCESS_LOGGER_FILENAME"),
			MaxSize:     config.GetInt("MW_ACCESS_LOGGER_MAXSIZE"),
			MaxAge:      config.GetInt("MW_ACCESS_LOGGER_MAXAGE"),
			MaxBackups:  config.GetInt("MW_ACCESS_LOGGER_MAXBACKUPS"),
			LocalTime:   config.GetBool("MW_ACCESS_LOGGER_LOCALTIME"),
			Compress:    config.GetBool("MW_ACCESS_LOGGER_COMPRESS"),
		}))
	}

	// Middleware - Force HTTPS
	if config.GetBool("MW_FORCE_HTTPS_ENABLED") {
		app.Use(middleware.ForceHTTPS())
	}

	// Middleware - Force trailing slash
	if config.GetBool("MW_FORCE_TRAILING_SLASH_ENABLED") {
		app.Use(middleware.ForceTrailingSlash())
	}

	// Middleware - HSTS
	if config.GetBool("MW_HSTS_ENABLED") {
		app.Use(middleware.HSTS(&middleware.HSTSConfig{
			MaxAge:            config.GetInt("MW_HSTS_MAXAGE"),
			IncludeSubdomains: config.GetBool("MW_HSTS_INCLUDESUBDOMAINS"),
			Preload:           config.GetBool("MW_HSTS_PRELOAD"),
		}))
	}

	// Middleware - Suppress WWW
	if config.GetBool("MW_SUPPRESS_WWW_ENABLED") {
		app.Use(middleware.SuppressWWW())
	}

	// Middleware - Recover
	if config.GetBool("MW_FIBER_RECOVER_ENABLED") {
		app.Use(recover.New())
	}

	// Middleware - Recover
	if config.GetBool("MW_FIBER_RECOVER_ENABLED") {
		app.Use(recover.New())
	}

	// TODO: Middleware - Basic Authentication

	// Middleware - Cache
	if config.GetBool("MW_FIBER_CACHE_ENABLED") {
		app.Use(cache.New(cache.Config{
			Expiration:   config.GetDuration("MW_FIBER_CACHE_EXPIRATION"),
			CacheControl: config.GetBool("MW_FIBER_CACHE_CACHECONTROL"),
		}))
	}

	// Middleware - Compress
	if config.GetBool("MW_FIBER_COMPRESS_ENABLED") {
		lvl := compress.Level(config.GetInt("MW_FIBER_COMPRESS_LEVEL"))
		app.Use(compress.New(compress.Config{
			Level: lvl,
		}))
	}

	// Middleware - CORS
	if config.GetBool("MW_FIBER_CORS_ENABLED") {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     config.GetString("MW_FIBER_CORS_ALLOWORIGINS"),
			AllowMethods:     config.GetString("MW_FIBER_CORS_ALLOWMETHODS"),
			AllowHeaders:     config.GetString("MW_FIBER_CORS_ALLOWHEADERS"),
			AllowCredentials: config.GetBool("MW_FIBER_CORS_ALLOWCREDENTIALS"),
			ExposeHeaders:    config.GetString("MW_FIBER_CORS_EXPOSEHEADERS"),
			MaxAge:           config.GetInt("MW_FIBER_CORS_MAXAGE"),
		}))
	}

	// Middleware - CSRF
	if config.GetBool("MW_FIBER_CSRF_ENABLED") {
		app.Use(csrf.New(csrf.Config{
			KeyLookup:      config.GetString("MW_FIBER_CSRF_KEYLOOKUP"),
			CookieName:     config.GetString("MW_FIBER_CSRF_COOKIE_NAME"),
			CookieSameSite: config.GetString("MW_FIBER_CSRF_COOKIE_SAMESITE"),
			Expiration:     config.GetDuration("MW_FIBER_CSRF_COOKIE_EXPIRES"),
			ContextKey:     config.GetString("MW_FIBER_CSRF_CONTEXTKEY"),
		}))
	}

	// Middleware - ETag
	if config.GetBool("MW_FIBER_ETAG_ENABLED") {
		app.Use(etag.New(etag.Config{
			Weak: config.GetBool("MW_FIBER_ETAG_WEAK"),
		}))
	}

	// Middleware - Expvar
	if config.GetBool("MW_FIBER_EXPVAR_ENABLED") {
		app.Use(expvar.New())
	}

	// Middleware - Favicon
	if config.GetBool("MW_FIBER_FAVICON_ENABLED") {
		app.Use(favicon.New(favicon.Config{
			File:         config.GetString("MW_FIBER_FAVICON_FILE"),
			CacheControl: config.GetString("MW_FIBER_FAVICON_CACHECONTROL"),
		}))
	}

	// TODO: Middleware - Filesystem

	// Middleware - Limiter
	if config.GetBool("MW_FIBER_LIMITER_ENABLED") {
		app.Use(limiter.New(limiter.Config{
			Max:        config.GetInt("MW_FIBER_LIMITER_MAX"),
			Expiration: config.GetDuration("MW_FIBER_LIMITER_EXPIRATION"),
			// TODO: Key
			// TODO: LimitReached
		}))
	}

	// Middleware - Monitor
	if config.GetBool("MW_FIBER_MONITOR_ENABLED") {
		app.Use(monitor.New())
	}

	// Middleware - Pprof
	if config.GetBool("MW_FIBER_PPROF_ENABLED") {
		app.Use(pprof.New())
	}

	// TODO: Middleware - Proxy

	// Middleware - RequestID
	if config.GetBool("MW_FIBER_REQUESTID_ENABLED") {
		app.Use(requestid.New(requestid.Config{
			Header: config.GetString("MW_FIBER_REQUESTID_HEADER"),
			// TODO: Generator
			ContextKey: config.GetString("MW_FIBER_REQUESTID_CONTEXTKEY"),
		}))
	}

	// TODO: Middleware - Timeout
}

// Exit Stop the Fiber application
func (app *App) Exit() {
	_ = app.Shutdown()
}
