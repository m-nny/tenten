package game

import "strings"

// BoardCell underlying type of each board cell
type BoardCell bool

// Board underlying type of board
type Board [Boardsize][Boardsize]BoardCell

const (
	// EmptyCell string representation
	EmptyCell string = "."
	// FilledCell string representationl
	FilledCell string = "#"
	// Boardsize ...
	Boardsize = 10
)

// State of game
type State struct {
	Board           Board
	AvailableShapes [ShapesNum]Shape
}

// newState constuctor
func newState() *State {
	state := &State{}
	return state
}

// CanFit try to fit shape in board in position (x, y)
func (state *State) CanFit(shape Shape, x, y uint8) bool {
	x0, y0 := x, y
	x1, y1 := x+shape.width, y+shape.height
	if x1 >= Boardsize || y1 >= Boardsize {
		return false
	}
	for i := y0; i < y1; i++ {
		for j := x0; j < x1; j++ {
			if state.Board[i][j] {
				return false
			}
		}
	}
	return true
}

// Fit ...
func (state *State) Fit(shape Shape, x, y uint8) error {
	x0, y0 := x, y
	x1, y1 := x+shape.width, y+shape.height
	if x1 > Boardsize || y1 > Boardsize {
		return ErrOutOfBounds(x1, Boardsize)
	}
	if !state.CanFit(shape, x, y) {
		return ErrAlreadyFilled(state, x0, y0)
	}
	for i := y0; i < y1; i++ {
		for j := x0; j < x1; j++ {
			state.Board[i][j] = true
		}
	}
	return nil
}

// Over ...
func (state *State) Over() bool {
	return false
}

// ToString of string
func (state *State) ToString() (result string) {
	for _, row := range state.Board {
		for _, cell := range row {
			if cell {
				result += FilledCell
			} else {
				result += EmptyCell
			}
		}
		result += "\n"
	}
	if len(result) > 0 {
		result = strings.TrimSuffix(result, "\n")
	}

	return
}
