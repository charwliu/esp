package config

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	glogger "gorm.io/gorm/logger"

	"go.vixal.xyz/esp/internal/event"
)

func NewLoggerConfig(verbose bool) *LoggerConfig {
	return &LoggerConfig{
		Type:        "console",
		Environment: "",
		Filename:    "",
		MaxSize:     500,
		MaxAge:      28,
		MaxBackups:  3,
		LocalTime:   false,
		Compress:    false,
		Verbose:     verbose,
	}
}
func (c *Config) LoggerConfig() *LoggerConfig {
	if c.loggerConfig == nil {
		c.loggerConfig = NewLoggerConfig(false)
	}
	return c.loggerConfig
}

func getLogWriter(root string, ls *LoggerConfig) zapcore.WriteSyncer {
	if ls.Filename == "" {
		return nil
	}
	logFilePath := filepath.Join(root, ls.Filename)
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
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder(c *LoggerConfig) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig

	switch c.Environment {
	case "Production":
		encoderConfig = zap.NewProductionEncoderConfig()
	case "Development":
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func NewLogger(root string, ls *LoggerConfig) (*zap.Logger, func()) {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.InfoLevel
	})
	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	var aLevel zap.AtomicLevel
	if ls.Verbose {
		aLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		aLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	}

	encoder := getEncoder(ls)
	var core zapcore.Core

	writerSyncer := getLogWriter(root, ls)
	if writerSyncer != nil {
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

	//event.Log = logger.WithOptions(zap.AddCallerSkip(1))
	return logger, func() {
		undoStd()
		undoGlobals()
	}
}

type LoggerConfig struct {
	// Type determines whether zap will be initialized as a file loggerConfig or,
	// by default, as a console loggerConfig.
	Type string `yaml:"LogType" json:"type"`

	// Environment determines whether zap will be initialized using a production
	// or a development loggerConfig.
	Environment string `yaml:"Environment" json:"environment"`

	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.
	Filename string `yaml:"Filename" json:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated.
	MaxSize int `yaml:"MaxSize" json:"max_size"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `yaml:"MaxAge" json:"max_age"`

	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int `yaml:"MaxBackups" json:"max_backups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.
	LocalTime bool `yaml:"LocalTime" json:"local_time"`

	// Compress determines if the rotated log files should be compressed
	// using gzip.
	Compress bool `yaml:"Compress" json:"compress"`

	Verbose bool `yaml:"Verbose" json:"verbose"`
}

type dblogger struct {
	glogger.Config
	logger                              *zap.SugaredLogger
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (l *dblogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newLogger := *l
	l.logger = l.logger.Desugar().WithOptions(zap.IncreaseLevel(zap.NewAtomicLevelAt(zapcore.Level(glogger.
		Info - level)))).Sugar()
	return &newLogger
}

func (l dblogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Infof(msg, data)
}

func (l dblogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Warnf(msg, data)
}

func (l dblogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Error(msg, data)
}

func (l dblogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && (!errors.Is(err, glogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.logger.Debugf(l.traceStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Debugf(l.traceStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= glogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL => %v", l.SlowThreshold)
		if rows == -1 {
			l.logger.Debugf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Debugf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == glogger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Debugf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.logger.Debugf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

// Colors
const (
	Reset       = "\033[0m"
	Yellow      = "\033[33m"
	Magenta     = "\033[35m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
)

func NewGormLogger(level glogger.LogLevel) glogger.Interface {
	var (
		traceStr     = "\n" + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr  = MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	)
	return &dblogger{
		Config: glogger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      level,
			Colorful:      true,
		},
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		logger:       L().WithOptions(zap.AddCallerSkip(1)).Sugar(),
	}
}
