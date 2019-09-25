package agents

import (
	"fmt"

	"github.com/m-nny/tenten/game"
)

// RealAgent ...
type RealAgent struct {
}

// Think ...
func (agent *RealAgent) Think(state game.State) game.Action {
	board := state.ToString()
	fmt.Printf("Board:\n%v\nAvailable shapes: %v\n", board, state.AvailableShapes)
	var idx, x, y uint8
	for {
		fmt.Printf("Enter action in form [index x y]: ")
		m, err := fmt.Scanf("%d %d %d", &idx, &x, &y)
		if m == 3 && err == nil {
			break
		}
	}
	return game.Action{Idx: idx, X: x, Y: y}
}
