package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/game"
)

//go:embed assets
var embeddedAssets embed.FS

func main() {
	eg, _ := game.NewGame(embeddedAssets)

	if err := ebiten.RunGame(eg); err != nil {
		log.Fatal(err)
	}
}
