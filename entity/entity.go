package entity

import (
	"github.com/kettek/goro"
)

// Entity is a type that represents an active entity in the world.
type Entity struct {
	X, Y  int
	Rune  rune
	Style goro.Style
	Name  string
	Flags Flags
}

// Move moves the entity a given amount.
func (e *Entity) Move(x, y int) {
	e.X += x
	e.Y += y
}

//NewEntity returns a pointer to a newly created Entity.
func NewEntity(x int, y int, r rune, s goro.Style, name string, flags Flags) *Entity {
	return &Entity{
		X:     x,
		Y:     y,
		Rune:  r,
		Style: s,
		Name:  name,
		Flags: flags,
	}
}

// FindEntityAtLocation finds and returns the first entity at x and y matching the provided flags. If none exists, it returns nil.
func FindEntityAtLocation(entities []*Entity, x, y int, checkMask Flags, matchFlags Flags) *Entity {
	for _, e := range entities {
		if (e.Flags&checkMask) == matchFlags && e.X == x && e.Y == y {
			return e
		}
	}
	return nil
}
