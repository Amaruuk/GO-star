package main

import (
	"log"

	"github.com/kettek/goro"

	"star/entity"
	"star/mapping"
)

func main() {
	// Initialize goro!
	if err := goro.InitTCell(); err != nil {
		log.Fatal(err)
	}

	goro.Run(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("StarSlayer")
		screen.SetSize(80, 24)
		screen.AutoSize = false

		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())

		// Our initial variables.
		mapWidth, mapHeight := 80, 24
		maxRooms, roomMinSize, roomMaxSize := 30, 6, 10

		colors := map[string]goro.Color{
			"darkWall":   goro.Color{0x80, 0x80, 0x80, 0xFF},
			"darkGround": goro.Color{0x80, 0x00, 0x00, 0xFF},
		}

		gameMap := mapping.GameMap{
			Width:  mapWidth,
			Height: mapHeight,
		}

		gameMap.Initialize()

		player := entity.NewEntity(screen.Columns/2, screen.Rows/2, '@', goro.Style{Foreground: goro.ColorWhite})
		npc := entity.NewEntity(screen.Columns/2-5, screen.Rows/2, '&', goro.Style{Foreground: goro.ColorOlive})

		entities := []*entity.Entity{
			player,
			npc,
		}

		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, player)

		for {
			// Draw screen.
			DrawAll(screen, entities, gameMap, colors)
			ClearAll(screen, entities)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if !gameMap.IsBlocked(player.X+action.X, player.Y+action.Y) {
						player.Move(action.X, action.Y)
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
