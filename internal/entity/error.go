package entity

import (
	"time"

	"go.uber.org/zap/zapcore"

	"go.vixal.xyz/esp/internal/event"
)

// Error represents an error message log.
type Error struct {
	ID           uint      `gorm:"primary_key" json:"ID" yaml:"ID"`
	ErrorTime    time.Time `sql:"index" json:"Time" yaml:"Time"`
	ErrorLevel   string    `gorm:"type:varchar(32)" json:"Level" yaml:"Level"`
	ErrorMessage string    `gorm:"type:varchar(2048)" json:"Message" yaml:"Message"`
}

type Errors []Error

func (Error) TableName() string {
	return "error"
}

// SaveErrorMessages subscribes to error logs and stored them in the errors table.
func SaveErrorMessages() {
	s := event.Subscribe("log.*")

	defer func() {
		event.Unsubscribe(s)
	}()

	for msg := range s.Receiver {
		level, ok := msg.Fields["level"]

		if !ok {
			continue
		}

		var logLevel zapcore.Level
		err := logLevel.UnmarshalText([]byte(level.(string)))

		if err != nil || logLevel <= zapcore.InfoLevel {
			continue
		}

		newError := Error{ErrorLevel: logLevel.String()}

		if val, ok := msg.Fields["message"]; ok {
			newError.ErrorMessage = val.(string)
		}

		if val, ok := msg.Fields["time"]; ok {
			newError.ErrorTime = val.(time.Time)
		}

		DB().Create(&newError)
	}
}
