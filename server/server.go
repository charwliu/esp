package server

import (
	"context"
	"fmt"
	"path/filepath"

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
	"github.com/gofiber/fiber/v2/middleware/session"
	hashing "github.com/thomasvvugt/fiber-hashing"

	"go.uber.org/zap"
	"gorm.io/gorm"

	configuration "go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/pkg/middleware"
	"go.vixal.xyz/esp/pkg/routes"
)

type App struct {
	*fiber.App

	DB      *gorm.DB
	Hasher  hashing.Driver
	Session *session.Store
}

func Create(config *configuration.Config) (*App, error) {
	app := &App{
		App:    fiber.New(*config.FiberConfig()),
		Hasher: hashing.New(config.HasherConfig()),
	}

	app.registerMiddlewares(config)

	app.DB = config.DB()

	// Register web routes
	web := app.Group("")
	routes.RegisterWeb(web, app.Hasher)

	return app, nil
}

func (app *App) registerMiddlewares(config *configuration.Config) {
	// Middleware - Custom Access Logger based on zap
	conf := config.AccessLogger()
	if conf.Enabled {
		app.Use(middleware.AccessLogger(&middleware.AccessLoggerConfig{
			Type:        conf.Type,
			Environment: conf.Environment,
			Filename:    conf.Filename,
			MaxSize:     conf.MaxSize,
			MaxAge:      conf.MaxAge,
			MaxBackups:  conf.MaxBackups,
			LocalTime:   conf.LocalTime,
			Compress:    conf.Compress,
		}))
	}

	// Middleware - Force HTTPS
	if config.ForceHTTPS() {
		app.Use(middleware.ForceHTTPS())
	}

	// Middleware - Force trailing slash
	if config.ForceTrailingSlash() {
		app.Use(middleware.ForceTrailingSlash())
	}

	// Middleware - HSTS
	if config.HSTSEnabled() {
		app.Use(middleware.HSTS(&middleware.HSTSConfig{
			MaxAge:            config.HSTSMaxAge(),
			IncludeSubdomains: config.HSTSIncludeSubdomains(),
			Preload:           config.HSTSPreload(),
		}))
	}

	// Middleware - Suppress WWW
	if config.SuppressWWW() {
		app.Use(middleware.SuppressWWW())
	}

	// Middleware - Recover
	if !config.FiberRecoverDisabled() {
		app.Use(recover.New())
	}

	// TODO: Middleware - Basic Authentication

	// Middleware - Cache
	if config.FiberCacheEnabled() {
		app.Use(cache.New(cache.Config{
			Expiration:   config.FiberCacheExpiration(),
			CacheControl: config.FiberCacheControl(),
		}))
	}

	// Middleware - Compress
	if config.FiberCompressEnabled() {
		lvl := compress.Level(config.FiberCompressLevel())
		app.Use(compress.New(compress.Config{
			Level: lvl,
		}))
	}

	// Middleware - CORS
	if config.FiberCORSEnabled() {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     config.FiberCORSAllowOrigins(),
			AllowMethods:     config.FiberCORSAllowMethods(),
			AllowHeaders:     config.FiberCORSAllowHeaders(),
			AllowCredentials: config.FiberCORSAllowCredentials(),
			ExposeHeaders:    config.FiberCORSExposeHeaders(),
			MaxAge:           config.FiberCORSMaxAge(),
		}))
	}

	// Middleware - CSRF
	if config.FiberCSRFEnabled() {
		app.Use(csrf.New(csrf.Config{
			KeyLookup:      config.FiberCSRFKeyLookup(),
			CookieName:     config.FiberCSRFCookieName(),
			CookieSameSite: config.FiberCSRFCookieSameSite(),
			Expiration:     config.FiberCSRFExpiration(),
			ContextKey:     config.FiberCSRFContextKey(),
		}))
	}

	// Middleware - ETag
	if config.FiberETagEnabled() {
		app.Use(etag.New(etag.Config{
			Weak: config.FiberETagWeak(),
		}))
	}

	// Middleware - Expvar
	if config.FiberExpvarEnabled() {
		app.Use(expvar.New())
	}

	// Middleware - Favicon
	if config.FiberFaviconEnabled() {
		app.Use(favicon.New(favicon.Config{
			File:         config.FiberFaviconFile(),
			CacheControl: config.FiberFaviconCacheControl(),
		}))
	}

	// TODO: Middleware - Filesystem

	// Middleware - Limiter
	if config.FiberLimiterEnabled() {
		app.Use(limiter.New(limiter.Config{
			Max:        config.FiberLimiterMax(),
			Expiration: config.FiberLimiterExpiration(),
			// TODO: Key
			// TODO: LimitReached
		}))
	}

	// Middleware - Monitor
	if config.FiberMonitorEnabled() {
		app.Use(monitor.New())
	}

	// Middleware - Pprof
	if config.FiberPprofEnabled() {
		app.Use(pprof.New())
	}

	// TODO: Middleware - Proxy

	// Middleware - RequestID
	if config.FiberRequestIDEnabled() {
		app.Use(requestid.New(requestid.Config{
			Header: config.FiberRequestIDHeader(),
			// TODO: Generator
			ContextKey: config.FiberRequestIDContextKey(),
		}))
	}

	// TODO: Middleware - Timeout
}

// Exit Stop the Fiber application
func (app *App) Exit() {
	_ = app.Shutdown()
}

func Start(ctx context.Context, config *configuration.Config) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		log.Error(err)
	//	}
	//}()

	app := &App{
		App:     fiber.New(*config.FiberConfig()),
		Hasher:  hashing.New(config.HasherConfig()),
		Session: session.New(config.SessionConfig()),
	}

	app.registerMiddlewares(config)

	// Register web routes
	web := app.Group("")
	routes.RegisterWeb(web, app.Hasher)

	// Register application API routes (using the /api/v1 group)
	api := app.Group("/api")
	apiv1 := api.Group("/v1")
	routes.RegisterAPI(apiv1)

	// Register static routes for the public directory
	app.Static("/", config.BuildPath())
	app.Static("/static", config.StaticPath())

	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile(filepath.Join(config.BuildPath(), "index.html"))
	})

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return nil
	})

	// Start listening on the specified address
	err := app.Listen(fmt.Sprintf("%s:%d", config.HttpHost(), config.HttpPort()))
	if err != nil {
		log.Error("Starting listening on the specified address", zap.Error(err))
		return
	}

}

var log = event.Log.Sugar()
