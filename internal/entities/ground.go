package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"leandro.com/v2/internal/enums"
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

func (g *Ground) Update(gSpeed float64, gScene enums.Scene) {
	if gScene == enums.SceneGame {
		g.posX = g.posX - gSpeed

		// a ideia é, quando o sprite sair da tela, reiniciar a posicao para o inicio
		// dando ideia de looping
		if g.posX <= -float64(g.spriteWidth) {
			g.posX += float64(g.spriteWidth)
		}
	}
}

func (g *Ground) Draw(screen *ebiten.Image) {
	posY := float64(screen.Bounds().Dy() - g.sprite.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())

	// x inicialmente começa com 0
	// screenWidth tem uns 640px
	// desenhha quantos sprites forem precisos para cobrir a tela
	// e continuar com a ideia de infinito
	for x := g.posX; x < screenWidth; x += float64(g.spriteWidth) {
		opG := &ebiten.DrawImageOptions{}
		opG.GeoM.Reset()
		opG.GeoM.Translate(x, posY)
		screen.DrawImage(g.sprite, opG)
	}
}
