package mapping

import (
	"star/entity"
	"star/interfaces"

	"github.com/kettek/goro"
)

// GameMap is our map data type.
type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
}

// Initialize initializes a GameMap's Tiles to match its Width and Height. It also sets up some coordinates to block movement and sight.
func (g *GameMap) Initialize() {
	g.Tiles = make([][]Tile, g.Width)

	for x := range g.Tiles {
		g.Tiles[x] = make([]Tile, g.Height)
		for y := range g.Tiles[x] {
			g.Tiles[x][y] = Tile{
				BlockSight:    true,
				BlockMovement: true,
			}
		}
	}
}

// MakeMap creates a new randomized map. This is built according to the passed arguments.
func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize int, entities *[]interfaces.Entity, maxMonsters int) {
	var rooms []*Rect

	for r := 0; r < maxRooms; r++ {
		// Generate a random width and height.
		width := roomMinSize + goro.Random.Intn(roomMaxSize)
		height := roomMinSize + goro.Random.Intn(roomMaxSize)
		// Generate a random position within the map boundaries.
		x := goro.Random.Intn(g.Width - width - 1)
		y := goro.Random.Intn(g.Height - height - 1)
		// Create a Rect according to our generated sizes.
		room := NewRect(x, y, width, height)

		// Iterate through our existing rooms to check for intersection with our new room.
		intersects := false
		for _, otherRoom := range rooms {
			if room.Intersect(otherRoom) {
				intersects = true
				break
			}
		}
		// Add the room if there is no intersection found.
		if !intersects {
			g.CreateRoom(room)
			roomCenterX, roomCenterY := room.Center()

			// Always place the player in the center of the first room.
			if len(rooms) == 0 {
				(*entities)[0].SetX(roomCenterX)
				(*entities)[0].SetY(roomCenterY)
			} else {
				prevCenterX, prevCenterY := rooms[len(rooms)-1].Center()

				// Flip a coin if we should tunnel horizontally or vertically first.
				if goro.Random.Intn(1) == 1 {
					g.CreateHTunnel(prevCenterX, roomCenterX, prevCenterY)
					g.CreateVTunnel(prevCenterY, roomCenterY, roomCenterX)
				} else {
					g.CreateVTunnel(prevCenterY, roomCenterY, prevCenterX)
					g.CreateHTunnel(prevCenterX, roomCenterX, roomCenterY)
				}
			}
			// Place random monsters in the room.
			g.PlaceEntities(room, entities, maxMonsters)

			// Append our new room to our rooms list.
			rooms = append(rooms, room)
		}
	}
}

// CreateRoom creates a room from a provided Rect.
func (g *GameMap) CreateRoom(r *Rect) {
	for x := r.X1 + 1; x < r.X2; x++ {
		for y := r.Y1 + 1; y < r.Y2; y++ {
			g.SetTile(x, y, Tile{})
		}
	}
}

// CreateHTunnel creates a horizontal tunnel from x1 to/from x2 starting at y.
func (g *GameMap) CreateHTunnel(x1, x2, y int) {
	for x := goro.MinInt(x1, x2); x <= goro.MaxInt(x1, x2); x++ {
		g.SetTile(x, y, Tile{})
	}
}

// CreateVTunnel creates a vertical tunnel from y1 to/from y2 starting at x.
func (g *GameMap) CreateVTunnel(y1, y2, x int) {
	for y := goro.MinInt(y1, y2); y <= goro.MaxInt(y1, y2); y++ {
		g.SetTile(x, y, Tile{})
	}
}

// PlaceEntities places 0 to maxMonsters monster entites in the provided room.
func (g *GameMap) PlaceEntities(room *Rect, entities *[]interfaces.Entity, maxMonsters int) {
	monstersCount := goro.Random.Intn(maxMonsters)

	for i := 0; i < monstersCount; i++ {
		var monster interfaces.Entity
		//Acquire a random location within the room.
		x := (1 + room.X1) + goro.Random.Intn(room.X2-room.X1-1)
		y := (1 + room.Y1) + goro.Random.Intn(room.Y2-room.Y1-1)

		if entity.FindEntityAtLocation(*entities, x, y, 0, 0) == nil {
			//Generate an Algol with 80% probability or a Saiph with 20%.
			if goro.Random.Intn(100) < 80 {
				monster = entity.NewEntity(x, y, '1', goro.Style{Foreground: goro.Color{0x2E, 0x2E, 0x2E, 0xFF}}, "Algol", entity.BlockMovement)
			} else {
				monster = entity.NewEntity(x, y, '8', goro.Style{Foreground: goro.Color{0x2E, 0x2E, 0x2E, 0xFF}}, "Saiph", entity.BlockMovement)
			}
			*entities = append(*entities, monster)
		}
	}
}

// Explored returns if the tile at x by y has been explored.
func (g *GameMap) Explored(x, y int) bool {
	if g.InBounds(x, y) {
		return g.Tiles[x][y].Explored
	}
	return false
}

// SetExplored sets the explored state of the tile at x by y to the passed explored bool.
func (g *GameMap) SetExplored(x, y int, explored bool) {
	if g.InBounds(x, y) {
		g.Tiles[x][y].Explored = explored
	}
}

// LastSeen returns the last seen rune.
func (g *GameMap) LastSeen(x, y int) rune {
	if g.InBounds(x, y) {
		return g.Tiles[x][y].LastSeen
	}
	return rune(0)
}

// SetLastSeen sets the last seen rune to the entity provided.
func (g *GameMap) SetLastSeen(x, y int, seen rune) {
	if g.InBounds(x, y) {
		g.Tiles[x][y].LastSeen = seen
	}
}

// IsBlocked returns if the given coordinates are blocking movement.
func (g *GameMap) IsBlocked(x, y int) bool {
	// Always block if outside our GameMap's bounds.
	if !g.InBounds(x, y) {
		return true
	}
	return g.Tiles[x][y].BlockMovement
}

// SetTile sets the tile at these specified coordinates to the value of t.
func (g *GameMap) SetTile(x, y int, t Tile) {
	if g.InBounds(x, y) {
		g.Tiles[x][y] = t
	}
}

// InBounds ensures that X and Y are within the map's bounds.
func (g *GameMap) InBounds(x, y int) bool {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return false
	}
	return true
}
