package game

import (
	"math/rand"
)

// Shape ...
type Shape struct {
	height, width uint8
}

// Shapes all allowed in-game shapes
var Shapes = [...]Shape{
	Shape{1, 1}, // 0
	Shape{1, 2}, // 1
	Shape{1, 3}, // 2
	Shape{1, 4}, // 3
	Shape{1, 5}, // 4
	Shape{2, 1}, // 5
	Shape{2, 2}, // 6
	Shape{3, 1}, // 7
	Shape{3, 3}, // 8
	Shape{4, 1}, // 9
	Shape{5, 1}, // 10
}

// Fit try to fit shape in board in position (x, y)
func (state *State) Fit(shape Shape, x, y uint8) error {
	if x+shape.width > Boardsize || y+shape.height > Boardsize {
		return ErrOutOfBounds(x+shape.width, Boardsize)
	}
	for i := y; i < y+shape.height; i++ {
		for j := x; j < x+shape.width; j++ {
			if state.board[i][j] {
				return ErrAlreadyFilled(state, x, y)
			}
		}
	}
	for i := y; i < y+shape.height; i++ {
		for j := x; j < x+shape.width; j++ {
			state.board[i][j] = true
		}
	}
	return nil
}

func PickShape(population []Shape) Shape {
	idx := rand.Intn(len(population))
	return Shapes[idx]
}

func (shape *Shape) IsEmpty() bool {
	return shape.height == 0 || shape.width == 0
}
