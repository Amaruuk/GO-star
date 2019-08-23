package interfaces

import (
	"github.com/kettek/goro/fov"
)

type Actor interface {
	Owner() Entity
	SetOwner(Entity)
	Target() Entity
	SetTarget(Entity)
	TakeTurn(fov.Map, GameMap, []Entity)
}
