package i18n

import (
	"errors"
	"fmt"
	"strings"

	"github.com/leonelquinteros/gotext"
)

//go:generate xgettext --no-wrap --language=c --from-code=UTF-8 --output=../../assets/locales/messages.pot messages.go

type Message int
type MessageMap map[Message]string

func gettext(s string) string {
	return gotext.Get(s)
}

func Msg(id Message, params ...interface{}) string {
	msg := gotext.Get(Messages[id])

	if strings.Contains(msg, "%") {
		msg = fmt.Sprintf(msg, params...)
	}

	return msg
}

func Error(id Message, params ...interface{}) error {
	return errors.New(Msg(id, params...))
}
