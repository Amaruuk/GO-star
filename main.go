package main

import (
	"log"

	"github.com/kettek/goro"
	"github.com/kettek/goro/fov"

	"star/entity"
	"star/mapping"
)

func main() {
	// Initialize goro!
	if err := goro.InitEbiten(); err != nil {
		log.Fatal(err)
	}

	goro.Run(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("StarSlayer")
		screen.SetSize(50, 50)
		screen.AutoSize = false
		screen.SetGlyphs(0, "Starslayer.ttf", 16)

		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())

		// Our initial variables.
		mapWidth, mapHeight := 50, 50
		maxRooms, roomMinSize, roomMaxSize := 30, 6, 10

		fovRadius := 10
		fovRecompute := true

		colors := map[string]goro.Color{
			"darkWall":    goro.Color{0x50, 0x50, 0x50, 0xFF},
			"darkGround":  goro.Color{0xC8, 0xC8, 0xC8, 0xFF},
			"lightWall":   goro.Color{0x80, 0x80, 0x80, 0xFF},
			"lightGround": goro.Color{0xDD, 0xDD, 0xDD, 0xFF},
		}

		gameMap := mapping.GameMap{
			Width:  mapWidth,
			Height: mapHeight,
		}

		gameMap.Initialize()

		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorBlack})
		npc := entity.NewEntity(screen.Columns/2-5, screen.Rows/2, '&', goro.Style{Foreground: goro.ColorBlack})

		entities := []*entity.Entity{
			player,
			npc,
		}

		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, player)

		fovMap := InitializeFoV(&gameMap)

		for {

			if fovRecompute {
				RecomputeFoV(fovMap, player.X, player.Y, fovRadius, fov.Light{})
			}

			// Draw screen.
			DrawAll(screen, entities, gameMap, fovMap, fovRecompute, colors)

			fovRecompute = false

			ClearAll(screen, entities)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if !gameMap.IsBlocked(player.X+action.X, player.Y+action.Y) {
						player.Move(action.X, action.Y)
						fovRecompute = true
					}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}
		}
	})
}
