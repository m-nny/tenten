package game

import (
	"fmt"
)

// ErrOutOfBounds ...
func ErrOutOfBounds(idx, bounds uint8) error {
	return fmt.Errorf("game: index %d is out of bounds %d", idx, bounds)
}

// ErrAlreadyFilled ...
func ErrAlreadyFilled(state *State, x, y uint8) error {
	return fmt.Errorf("game: board at (%d,%d) is already filled: %v", x, y, state.ToString())
}

// ErrAlredyUsed ...
func ErrAlredyUsed(shapeIdx uint8, availableShapes *[ShapesNum]Shape) error {
	return fmt.Errorf("game: shape (%d) was already used: %v", shapeIdx, availableShapes)
}
