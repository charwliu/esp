package event

import (
	"github.com/leandro-lugaresi/hub"

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
	S().Error(msg)
	Publish("notify.error", Data{"message": msg})
}

func Success(msg string) {
	S().Info(msg)
	Publish("notify.success", Data{"message": msg})
}

func Info(msg string) {
	S().Info(msg)
	Publish("notify.info", Data{"message": msg})
}


func Warning(msg string) {
	S().Warn(msg)
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
