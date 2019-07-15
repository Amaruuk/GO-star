package main

import (
	"fmt"
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

	goro.Setup(func(screen *goro.Screen) {
		// Screen configuration.
		screen.SetTitle("StarSlayer")
		screen.SetSize(64, 36)
		screen.AutoSize = false
		screen.SetGlyphs(0, "Starslayer.ttf", 16)

		// Randomize our seed so the map is randomized per run.
		goro.SetSeed(goro.RandomSeed())
	})

	goro.Run(func(screen *goro.Screen) {
		// Our initial variables.
		mapWidth, mapHeight := 64, 36
		maxRooms, roomMinSize, roomMaxSize := 30, 6, 10
		maxMonstersPerRoom := 3
		gameState := PlayerTurnState

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

		player := entity.NewEntity(0, 0, '@', goro.Style{Foreground: goro.ColorBlack}, "Player", entity.BlockMovement)

		entities := []*entity.Entity{
			player,
		}

		gameMap.MakeMap(maxRooms, roomMinSize, roomMaxSize, &entities, maxMonstersPerRoom)

		fovMap := InitializeFoV(&gameMap)

		for {

			if fovRecompute {
				RecomputeFoV(fovMap, entities, gameMap, player.X, player.Y, fovRadius, fov.Light{})
			}

			// Draw screen.
			DrawAll(screen, entities, gameMap, fovMap, fovRecompute, colors)

			fovRecompute = false

			ClearAll(screen, entities, fovMap)

			// Handle events.
			switch event := screen.WaitEvent().(type) {
			case goro.EventKey:
				switch action := handleKeyEvent(event).(type) {
				case ActionMove:
					if gameState == PlayerTurnState {
						x := player.X + action.X
						y := player.Y + action.Y
						if !gameMap.IsBlocked(x, y) {
							otherEntity := entity.FindEntityAtLocation(entities, x, y, entity.BlockMovement, entity.BlockMovement)
							if otherEntity != nil {
								fmt.Printf("You stick your hand through the %s's body. It is quite clammy.\n", otherEntity.Name)
							} else {
								player.Move(action.X, action.Y)
								fovRecompute = true
						}
					}
					gameState = NPCTurnState
				}
				case ActionQuit:
					goro.Quit()
				}
			case goro.EventQuit:
				return
			}

			// Handle entity updates.
			if gameState == NPCTurnState {
				for i, e := range entities {
					if i > 0 {
						fmt.Printf("The %s spams the terminal.\n", e.Name)
					}
				}
				gameState = PlayerTurnState
			}
		}
	})
}
