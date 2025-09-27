package main

import (
	"bufio"
	"number-guessing-go/game"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	g := game.NewGame(reader)
	g.Play()
}
