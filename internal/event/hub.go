package event

import (
	"sync"

	"github.com/leandro-lugaresi/hub"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/i18n"
)

type Hub = hub.Hub
type Data = hub.Fields
type Message = hub.Message

var channelCap = 100
var sharedHub = NewHub()

func NewHub() *Hub {
	return hub.New()
}

func SharedHub() *Hub {
	return sharedHub
}

func Error(msg string) {
	once.Do(func() {
		log = Log.WithOptions(zap.AddCallerSkip(1))
	})
	log.Error(msg)
	Publish("notify.error", Data{"message": msg})
}

func Success(msg string) {
	once.Do(func() {
		log = Log.WithOptions(zap.AddCallerSkip(1))
	})
	log.Info(msg)
	Publish("notify.success", Data{"message": msg})
}

func Info(msg string) {
	once.Do(func() {
		log = Log.WithOptions(zap.AddCallerSkip(1))
	})
	log.Info(msg)
	Publish("notify.info", Data{"message": msg})
}

var once sync.Once
var log *zap.Logger

func Warning(msg string) {
	once.Do(func() {
		log = Log.WithOptions(zap.AddCallerSkip(1))
	})
	log.Warn(msg)
	Publish("notify.warning", Data{"message": msg})
}

func ErrorMsg(id i18n.Message, params ...interface{}) {
	Error(i18n.Msg(id, params...))
}

func SuccessMsg(id i18n.Message, params ...interface{}) {
	Success(i18n.Msg(id, params...))
}

func InfoMsg(id i18n.Message, params ...interface{}) {
	Info(i18n.Msg(id, params...))
}

func WarningMsg(id i18n.Message, params ...interface{}) {
	Warning(i18n.Msg(id, params...))
}

func Publish(event string, data Data) {
	SharedHub().Publish(Message{
		Name:   event,
		Fields: data,
	})
}

func Subscribe(topics ...string) hub.Subscription {
	return SharedHub().NonBlockingSubscribe(channelCap, topics...)
}

func Unsubscribe(s hub.Subscription) {
	SharedHub().Unsubscribe(s)
}
