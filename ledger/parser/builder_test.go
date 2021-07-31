package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	assert.Equal(t, "Здравейте", Capitalize("здравейте"))
	assert.Equal(t, "Ｔhis is 您好！", Capitalize("ｔhis is 您好！"))
	assert.Equal(t, "Ｔｈｉｓ is 您好！", Capitalize("ｔｈｉｓ is 您好！"))
}
