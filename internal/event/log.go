package event

import (
	"github.com/leandro-lugaresi/hub"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

type Hook struct {
	hub *hub.Hub
}

func NewHook(hub *hub.Hub) *Hook {
	return &Hook{hub: hub}
}

func (h *Hook) Fire(entry zapcore.Entry) error {
	h.hub.Publish(Message{
		Name: "log." + entry.Level.String(),
		Fields: Data{
			"time":    entry.Time,
			"level":   entry.Level.String(),
			"message": entry.Message,
		},
	})
	return nil
}

func init() {
	hook := NewHook(SharedHub())
	Log, _ = zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	zapcore.RegisterHooks(Log.Core(), hook.Fire)
}
