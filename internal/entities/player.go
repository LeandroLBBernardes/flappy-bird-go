package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	playerDownflapSprite = "assets/sprites/{0}bird-downflap.png"
	playerMidflapSprite  = "assets/sprites/{0}bird-midflap.png"
	playerUpflapSprite   = "assets/sprites/{0}bird-upflap.png"
)

// TODO: Criar o player
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
