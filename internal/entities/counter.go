package entities

import (
	"fmt"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/utils"
)

type Counter struct {
	sprites [10]*ebiten.Image

	value int

	assets *utils.Assets
}

func NewCounter(assets *utils.Assets) *Counter {
	c := &Counter{}
	c.assets = assets
	c.value = 0
	c.loadSprites()
	return c
}

func (c *Counter) loadSprites() {
	for i := range 10 {
		path := fmt.Sprintf("%s/%d.png", constants.SPRITE_PATH, i)
		c.sprites[i] = c.assets.LoadImage(path)
	}
}

func (c *Counter) PlusValue() {
	c.value += 1
}

func (c *Counter) Draw(screen *ebiten.Image) {
	strValue := strconv.Itoa(c.value)

	totalWidth := calcTotalImageWidth(strValue, c.sprites)

	screenWidth := float64(screen.Bounds().Dx())

	posX := (screenWidth - totalWidth) / 2
	posY := 20.0

	for _, char := range strValue {
		digit := int(char - '0')

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(posX, posY)
		op.GeoM.Scale(1, 1)

		imageWidth := float64(c.sprites[digit].Bounds().Dx())

		screen.DrawImage(c.sprites[digit], op)

		posX += imageWidth
	}
}

func calcTotalImageWidth(strValue string, sprites [10]*ebiten.Image) float64 {
	totalWidth := 0.0

	for _, char := range strValue {
		digit := int(char - '0')
		totalWidth += float64(sprites[digit].Bounds().Dx())
	}

	return totalWidth
}
