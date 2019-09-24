package game

import "testing"

func TestIsEmpty(t *testing.T) {
	shape := Shape{}
	want := true
	if got := shape.IsEmpty(); got != want {
		t.Errorf("Shape{}.IsEmpty() = %v, want %v", got, want)
	}
	shape = Shapes[0]
	want = false
	if got := shape.IsEmpty(); got != want {
		t.Errorf("Shapes[0].IsEmpty() = %v, want %v", got, want)
	}
}

func TestApply(t *testing.T) {
	state := newState()

	action := Action{7, 0, 0} // 3x1
	if got := state.Fit(Shapes[action.Idx], action.X, action.Y); got != nil {
		t.Errorf("state.Fit(%d) threw error %v, but it should not", action, got)
	}

	state.Board[action.Y][action.X] = true
	if got := state.Fit(Shapes[action.Idx], action.X, action.Y); got == nil {
		t.Errorf("state.Fit(%d) did not throw error %v, but it should", action, got)
	}

	action.X, action.Y = 9, 9
	if got := state.Fit(Shapes[action.Idx], action.X, action.Y); got == nil {
		t.Errorf("state.Fit(%d) did not throw error %v, but it should", action, got)
	}
}
