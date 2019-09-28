package game

import (
	"testing"
)

func TestNewState(t *testing.T) {
	state := newState()
	want := Boardsize
	if got := len(state.Board); got != want {
		t.Errorf("len(NewState()) = %d, want %d", got, want)
	}
	if got := len(state.Board[0]); got != want {
		t.Errorf("len(NewState()[0]) = %d, want %d", got, want)
	}
}

func TestToString(t *testing.T) {
	state := newState()
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

func TestCanFit(t *testing.T) {
	state := newState()
	want, shape, x, y := true, Shapes[0], uint8(0), uint8(0)
	if got := state.CanFit(shape, x, y); got != want {
		t.Errorf("state.CanFit(%v, %v, %v) = %v, want %v", shape, x, y, got, want)
	}
	want, state.Board[x][y] = false, true
	if got := state.CanFit(shape, x, y); got != want {
		t.Errorf("state.CanFit(%v, %v, %v) = %v, want %v", shape, x, y, got, want)
	}
	want, x, y = false, Boardsize, Boardsize
	if got := state.CanFit(shape, x, y); got != want {
		t.Errorf("state.CanFit(%v, %v, %v) = %v, want %v", shape, x, y, got, want)
	}
}

func TestOver(t *testing.T) {
	state, want := newState(), true
	if got := state.Over(); got != want {
		t.Errorf("state.Over() = %v, want %v", got, want)
	}
	state.refillShapes()
	want = false
	if got := state.Over(); got != want {
		t.Errorf("state.Over() = %v, want %v", got, want)
	}
}
