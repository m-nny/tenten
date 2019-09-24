package agents

import "github.com/m-nny/tenten/game"

// LazyAgent ...
type LazyAgent struct {
}

// Think Agent's thinking mechanism
func (agent *LazyAgent) Think(state game.State) (action game.Action) {
	var chosen uint8
	for chosen = 0; chosen < game.ShapesNum; chosen++ {
		if !state.AvailableShapes[chosen].IsEmpty() {
			break
		}
	}
	chosenShape := state.AvailableShapes[chosen]
	var x, y uint8
	for x = 0; x < game.Boardsize; x++ {
		for y = 0; y < game.Boardsize; y++ {
			if state.CanFit(chosenShape, x, y) {
				return game.Action{chosen, x, y}
			}
		}
	}
	return game.Action{255, 255, 255}
}
