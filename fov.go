package main

import (
	"star/interfaces"

	"github.com/kettek/goro/fov"
)

// InitializeFoV initializes the FoV
func InitializeFoV(g interfaces.GameMap) fov.Map {
	fovMap := fov.NewMap(g.Width(), g.Height(), fov.AlgorithmBBQ)

	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			fovMap.SetBlocksMovement(x, y, g.IsBlocked(x, y))
			fovMap.SetBlocksLight(x, y, g.IsOpaque(x, y))
		}
	}
	return fovMap
}

// RecomputeFoV recomputes the FoV
func RecomputeFoV(fovMap fov.Map, entities []interfaces.Entity, gameMap interfaces.GameMap, centerX, centerY int, radius int, light fov.Light) {
	entitiesInView := make([]bool, len(entities))
	for i, e := range entities {
		if fovMap.Visible(e.X(), e.Y()) {
			entitiesInView[i] = true
		}
	}
	fovMap.Recompute(centerX, centerY, radius, light)
	for i, b := range entitiesInView {
		if b && !fovMap.Visible(entities[i].X(), entities[i].Y()) {
			gameMap.SetLastSeen(entities[i].X(), entities[i].Y(), entities[i].Rune())
		}
	}
}
