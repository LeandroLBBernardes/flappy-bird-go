package utils

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	FAVICON_PATH          = "assets/favicon.png"
	BACKGROUND_DAY_PATH   = "assets/sprites/background-day.png"
	BACKGROUND_NIGHT_PATH = "assets/sprites/background-night.png"
	MENU_PATH             = "assets/sprites/message.png"
	GAMEOVER_PATH         = "assets/sprites/gameover.png"
)

type Assets struct {
	Icon            *ebiten.Image
	BackgroundDay   *ebiten.Image
	BackgroundNight *ebiten.Image
	Menu            *ebiten.Image
	GameOver        *ebiten.Image

	embeddedAssets embed.FS
}

func NewAssets(embeddedAssets embed.FS) *Assets {
	a := &Assets{}
	a.embeddedAssets = embeddedAssets
	a.loadAssets()
	return a
}

func (a *Assets) LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(a.embeddedAssets, path)
	if err != nil {
		log.Fatalf("error to load image: %v", err)
	}

	return img
}

func (a *Assets) loadAssets() {
	a.loadImageFromTarget(&a.Icon, FAVICON_PATH)
	a.loadImageFromTarget(&a.BackgroundDay, BACKGROUND_DAY_PATH)
	a.loadImageFromTarget(&a.BackgroundNight, BACKGROUND_NIGHT_PATH)
	a.loadImageFromTarget(&a.Menu, MENU_PATH)
	a.loadImageFromTarget(&a.GameOver, GAMEOVER_PATH)
}

func (a *Assets) loadImageFromTarget(target **ebiten.Image, path string) {
	img, _, err := ebitenutil.NewImageFromFileSystem(a.embeddedAssets, path)
	if err != nil {
		log.Fatalf("error to file image: %v", err)
	}

	*target = img
}
