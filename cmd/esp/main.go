package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.vixal.xyz/esp/internal/commands"
	"go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/internal/event"
)

var version = "development"
var log = event.Log

func main() {
	app := &cli.App{
		Name:                 "Esp",
		Usage:                "E-Service platform",
		Version:              version,
		Copyright:            "© 2021 Charlie W. Liu <charwliu@gmail.com>",
		EnableBashCompletion: true,
		Flags:                config.GlobalFlags,
	}

	app.Commands = []*cli.Command{
		commands.StartCommand,
		commands.StopCommand,
		commands.ConfigCommand,
		commands.VersionCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
	}
}

func configureLogger(config *config.Config) (*zap.Logger, func()) {
	ls := config.LoggerConfig()
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
		logFilePath := filepath.Join(config.StoragePath(), ls.Filename)
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
	zapcore.RegisterHooks(core, event.NewHook(event.SharedHub()).Fire)
	logger := zap.New(core, zap.AddCaller())
	undoGlobals := zap.ReplaceGlobals(logger)
	undoStd := zap.RedirectStdLog(logger)

	event.Log = logger.WithOptions(zap.AddCallerSkip(1))
	return logger, func() {
		undoStd()
		undoGlobals()
	}
}
