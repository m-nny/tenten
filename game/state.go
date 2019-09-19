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
	board Board
}

// newState constuctor
func newState() *State {
	state := &State{}
	return state
}

// toString of string
func (s *State) toString() (result string) {
	for _, row := range s.board {
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
