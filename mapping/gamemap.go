package mapping

import (
	"star/entity"

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
func (g *GameMap) MakeMap(maxRooms, roomMinSize, roomMaxSize int, player *entity.Entity) {
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
				player.X = roomCenterX
				player.Y = roomCenterY
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
