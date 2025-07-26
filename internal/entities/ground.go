package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"leandro.com/v2/internal/constants"
)

const (
	GROUND_SPRITE_PATH = "../../assets/sprites/base.png"
)

type Ground struct {
	sprite      *ebiten.Image
	posX        float64
	spriteWidth int
}

func NewGround() *Ground {
	g := &Ground{}
	g.init()
	return g
}

func (g *Ground) init() {
	var err error
	g.posX = 0
	g.sprite, _, err = ebitenutil.NewImageFromFile(GROUND_SPRITE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	g.spriteWidth = g.sprite.Bounds().Dx()
}

func (g *Ground) Update() {
	if g.posX <= -float64(g.spriteWidth) {
		g.posX += float64(g.spriteWidth)
	} else {
		g.posX -= constants.GAME_SPEED
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
