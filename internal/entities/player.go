package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PLAYER_DOWN_SPRITE_PATH = "../../assets/sprites/{0}bird-downflap.png"
	PLAYER_MID_SPRITE_PATH  = "../../assets/sprites/{0}bird-midflap.png"
	PLAYER_UP_SPRITE_PATH   = "../../assets/sprites/{0}bird-upflap.png"
)

type Player struct {
	posY float64
}

func NewPlayer() *Player {
	p := &Player{}
	p.init()
	return p
}

func (p *Player) init() {

}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {

}
