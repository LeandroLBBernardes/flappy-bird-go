package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/utils"
)

const (
	PIPE_SPRITE_PATH = "../../assets/sprites/pipe-green.png"
	GAP              = 40.0
)

type Pipe struct {
	posY float64
	posX float64

	spriteWidth int
	screenWidth int

	spriteUp   *ebiten.Image
	spriteDown *ebiten.Image
}

func NewPipe(posX float64) *Pipe {
	p := &Pipe{}
	p.init()
	p.spriteWidth = p.spriteUp.Bounds().Dx()
	p.screenWidth = int(posX)

	p.posX = posX
	p.posY = 300.0
	return p
}

func (p *Pipe) init() {
	p.spriteUp = utils.LoadImage(PIPE_SPRITE_PATH)
	p.spriteDown = utils.LoadImage(PIPE_SPRITE_PATH)
}

func (p *Pipe) Update() {
	p.posX -= constants.GAME_SPEED

	// if p.posX <= -float64(p.spriteWidth) {
	// 	p.posX += float64(p.screenWidth + (p.spriteWidth + 100))
	// }
}

func (p *Pipe) Draw(screen *ebiten.Image) {
	p.drawPipe(screen, p.spriteUp, -1)
	p.drawPipe(screen, p.spriteDown, 1)
}

func (p *Pipe) drawPipe(screen *ebiten.Image, sprite *ebiten.Image, scale float64) {
	op := &ebiten.DrawImageOptions{}

	screenHeight := float64(screen.Bounds().Dy())
	posY := screenHeight - (p.posY + (-GAP * scale))

	op.GeoM.Scale(1.0, 1.0*scale)

	op.GeoM.Translate(p.posX, posY)
	screen.DrawImage(sprite, op)
}
