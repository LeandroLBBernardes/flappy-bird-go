package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PIPE_SPRITE_PATH = "../../assets/sprites/pipe-green.png"
)

type Pipe struct {
	posY float64
	posx float64
}

func NewPipe() *Pipe {
	p := &Pipe{}
	p.init()
	return p
}

func (p *Pipe) init() {

}

func (p *Pipe) Update() {

}

func (p *Pipe) Draw(screen *ebiten.Image) {

}
