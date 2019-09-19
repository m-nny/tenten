package game

import (
	"testing"
)

func TestNewState(t *testing.T) {
	state := newState()
	want := Boardsize
	if got := len(state.board); got != want {
		t.Errorf("len(NewState()) = %d, want %d", got, want)
	}
	if got := len(state.board[0]); got != want {
		t.Errorf("len(NewState()[0]) = %d, want %d", got, want)
	}
}

func TestToString(t *testing.T) {
	state := newState()
	want := "..........\n..........\n..........\n..........\n..........\n..........\n..........\n..........\n..........\n.........."
	if got := state.toString(); got != want {
		t.Errorf("ToString() = %q, want %q", got, want)
	}
	want = "#.........\n.#........\n..#.......\n...#......\n....#.....\n.....#....\n......#...\n.......#..\n........#.\n.........#"
	for i := 0; i < Boardsize; i++ {
		state.board[i][i] = true
	}
	if got := state.toString(); got != want {
		t.Errorf("ToString() = %q, want %q", got, want)
	}
}
