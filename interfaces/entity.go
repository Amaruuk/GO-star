package interfaces

import (
	"github.com/kettek/goro"
)

// Entity is an interface that gets and sets parameters for entities.
type Entity interface {
	Move(int, int)
	DistanceTo(other Entity) float64
	// Basic property accessors
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
	// Actor and Fighter accessors
	Actor() Actor
	SetActor(Actor)
	Fighter() Fighter
	SetFighter(Fighter)
}
