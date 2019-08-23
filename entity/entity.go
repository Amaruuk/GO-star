package entity

import (
	"math"
	"star/interfaces"

	"github.com/kettek/goro"
)

// Entity is a type that represents an active entity in the world.
type Entity struct {
	x, y    int
	rune    rune
	style   goro.Style
	name    string
	flags   Flags
	fighter interfaces.Fighter
	actor   interfaces.Actor
}

//NewEntity returns a pointer to a newly created Entity.
func NewEntity(x int, y int, r rune, s goro.Style, name string, flags Flags) interfaces.Entity {
	return &Entity{
		x:     x,
		y:     y,
		rune:  r,
		style: s,
		name:  name,
		flags: flags,
	}
}

// Move moves the entity a given amount.
func (e *Entity) Move(x, y int) {
	e.x += x
	e.y += y
}

// DistanceTo gets the distance between two entities.
func (e *Entity) DistanceTo(other interfaces.Entity) float64 {
	destX := float64(e.x - other.X())
	destY := float64(e.y - other.Y())

	return math.Sqrt(math.Pow(destX, 2) + math.Pow(destY, 2))
}

// X returns the entity's x value.
func (e *Entity) X() int {
	return e.x
}

// SetX sets the entity's x value.
func (e *Entity) SetX(x int) {
	e.x = x
}

// Y returns the entity's y value.
func (e *Entity) Y() int {
	return e.y
}

// SetY sets the entity's x value.
func (e *Entity) SetY(y int) {
	e.y = y
}

// Rune returns the entity's rune.
func (e *Entity) Rune() rune {
	return e.rune
}

// SetRune sets the entity's rune.
func (e *Entity) SetRune(r rune) {
	e.rune = r
}

// Style returns the entity's style.
func (e *Entity) Style() goro.Style {
	return e.style
}

// SetStyle sets the entity's style.
func (e *Entity) SetStyle(s goro.Style) {
	e.style = s
}

// Name returns the entity's name.
func (e *Entity) Name() string {
	return e.name
}

// SetName sets the entity's name.
func (e *Entity) SetName(n string) {
	e.name = n
}

// Flags returns the entity's flags.
func (e *Entity) Flags() uint {
	return e.flags
}

// SetFlags sets the entity's flags.
func (e *Entity) SetFlags(f uint) {
	e.flags = f
}

// Actor gets the Actor interface owned by the entity.
func (e *Entity) Actor() interfaces.Actor {
	return e.actor
}

// SetActor sets the entity's Actor to the passed one.
func (e *Entity) SetActor(a interfaces.Actor) {
	e.actor = a
	e.actor.SetOwner(e)
}

//Fighter gets the Fighter interface owned by the entity.
func (e *Entity) Fighter() interfaces.Fighter {
	return e.fighter
}

// SetFighter sets the entity's Fighter to the passed one.
func (e *Entity) SetFighter(f interfaces.Fighter) {
	e.fighter = f
	e.fighter.SetOwner(e)
}

// FindEntityAtLocation finds and returns the first entity at x and y matching the provided flags. If none exists, it returns nil.
func FindEntityAtLocation(entities []interfaces.Entity, x, y int, checkMask Flags, matchFlags Flags) interfaces.Entity {
	for _, e := range entities {
		if (e.Flags()&checkMask) == matchFlags && e.X() == x && e.Y() == y {
			return e
		}
	}
	return nil
}
