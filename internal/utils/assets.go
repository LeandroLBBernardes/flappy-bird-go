package utils

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Assets struct {
	Icon            *ebiten.Image
	BackgroundDay   *ebiten.Image
	BackgroundNight *ebiten.Image
	Menu            *ebiten.Image
	GameOver        *ebiten.Image
}

func NewAssets() *Assets {
	a := &Assets{}
	a.init()
	return a
}

func (a *Assets) init() {
	var err error

	a.Icon, _, err = ebitenutil.NewImageFromFile("assets/favicon.png")
	if err != nil {
		log.Fatal(err)
	}

	a.BackgroundDay, _, err = ebitenutil.NewImageFromFile("assets/sprites/background-day.png")
	if err != nil {
		log.Fatal(err)
	}

	a.BackgroundNight, _, err = ebitenutil.NewImageFromFile("assets/sprites/background-night.png")
	if err != nil {
		log.Fatal(err)
	}

	a.Menu, _, err = ebitenutil.NewImageFromFile("assets/sprites/message.png")
	if err != nil {
		log.Fatal(err)
	}

	a.GameOver, _, err = ebitenutil.NewImageFromFile("assets/sprites/gameover.png")
	if err != nil {
		log.Fatal(err)
	}
}
