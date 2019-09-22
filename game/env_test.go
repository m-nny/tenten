package game

import (
	"fmt"
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
	state, shapes := env.Init()
	want = false
	for i, shape := range env.availableShapes {
		if got := shape.IsEmpty(); got != want {
			t.Errorf("NewEnvironment().Init().shape[%d].IsEmpty() = %t, want %t: %v", i, got, want, shape)
		}
	}

	wantBoard := env.state
	if got := state; got != wantBoard {
		t.Errorf("NewEnvironment().Init().board = %v, want %v", got, wantBoard)
	}
	wantShapes := env.availableShapes
	if got := shapes; got != wantShapes {
		t.Errorf("NewEnvironment().Init().shapes = %v, want %v", got, wantShapes)
	}
}

func TestMakeMove(t *testing.T) {
	rand.Seed(0)
	env := NewEnvironment()
	emptyState, _ := env.Init()

	var x, y, shapeIdx uint8 = 0, 0, 0

	state, shape, err := env.MakeMove(shapeIdx, x, y)
	fmt.Printf("Returned %#v -> %v %v\n\n", state.toString(), shape, err)
	if state.toString() == emptyState.toString() {
		t.Errorf("0. env.MakeMove(%d, %d, %d).state is %v, when it should not be", shapeIdx, x, y, state.toString())
	}
	if !shape[shapeIdx].IsEmpty() {
		t.Errorf("1. env.MakeMove(%d, %d, %d).shapes[%d] is %v not empty, when it should be", shapeIdx, x, y, shapeIdx, shape)
	}
	if err != nil {
		t.Errorf("2. env.MakeMove(%d, %d, %d) threw error %v, when it should not", shapeIdx, x, y, err)
	}

	// if state, shape, err = env.MakeMove(shapeIdx, x, y); err == nil {
	// 	t.Errorf("3. env.MakeMove(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, err)
	// }

	// x, y = 5, 5
	// if state, shape, err = env.MakeMove(shapeIdx, x, y); err == nil {
	// 	t.Errorf("4. env.MakeMove(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, err)
	// }

	// shapeIdx = ShapesNum
	// if _, _, err := env.MakeMove(shapeIdx, x, y); err == nil {
	// 	t.Errorf("5. env.Fit(%d, %d, %d) did not throw error %v, when it should", shapeIdx, x, y, err)
	// }
}
