package mapping

// Tile represents the state of a given location in GameMap.
type Tile struct {
	BlockMovement bool
	BlockSight    bool
	Explored      bool
	LastSeen      rune
}
