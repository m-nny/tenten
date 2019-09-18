package game

import "strings"

// BoardCell underlying type of each board cell
type BoardCell bool

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
	Board [Boardsize][Boardsize]BoardCell
}

// NewState constuctor
func NewState() *State {
	state := &State{}
	return state
}

// ToString of string
func (s *State) ToString() (result string) {
	for _, row := range s.Board {
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
