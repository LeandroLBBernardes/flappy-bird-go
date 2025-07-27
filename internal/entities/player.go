package entities

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

const (
	PLAYER_DOWN_SPRITE_PATH = "assets/sprites/yellowbird-downflap.png"
	PLAYER_MID_SPRITE_PATH  = "assets/sprites/yellowbird-midflap.png"
	PLAYER_UP_SPRITE_PATH   = "assets/sprites/yellowbird-upflap.png"
	GRAVITY                 = 800.0
	JUMP_FORCE              = 250.0
)

type Player struct {
	posY float64
	PosX float64
	velY float64

	screenHeight int
	screenWidth  int

	SpriteWidth  float64
	SpriteHeight float64

	sprites [3]*ebiten.Image

	imageIndex     int
	rotation       float64
	tiltTimer      float64
	targetRotation float64
	animationTimer float64

	assets *utils.Assets
}

func NewPlayer(screenWidth int, screenHeight int, assets *utils.Assets) *Player {
	p := &Player{}
	p.assets = assets
	p.loadSprites()
	p.screenHeight = screenHeight
	p.screenWidth = screenWidth
	p.SpriteHeight = float64(p.sprites[0].Bounds().Dy())
	p.posY = (float64(screenHeight) - p.SpriteHeight) / 2
	p.PosX = 40
	p.SpriteWidth = float64(p.sprites[0].Bounds().Dx())
	return p
}

func (p *Player) loadSprites() {
	p.sprites[0] = p.assets.LoadImage(PLAYER_MID_SPRITE_PATH)
	p.sprites[1] = p.assets.LoadImage(PLAYER_DOWN_SPRITE_PATH)
	p.sprites[2] = p.assets.LoadImage(PLAYER_UP_SPRITE_PATH)
}

func (p *Player) Update(audio *utils.Audio, groundCollisionHeight float64, gc GameContext) {
	spriteHeight := float64(p.sprites[p.imageIndex].Bounds().Dy())

	p.addAnimation()
	p.addRotation()

	p.checkSkyCollision(spriteHeight, audio, gc)
	p.checkGroundCollision(spriteHeight, groundCollisionHeight, audio, gc)

	p.addGravity()

	p.actions(audio)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest

	w, h := p.sprites[p.imageIndex].Bounds().Dx(), p.sprites[p.imageIndex].Bounds().Dy()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(math.Floor(p.PosX+float64(w)/2), math.Floor(p.posY+float64(h)/2))

	screen.DrawImage(p.sprites[p.imageIndex], op)
}

func (p *Player) addAnimation() {
	p.animationTimer += constants.DELTA_TIME

	if p.animationTimer >= 0.1 {
		p.imageIndex++
		if p.imageIndex > 2 {
			p.imageIndex = 0
		}
		p.animationTimer = 0
	}
}

func (p *Player) addRotation() {
	if p.tiltTimer > 0 {
		p.tiltTimer -= constants.DELTA_TIME
		p.targetRotation = -math.Pi / 6
	} else {
		p.targetRotation = math.Pi / 4
	}

	rotationSpeed := 0.1
	p.rotation += (p.targetRotation - p.rotation) * rotationSpeed
}

func (p *Player) addGravity() {
	p.velY += GRAVITY * constants.DELTA_TIME
	p.posY += p.velY * constants.DELTA_TIME
}

func (p *Player) actions(audio *utils.Audio) {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		audio.PlayOnce(enums.WingAudio)
		p.velY = -JUMP_FORCE
		p.tiltTimer = 0.3
	}
}

func (p *Player) checkGroundCollision(spriteHeight float64, groundCollisionHeight float64, audio *utils.Audio, gc GameContext) {
	if p.posY+spriteHeight >= groundCollisionHeight {
		p.posY = groundCollisionHeight - spriteHeight
		p.velY = 0
		audio.PlayOnce(enums.HitAudio)
		audio.PlayOnce(enums.DieAudio)
		gc.SetGameOver()
	}
}

func (p *Player) checkSkyCollision(spriteHeight float64, audio *utils.Audio, gc GameContext) {
	if p.posY <= -(1.5 * spriteHeight) {
		audio.PlayOnce(enums.HitAudio)
		audio.PlayOnce(enums.DieAudio)
		gc.SetGameOver()
	}
}

func (p *Player) GetCollisionRect() image.Rectangle {
	return image.Rect(
		int(p.PosX),
		int(p.posY),
		int(p.PosX+p.SpriteWidth),
		int(p.posY+p.SpriteHeight),
	)
}
