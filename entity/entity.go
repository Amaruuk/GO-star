package entity

import (
	"github.com/kettek/goro"
)

// Entity is a type that represents an active entity in the world.
type Entity struct {
	X, Y  int
	Rune  rune
	Style goro.Style
}

// Move moves the entity a given amount.
func (e *Entity) Move(x, y int) {
	e.X += x
	e.Y += y
}

//NewEntity returns a pointer to a newly created Entity.
func NewEntity(x int, y int, r rune, s goro.Style) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		Rune:  r,
		Style: s,
	}
}
