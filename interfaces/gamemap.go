package interfaces

// GameMap is an interface that provides access to tile state and more.
type GameMap interface {
	MakeMap(maxRooms, roomMinSize, roomMaxSize int, entities *[]interfaces.Entity, maxMonsters int)
	Width() int
	Height() int
	IsBlocked(x, y int) bool
	IsOpaque(x, y int) bool
	Explored(x, y int) bool
	SetExplored(x, y int, explored bool)
	LastSeen(x, y int) rune
	SetLastSeen(x, y int, seen rune)
}