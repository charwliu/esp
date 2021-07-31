package api

import (
	"github.com/fatih/color"
)

var YellowFn = color.New(color.FgYellow).SprintFunc()
var RedFn = color.New(color.FgRed).SprintFunc()

func Color(fg color.Attribute, s string) interface{} {
	switch fg {
	case color.FgYellow:
		return YellowFn(s)
	case color.FgRed:
		return RedFn(s)
	}
	panic("impossible situation")
}
