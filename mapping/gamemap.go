package mapping

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
	}

	g.Tiles[30][22] = Tile{
		BlockMovement: true,
		BlockSight:    true,
	}
	g.Tiles[31][22] = Tile{
		BlockMovement: true,
		BlockSight:    true,
	}
	g.Tiles[32][22] = Tile{
		BlockMovement: true,
		BlockSight:    true,
	}
}

// IsBlocked returns if the given coordinates are blocking movement.
func (g *GameMap) IsBlocked(x, y int) bool {
	// Always block if outside our GameMap's bounds.
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return true
	}
	return g.Tiles[x][y].BlockMovement
}
