package main

import (
	"math/rand"

	"github.com/m-nny/tenten/game"
)

func main() {
	env := game.NewEnvironment()
	_, shapes := env.Init()
	chosenShape := uint8(rand.Intn(len(shapes)))
	env.Fit(chosenShape, 0, 0)
}
