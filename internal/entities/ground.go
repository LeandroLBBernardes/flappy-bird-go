package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/utils"
)

const (
	GROUND_SPRITE_PATH = "assets/sprites/base.png"
)

type Ground struct {
	sprite      *ebiten.Image
	posX        float64
	spriteWidth int

	assets *utils.Assets
}

func NewGround(assets *utils.Assets) *Ground {
	g := &Ground{}
	g.assets = assets
	g.posX = 0
	g.sprite = g.assets.LoadImage(GROUND_SPRITE_PATH)
	g.spriteWidth = g.sprite.Bounds().Dx()
	return g
}

func (g *Ground) GetGroundHeight() float64 {
	return float64(g.sprite.Bounds().Dy())
}

func (g *Ground) Update() {
	move := constants.GAME_SPEED * constants.DELTA_TIME

	if g.posX <= -float64(g.spriteWidth) {
		g.posX += float64(g.spriteWidth)
	} else {
		g.posX -= move
	}
}

func (g *Ground) Draw(screen *ebiten.Image) {
	posY := float64(screen.Bounds().Dy() - g.sprite.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())

	for x := g.posX; x < screenWidth; x += float64(g.spriteWidth) {
		opG := &ebiten.DrawImageOptions{}
		opG.GeoM.Reset()
		opG.GeoM.Translate(x, posY)
		screen.DrawImage(g.sprite, opG)
	}
}
