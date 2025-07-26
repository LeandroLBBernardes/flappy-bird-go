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

func LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("error to load image: %v", err)
	}

	return img
}

func DrawCentralizedImage(image *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	imageWidth := float64(image.Bounds().Dx())
	imageHeight := float64(image.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())
	posX := (screenWidth - imageWidth) / 2
	posY := (screenHeight - imageHeight) / 2

	op.GeoM.Translate(posX, posY)
	op.GeoM.Scale(1, 1)

	screen.DrawImage(image, op)
}

func (a *Assets) loadAssets() {
	loadImageFromTarget(&a.Icon, FAVICON_PATH)
	loadImageFromTarget(&a.BackgroundDay, BACKGROUND_DAY_PATH)
	loadImageFromTarget(&a.BackgroundNight, BACKGROUND_NIGHT_PATH)
	loadImageFromTarget(&a.Menu, MENU_PATH)
	loadImageFromTarget(&a.GameOver, GAMEOVER_PATH)
}

func loadImageFromTarget(target **ebiten.Image, path string) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("error to file image: %v", err)
	}

	*target = img
}
