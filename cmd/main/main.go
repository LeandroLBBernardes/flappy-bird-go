package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/game"
)

func main() {
	eg, _ := game.NewGame()

	if err := ebiten.RunGame(eg); err != nil {
		log.Fatal(err)
	}
}
