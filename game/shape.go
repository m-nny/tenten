package game

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

// CanFit test if shape can fit in board in specifit position
func (state *State) CanFit(shape Shape, x, y uint8) bool {
	if x+shape.width <= Boardsize && y+shape.height <= Boardsize {
		for i := y; i < y+shape.height; i++ {
			for j := x; j < x+shape.width; j++ {
				if state.Board[i][j] {
					return false
				}
			}
		}
		return true
	}
	return false
}
