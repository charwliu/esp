package entity

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestMain(m *testing.M) {
	logger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.ErrorLevel))
	log = logger.Sugar()

	if err := os.Remove(".test.db"); err == nil {
		log.Debug("removed .test.db")
	}

	db := InitTestDb(Postgres,
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=UTC",
			"esp", "Pentinum#1", "esp_db", "localhost", 5432))
	defer db.Close()

	code := m.Run()

	if err := os.Remove(".test.db"); err == nil {
		log.Debug("removed .test.db")
	}
	os.Exit(code)
}
