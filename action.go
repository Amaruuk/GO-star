package main

// Action is our interface for game actions.
type Action interface{}

// ActionMove is our action for moving something.
type ActionMove struct {
	X int
	Y int
}

// ActionQuit is our action to quit the program.
type ActionQuit struct{}
