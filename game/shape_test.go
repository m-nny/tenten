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

func TestCanFit(t *testing.T) {
	state := newState()

	var x, y uint8 = 0, 0
	shape := Shapes[7] // 3x1
	if got := state.Fit(shape, x, y); got != nil {
		t.Errorf("state.Fit({%d, %d}, %d, %d) threw error %v, but it should not", shape.height, shape.width, x, y, got)
	}

	state.board[y][x] = true
	if got := state.Fit(shape, x, y); got == nil {
		t.Errorf("state.Fit({%d, %d}, %d, %d) did not throw error %v, but it should", shape.height, shape.width, x, y, got)
	}

	x, y = 9, 9
	if got := state.Fit(shape, x, y); got == nil {
		t.Errorf("state.Fit({%d, %d}, %d, %d) did not throw error %v, but it should", shape.height, shape.width, x, y, got)
	}
}
