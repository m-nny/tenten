package agents

import "github.com/m-nny/tenten/game"

// Agent interface
type Agent interface {
	think(state game.State) (action game.Action)
}
