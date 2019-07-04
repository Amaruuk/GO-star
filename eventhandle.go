package main

import (
	"github.com/kettek/goro"
)

// handleKeyEvent converts a KeyEvent into a corresponding Action.
func handleKeyEvent(ev goro.EventKey) Action {
	switch ev.Key {
	case goro.KeyUp:
		return ActionMove{Y: -1}
	case goro.KeyDown:
		return ActionMove{Y: 1}
	case goro.KeyRight:
		return ActionMove{X: 1}
	case goro.KeyLeft:
		return ActionMove{X: -1}
	case goro.KeyEscape:
		return ActionQuit{}
	}
	return nil
}
