package entity

// Flags represents multiple states stored in a bitflag.
type Flags = uint

// Our entity flags.
const (
	BlockMovement Flags = 1 << iota
)
