package utils

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	FAVICON_PATH          = "../../assets/favicon.png"
	BACKGROUND_DAY_PATH   = "../../assets/sprites/background-day.png"
	BACKGROUND_NIGHT_PATH = "../../assets/sprites/background-night.png"
	MENU_PATH             = "../../assets/sprites/message.png"
	GAMEOVER_PATH         = "../../assets/sprites/gameover.png"
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
	a.loadAssets()
	return a
}

func (a *Assets) loadAssets() {
	loadImage(&a.Icon, FAVICON_PATH)
	loadImage(&a.BackgroundDay, BACKGROUND_DAY_PATH)
	loadImage(&a.BackgroundNight, BACKGROUND_NIGHT_PATH)
	loadImage(&a.Menu, MENU_PATH)
	loadImage(&a.GameOver, GAMEOVER_PATH)
}

func loadImage(target **ebiten.Image, path string) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("error to file image: %v", err)
	}

	*target = img
}
