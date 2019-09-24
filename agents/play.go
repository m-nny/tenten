package agents

import "github.com/m-nny/tenten/game"

// Play ...
func Play(agent Agent) (score uint32) {
	env := game.NewEnvironment()
	for state := env.Init(); !state.Over(); {
		action := agent.think(state)
		newState, reward, err := env.MakeMove(action)
		if err == nil {
			score += uint32(reward)
			state = newState
		}
	}
	return score
}
