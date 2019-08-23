package entity

import (
	"fmt"
	"math"

	"star/interfaces"

	"github.com/kettek/goro/fov"
	"github.com/kettek/goro/pathing"
)

// MonsterActor is the Actor that controls an monster Entity.
type MonsterActor struct {
	owner  interfaces.Entity
	target interfaces.Entity
}

// Owner gets the actor's owner.
func (a *MonsterActor) Owner() interfaces.Entity {
	return a.owner
}

// SetOwner sets the actor's owner to the given entity.
func (a *MonsterActor) SetOwner(e interfaces.Entity) {
	a.owner = e
}

// Target returns the actor's current target.
func (a *MonsterActor) Target() interfaces.Entity {
	return a.target
}

// SetTarget sets the actor's current target to the provided one.
func (a *MonsterActor) SetTarget(e interfaces.Entity) {
	a.target = e
}

// TakeTurn is the actor's turn-based logic step.
func (a *MonsterActor) TakeTurn(fovMap fov.Map, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	if fovMap.Visible(a.owner.X(), a.owner.Y()) {
		if a.target != nil && a.owner.DistanceTo(a.target) >= 3 {
			a.PathfindTo(a.target, gameMap, entities)
		} else if a.target != nil {
			fmt.Printf("The %s stares coldly at you.\n", a.owner.Name())
		}
	}
}

// PathfindTo pathfinds the actor's entity to the given target.
func (a *MonsterActor) PathfindTo(target interfaces.Entity, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	path := pathing.NewPathFromFunc(gameMap.Width(), gameMap.Height(), func(x, y int) int {
		if gameMap.IsBlocked(x, y) {
			return pathing.MaximumCost
		}
		other := FindEntityAtLocation(entities, x, y, BlockMovement, BlockMovement)
		if other != nil && other != a.owner && other != target {
			return pathing.MaximumCost
		}
		return 0
	}, pathing.AlgorithmAStar)

	steps := path.Compute(a.owner.X(), a.owner.Y(), target.X(), target.Y())

	if len(steps) > 0 {
		x := steps[0].X()
		y := steps[0].Y()
		a.owner.SetX(x)
		a.owner.SetY(y)
	} else {
		a.MoveTo(target.X(), target.Y(), gameMap, entities)
	}
}

// MoveTo move's the actor's entity towards a given location.
func (a *MonsterActor) MoveTo(x, y int, gameMap interfaces.GameMap, entities []interfaces.Entity) {
	destX := float64(x - a.owner.X())
	destY := float64(x - a.owner.Y())

	distance := math.Sqrt(math.Pow(destX, 2) + math.Pow(destY, 2))

	destX = math.Round(destX / distance)
	destY = math.Round(destY / distance)

	targetX := a.owner.X() + int(destX)
	targetY := a.owner.Y() + int(destY)

	if !gameMap.IsBlocked(targetX, targetY) && FindEntityAtLocation(entities, targetX, targetY, BlockMovement, BlockMovement) == nil {
		a.owner.Move(int(destX), int(destY))
	}
}
