package main

import (
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
func RecomputeFoV(fovMap fov.Map, centerX, centerY int, radius int, light fov.Light) {
	fovMap.Recompute(centerX, centerY, radius, light)
}
