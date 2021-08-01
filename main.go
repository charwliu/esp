package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	configuration "go.vixal.xyz/esp/config"
	_ "go.vixal.xyz/esp/docs" // load API Docs files (Swagger)
	"go.vixal.xyz/esp/pkg/routes"
	"go.vixal.xyz/esp/server"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description ESP API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email charwliu@ngboss.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configuration.New()

	logger, undo := configureLogger(config)

	defer func() { undo() }()

	// Server initialization
	app, err := server.Create(config)

	if err != nil {
		logger.Error("%s", zap.Error(err))
		return
	}

	// Register web routes
	web := app.Group("")
	routes.RegisterWeb(web, app.Session, config.GetString("SESSION_LOOKUP"), app.DB, app.Hasher)

	// Register application API routes (using the /api/v1 group)
	api := app.Group("/api")
	apiv1 := api.Group("/v1")
	routes.RegisterAPI(apiv1, app.DB)

	// Register static routes for the public directory
	app.Static("/", "./public")

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		if err := c.SendStatus(fiber.StatusNotFound); err != nil {
			panic(err)
		}
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return err
	})

	// Close any connections on interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		<-c
		app.Exit()
	}()

	// Start listening on the specified address
	err = app.Listen(config.GetString("APP_ADDR"))
	if err != nil {
		logger.Error("Starting listening on the specified address", zap.Error(err))
		app.Exit()
	}
}

func configureLogger(config *configuration.Config) (*zap.Logger, func()) {
	ls := config.GetLoggerConfig()
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.InfoLevel
	})
	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	var consoleDebugging, consoleErrors zapcore.WriteSyncer
	//if ls.Verbose == true {
	consoleDebugging = zapcore.Lock(os.Stdout)
	consoleErrors = zapcore.Lock(os.Stderr)
	//} else {
	//	consoleDebugging = zapcore.AddSync(io.Discard)
	//	consoleErrors = zapcore.AddSync(io.Discard)
	//}
	var aLevel zap.AtomicLevel
	if ls.Verbose {
		aLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		aLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "ts",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	var core zapcore.Core
	var writerSyncer zapcore.WriteSyncer
	if ls.Filename != "" {
		logFilePath := filepath.Join(config.GetString("APP_WORKDIR"), ls.Filename)
		if filepath.IsAbs(ls.Filename) {
			logFilePath = ls.Filename
		}
		_, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			fmt.Printf("cannot create a log file: %s", err)
		}
		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			Compress:   ls.Compress, // disabled by default
			LocalTime:  ls.LocalTime,
			MaxBackups: ls.MaxBackups,
			MaxSize:    ls.MaxSize, // megabytes
			MaxAge:     ls.MaxAge,  // days
		}
		writerSyncer = zapcore.AddSync(lumberJackLogger)
		core = zapcore.NewCore(encoder, writerSyncer, aLevel)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, consoleErrors, highPriority),
			zapcore.NewCore(encoder, consoleDebugging, lowPriority),
		)
	}
	logger := zap.New(core, zap.AddCaller())
	undoGlobals := zap.ReplaceGlobals(logger)
	undoStd := zap.RedirectStdLog(logger)

	return logger, func() {
		undoStd()
		undoGlobals()
	}
}
