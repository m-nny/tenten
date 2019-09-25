package agents

import (
	"fmt"

	"github.com/m-nny/tenten/game"
)

// Play ...
func Play(agent Agent) (score uint32) {
	env := game.NewEnvironment()
	state := env.Init()
	for {
		if state.Over() {
			break
		}
		// fmt.Printf("%v\n%v\n", state.ToString(), state.AvailableShapes)
		action := agent.Think(state)
		// fmt.Printf("%v\n\n", action)
		newState, reward, err := env.MakeMove(action)
		if err == nil {
			score += uint32(reward)
			state = newState
		} else {
			fmt.Println(err)
			break
		}
	}
	fmt.Printf("%v\n%v\n", state.ToString(), state.AvailableShapes)
	return score
}

// MultiPlay ...
func MultiPlay(agent Agent, n int) (score uint32) {
	for i := 0; i < n; i++ {
		score += Play(agent)
	}
	return score
}
