package game

import (
	"testing"
)

func TestNewState(t *testing.T) {
	state := NewState()
	want := Boardsize
	if got := len(state.Board); got != want {
		t.Errorf("len(NewState()) = %d, want %d", got, want)
	}
	if got := len(state.Board[0]); got != want {
		t.Errorf("len(NewState()[0]) = %d, want %d", got, want)
	}
}

func TestToString(t *testing.T) {
	state := NewState()
	want := "..........\n..........\n..........\n..........\n..........\n..........\n..........\n..........\n..........\n.........."
	if got := state.ToString(); got != want {
		t.Errorf("ToString() = %q, want %q", got, want)
	}
	want = "#.........\n.#........\n..#.......\n...#......\n....#.....\n.....#....\n......#...\n.......#..\n........#.\n.........#"
	for i := 0; i < Boardsize; i++ {
		state.Board[i][i] = true
	}
	if got := state.ToString(); got != want {
		t.Errorf("ToString() = %q, want %q", got, want)
	}
}
