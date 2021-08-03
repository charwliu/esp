package i18n

import (
	"os"
	"testing"

	"github.com/leonelquinteros/gotext"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gotext.Configure(localeDir, string(locale), "default")

	code := m.Run()

	os.Exit(code)
}

func TestMsg(t *testing.T) {
	t.Run("already exists", func(t *testing.T) {
		msg := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msg)
	})

	t.Run("unexpected error", func(t *testing.T) {
		msg := Msg(ErrUnexpected, "A cat")
		assert.Equal(t, "Unexpected error, please try again", msg)
	})

	t.Run("already exists chinese", func(t *testing.T) {
		SetLocale("zh")
		msgTrans := Msg(ErrAlreadyExists, "小猫")
		assert.Equal(t, "小猫已存在", msgTrans)
		SetLocale("")
		msgDefault := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msgDefault)
	})

}

func TestError(t *testing.T) {
	t.Run("already exists", func(t *testing.T) {
		err := Error(ErrAlreadyExists, "A cat")
		assert.EqualError(t, err, "A cat already exists")
	})

	t.Run("unexpected error", func(t *testing.T) {
		err := Error(ErrUnexpected, "A cat")
		assert.EqualError(t, err, "Unexpected error, please try again")
	})

	t.Run("already exists chinese", func(t *testing.T) {
		SetLocale("zh")
		errGerman := Error(ErrAlreadyExists, "小猫")
		assert.EqualError(t, errGerman, "小猫已存在")
		SetLocale("")
		errDefault := Error(ErrAlreadyExists, "A cat")
		assert.EqualError(t, errDefault, "A cat already exists")
	})
}
