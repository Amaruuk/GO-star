package main

import (
	"star/entity"
	"star/mapping"

	"github.com/kettek/goro/fov"
)

// InitializeFoV initializes the FoV
func InitializeFoV(g *mapping.GameMap) fov.Map {
	fovMap := fov.NewMap(g.Width, g.Height, fov.AlgorithmBBQ)

	for x := range g.Tiles {
		for y, tile := range g.Tiles[x] {
			fovMap.SetBlocksMovement(x, y, tile.BlockMovement)
			fovMap.SetBlocksLight(x, y, tile.BlockSight)
		}
	}

	return fovMap
}

// RecomputeFoV recomputes the FoV
func RecomputeFoV(fovMap fov.Map, entities []*entity.Entity, gameMap mapping.GameMap, centerX, centerY int, radius int, light fov.Light) {
	entitiesInView := make([]bool, len(entities))
	for i, e := range entities {
		if fovMap.Visible(e.X, e.Y) {
			entitiesInView[i] = true
		}
	}
	fovMap.Recompute(centerX, centerY, radius, light)
	for i, b := range entitiesInView {
		if b && !fovMap.Visible(entities[i].X, entities[i].Y) {
			gameMap.SetLastSeen(entities[i].X, entities[i].Y, entities[i].Rune) 
		}
	}
}
