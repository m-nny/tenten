package main

import (
	"fmt"

	"github.com/m-nny/tenten/agents"
)

func main() {
	lazy := agents.LazyAgent{}
	score := agents.Play(lazy)
	fmt.Printf("%v", score)
}
