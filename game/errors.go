package game

import (
	"fmt"
)

func ErrOutOfBounds(idx, bounds uint8) error {
	return fmt.Errorf("game: index %d is out of bounds %d", idx, bounds)
}

func ErrAlreadyFilled(state *State, x, y uint8) error {
	return fmt.Errorf("game: board at (%d,%d) is already filled: %v", x, y, state.toString())
}

func ErrAlredyUsed(shapeIdx uint8, availableShapes *[ShapesNum]Shape) error {
	return fmt.Errorf("game: shape (%d) was already used: %v", shapeIdx, availableShapes)
}

// func ErrOutOfBounds(idx, bounds uint8) error {
// 	return fmt.Errorf("game: index %d out of bounds %d", idx, bounds)
// }