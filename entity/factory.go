package entity

import (
	"star/interfaces"

	"github.com/kettek/goro"
)

// NewPlayerEntity creates and defines the player's entity.
func NewPlayerEntity() interfaces.Entity {
	pc := &Entity{
		rune: '@',
		name: "Player",
		style: goro.Style{
			Foreground: goro.ColorBlack,
		},
		flags: BlockMovement,
	}
	return pc
}

// NewMonsterEntity does a thing.
func NewMonsterEntity(x, y int, r rune, style goro.Style, name string) interfaces.Entity {
	c := &Entity{
		x: x,
		y: y,
		rune: r,
		name: name,
		style: style,
		flags: BlockMovement,
	}
	c.SetActor(&MonsterActor{
		owner: c,
	})

	return c
}