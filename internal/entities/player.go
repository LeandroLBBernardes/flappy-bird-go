package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

const (
	PLAYER_DOWN_SPRITE_PATH = "../../assets/sprites/yellowbird-downflap.png"
	PLAYER_MID_SPRITE_PATH  = "../../assets/sprites/yellowbird-midflap.png"
	PLAYER_UP_SPRITE_PATH   = "../../assets/sprites/yellowbird-upflap.png"
)

type Player struct {
	posY float64
	posX float64

	screenHeight int
	screenWidth  int

	sprites [3]*ebiten.Image

	imageIndex     int
	animationTimer float64
}

func NewPlayer(screenWidth int, screenHeight int) *Player {
	p := &Player{}

	p.loadSprites()

	p.screenHeight = screenHeight
	p.screenWidth = screenWidth

	imageHeight := float64(p.sprites[0].Bounds().Dy())
	p.posY = (float64(screenHeight) - imageHeight) / 2
	p.posX = 40

	return p
}

func (p *Player) loadSprites() {
	p.sprites[0] = utils.LoadImage(PLAYER_MID_SPRITE_PATH)
	p.sprites[1] = utils.LoadImage(PLAYER_DOWN_SPRITE_PATH)
	p.sprites[2] = utils.LoadImage(PLAYER_UP_SPRITE_PATH)
}

func (p *Player) Update(audio *utils.Audio) {
	p.animation()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		audio.PlayOnce(enums.WingAudio)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.posX, p.posY)
	op.GeoM.Scale(1, 1)

	screen.DrawImage(p.sprites[p.imageIndex], op)
}

func (p *Player) animation() {
	p.animationTimer += constants.DELTA_TIME

	if p.animationTimer >= 0.1 {
		p.imageIndex++
		if p.imageIndex > 2 {
			p.imageIndex = 0
		}
		p.animationTimer = 0
	}
}
