package utils

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	SAMPLE_RATE = 44100
)

type Audio struct {
	audioContext *audio.Context

	SwooshPlayer *audio.Player
	PointPlayer  *audio.Player
	DiePlayer    *audio.Player
	WingPlayer   *audio.Player
}

func NewAudio() *Audio {
	a := &Audio{}
	a.init()
	return a
}

func (a *Audio) init() {
	a.audioContext = audio.NewContext(SAMPLE_RATE)

	swooshWav, err := os.Open("../../assets/audio/swoosh.wav")
	if err != nil {
		log.Fatal(err)
	}

	pointWav, err := os.Open("../../assets/audio/point.wav")
	if err != nil {
		log.Fatal(err)
	}

	dieWav, err := os.Open("../../assets/audio/die.wav")
	if err != nil {
		log.Fatal(err)
	}

	wingWav, err := os.Open("../../assets/audio/wing.wav")
	if err != nil {
		log.Fatal(err)
	}

	// Decodifica o áudio WAV
	swooshD, err := wav.DecodeF32(swooshWav)
	if err != nil {
		log.Fatal(err)
	}

	pointD, err := wav.DecodeF32(pointWav)
	if err != nil {
		log.Fatal(err)
	}

	dieD, err := wav.DecodeF32(dieWav)
	if err != nil {
		log.Fatal(err)
	}

	wingD, err := wav.DecodeF32(wingWav)
	if err != nil {
		log.Fatal(err)
	}

	// seta os players
	a.SwooshPlayer, err = a.audioContext.NewPlayerF32(swooshD)
	if err != nil {
		log.Fatal(err)
	}

	a.PointPlayer, err = a.audioContext.NewPlayerF32(pointD)
	if err != nil {
		log.Fatal(err)
	}

	a.DiePlayer, err = a.audioContext.NewPlayerF32(dieD)
	if err != nil {
		log.Fatal(err)
	}

	a.WingPlayer, err = a.audioContext.NewPlayerF32(wingD)
	if err != nil {
		log.Fatal(err)
	}
}
