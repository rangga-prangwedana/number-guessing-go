package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	game := game.NewGame(reader)
	g.Play()
}
