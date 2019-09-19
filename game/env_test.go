package game

import (
	"math/rand"
	"testing"
)

func TestNewEnvironmentAndInit(t *testing.T) {
	env := NewEnvironment()
	want := true
	for i, shape := range env.availableShapes {
		if got := shape.IsEmpty(); got != want {
			t.Errorf("NewEnvironment().shape[%d].IsEmpty() = %t, want %t: %v", i, got, want, shape)
		}
	}
	board, shapes := env.Init()
	want = false
	for i, shape := range env.availableShapes {
		if got := shape.IsEmpty(); got != want {
			t.Errorf("NewEnvironment().Init().shape[%d].IsEmpty() = %t, want %t: %v", i, got, want, shape)
		}
	}

	wantBoard := env.state.board
	if got := board; got != wantBoard {
		t.Errorf("NewEnvironment().Init().board = %v, want %v", got, wantBoard)
	}
	wantShapes := env.availableShapes
	if got := shapes; got != wantShapes {
		t.Errorf("NewEnvironment().Init().shapes = %v, want %v", got, wantShapes)
	}
}

func TestFit(t *testing.T) {
	rand.Seed(0)
	env := NewEnvironment()
	env.Init()

	var x, y, shapeIdx uint8 = 0, 0, 0

	if got := env.Fit(shapeIdx, x, y); got != nil {
		t.Errorf("env.Fit(%d, %d, %d) threw error %v, when it should not", shapeIdx, x, y, got)
	}

	if got := env.Fit(shapeIdx, x, y); got == nil {
		t.Errorf("env.Fit(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, got)
	}

	x, y = 5, 5
	if got := env.Fit(shapeIdx, x, y); got == nil {
		t.Errorf("env.Fit(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, got)
	}

	shapeIdx = ShapesNum
	if got := env.Fit(shapeIdx, x, y); got == nil {
		t.Errorf("env.Fit(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, got)
	}
}
