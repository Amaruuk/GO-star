package interfaces

import (
	"github.com/kettek/goro"
)

type Entity interface {
	Move(int, int)
	X() int
	SetX(int)
	Y() int
	SetY(int)
	Rune() rune
	SetRune(rune)
	Style() goro.Style
	SetStyle(goro.Style)
	Name() string
	SetName(string)
	Flags() uint
	SetFlags(uint)
}
