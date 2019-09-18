package main

import (
	"fmt"

	"github.com/m-nny/tenten/game"
)

func main() {
	state := game.NewState()
	fmt.Printf("%#v\n", state.ToString())
	for i := 0; i < game.Boardsize; i++ {
		state.Board[i][i] = true
	}
	fmt.Printf("%#v\n", state.ToString())
}
