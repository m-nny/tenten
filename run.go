package main

import (
	"fmt"

	"github.com/m-nny/tenten/agents"
)

func main() {
	runs := 1000
	lazy := &agents.LazyAgent{}
	score := agents.MultiPlay(lazy, runs)
	meanScore := float32(score) / float32(runs)
	fmt.Printf("Avg Score: %v\n", meanScore)
}
