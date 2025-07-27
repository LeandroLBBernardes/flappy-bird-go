package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawCentralizedImage(image *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	imageWidth := float64(image.Bounds().Dx())
	imageHeight := float64(image.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())
	posX := (screenWidth - imageWidth) / 2
	posY := (screenHeight - imageHeight - 112) / 2

	op.GeoM.Translate(posX, posY)
	op.GeoM.Scale(1, 1)

	screen.DrawImage(image, op)
}
