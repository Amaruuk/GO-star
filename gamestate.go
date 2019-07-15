package main

// GameState is a numerical representation of a game state.
type GameState = uint8

// Our various game states.
const (
	PlayerTurnState GameState = iota
	NPCTurnState
)
