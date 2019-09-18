package game

import "testing"

func TestCanFit(t *testing.T) {
	state := NewState()

	var x, y uint8 = 0, 0
	shape := Shapes[7] // 3x1
	want := true
	if got := state.CanFit(shape, x, y); got != want {
		t.Errorf("state.CanFit({%d, %d}, %d, %d) = %t, want %t", shape.height, shape.width, x, y, got, want)
	}

	state.Board[y][x] = true
	want = false

	x, y, want = 9, 9, false
	if got := state.CanFit(shape, x, y); got != want {
		t.Errorf("state.CanFit({%d, %d}, %d, %d) = %t, want %t", shape.height, shape.width, x, y, got, want)
	}
}
