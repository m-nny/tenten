package main

import (
	"fmt"

	"github.com/m-nny/tenten/agents"
)

func singlePlay(agent agents.Agent) {
	score := agents.Play(agent)
	fmt.Printf("Score: %v\n", score)
}

func multiPlay(agent agents.Agent) {
	runs := 3
	score := agents.MultiPlay(agent, runs)
	meanScore := float32(score) / float32(runs)
	fmt.Printf("Score: %v\n", meanScore)
}

func main() {
	agent := &agents.LazyAgent{}
	singlePlay(agent)
	multiPlay(agent)
}
