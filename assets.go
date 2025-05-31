package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Assets struct {
	icon            *ebiten.Image
	backgroundDay   *ebiten.Image
	backgroundNight *ebiten.Image
	menu            *ebiten.Image
	gameOver        *ebiten.Image
}

func NewAssets() *Assets {
	a := &Assets{}
	a.init()
	return a
}

func (a *Assets) init() {
	var err error

	a.icon, _, err = ebitenutil.NewImageFromFile("assets/favicon.png")
	if err != nil {
		log.Fatal(err)
	}

	a.backgroundDay, _, err = ebitenutil.NewImageFromFile("assets/sprites/background-day.png")
	if err != nil {
		log.Fatal(err)
	}

	a.backgroundNight, _, err = ebitenutil.NewImageFromFile("assets/sprites/background-night.png")
	if err != nil {
		log.Fatal(err)
	}

	a.menu, _, err = ebitenutil.NewImageFromFile("assets/sprites/message.png")
	if err != nil {
		log.Fatal(err)
	}

	a.gameOver, _, err = ebitenutil.NewImageFromFile("assets/sprites/gameover.png")
	if err != nil {
		log.Fatal(err)
	}
}
