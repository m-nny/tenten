package game

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewEnvironmentAndInit(t *testing.T) {
	env := NewEnvironment()
	want := true
	for i, shape := range env.state.AvailableShapes {
		if got := shape.IsEmpty(); got != want {
			t.Errorf("NewEnvironment().shape[%d].IsEmpty() = %t, want %t: %v", i, got, want, shape)
		}
	}
	state := env.Init()
	shapes := state.AvailableShapes
	want = false
	for i, shape := range env.state.AvailableShapes {
		if got := shape.IsEmpty(); got != want {
			t.Errorf("NewEnvironment().Init().shape[%d].IsEmpty() = %t, want %t: %v", i, got, want, shape)
		}
	}

	wantBoard := env.state
	if got := state; got != wantBoard {
		t.Errorf("NewEnvironment().Init().board = %v, want %v", got, wantBoard)
	}
	wantShapes := env.state.AvailableShapes
	if got := shapes; got != wantShapes {
		t.Errorf("NewEnvironment().Init().shapes = %v, want %v", got, wantShapes)
	}
}

func TestMakeMove(t *testing.T) {
	rand.Seed(0)
	env := NewEnvironment()
	emptyState := env.Init()

	action := Action{0, 0, 0}

	state, _, err := env.MakeMove(action)
	fmt.Printf("Returned %#v -> %v\n\n", state.ToString(), err)
	if state.ToString() == emptyState.ToString() {
		t.Errorf("0. env.MakeMove(%v).state is %v, when it should not be", action, state.ToString())
	}
	if !state.AvailableShapes[action.Idx].IsEmpty() {
		t.Errorf("1. env.MakeMove(%v).shapes[%d] is not empty, when it should be", action, action.Idx)
	}
	if err != nil {
		t.Errorf("2. env.MakeMove(%v) threw error %v, when it should not", action, err)
	}

	if state, _, err = env.MakeMove(action); err == nil {
		t.Errorf("3. env.MakeMove(%v) did not throw error %v, when it should", action, err)
	}

	action.X, action.Y = 5, 5
	if state, _, err = env.MakeMove(action); err == nil {
		t.Errorf("4. env.MakeMove(%v) did not throw error %v, when it should", action, err)
	}

	action.Idx = ShapesNum
	if _, _, err := env.MakeMove(action); err == nil {
		t.Errorf("5. env.Fit(%v) did not throw error %v, when it should", action, err)
	}
}
