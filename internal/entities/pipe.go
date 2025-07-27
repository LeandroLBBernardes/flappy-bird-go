package entities

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/utils"
)

const (
	PIPE_SPRITE_PATH = "../../assets/sprites/pipe-green.png"
	GAP              = 38
)

type Pipe struct {
	posY float64
	PosX float64

	SpriteWidth int
	screenWidth int

	spriteUp   *ebiten.Image
	spriteDown *ebiten.Image

	AlreadyPassed bool
}

func NewPipe(posX float64, screenWidth int) *Pipe {
	p := &Pipe{}
	p.init()
	p.SpriteWidth = p.spriteUp.Bounds().Dx()
	p.screenWidth = screenWidth

	p.PosX = posX
	p.posY = randomHeight()
	return p
}

func (p *Pipe) init() {
	p.spriteUp = utils.LoadImage(PIPE_SPRITE_PATH)
	p.spriteDown = utils.LoadImage(PIPE_SPRITE_PATH)
}

func (p *Pipe) Update() {
	move := constants.GAME_SPEED * constants.DELTA_TIME

	if p.PosX <= -float64(p.SpriteWidth) {
		p.posY = randomHeight()
		p.PosX = float64(p.screenWidth + constants.PIPE_SPACING + (p.SpriteWidth / 2))
		p.AlreadyPassed = false
	} else {
		p.PosX -= move
	}
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

	op.GeoM.Translate(p.PosX, posY)
	screen.DrawImage(sprite, op)
}

func randomHeight() float64 {
	min := 190.0
	max := 410.0
	return min + rand.Float64()*(max-min)
}

func (p *Pipe) GetCollisionRects(screenHeight float64) (image.Rectangle, image.Rectangle) {
	rectUp := image.Rect(
		int(p.PosX),
		int(screenHeight-p.posY-GAP),
		int(p.PosX)+p.SpriteWidth,
		0,
	)

	rectDown := image.Rect(
		int(p.PosX),
		int(screenHeight-p.posY+GAP),
		int(p.PosX)+p.SpriteWidth,
		int(screenHeight),
	)

	return rectUp, rectDown
}
