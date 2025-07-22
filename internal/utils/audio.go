package utils

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"leandro.com/v2/internal/enums"
)

const (
	SAMPLE_RATE = 44100

	SWOOSH_AUDIO_PATH = "../../assets/audio/swoosh.wav"
	POINT_AUDIO_PATH  = "../../assets/audio/point.wav"
	DIE_AUDIO_PATH    = "../../assets/audio/die.wav"
	WING_AUDIO_PATH   = "../../assets/audio/wing.wav"
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
	a.loadAudios()
	return a
}

func (a *Audio) PlayOnce(audioType enums.AudioType) {
	player := audioPlayerFactory(a, audioType)

	if err := player.Rewind(); err != nil {
		log.Fatalf("error to play audio: %v", err)
	}

	player.Play()
}

func (a *Audio) loadAudios() {
	a.audioContext = audio.NewContext(SAMPLE_RATE)

	loadAudio(&a.SwooshPlayer, a.audioContext, SWOOSH_AUDIO_PATH)
	loadAudio(&a.PointPlayer, a.audioContext, POINT_AUDIO_PATH)
	loadAudio(&a.DiePlayer, a.audioContext, DIE_AUDIO_PATH)
	loadAudio(&a.WingPlayer, a.audioContext, WING_AUDIO_PATH)
}

func loadAudio(target **audio.Player, context *audio.Context, path string) {
	wavFile := setAudio(path)
	wavStream := decodeAudio(wavFile)
	setPlayer(target, context, wavStream)
}

func setAudio(path string) *os.File {
	wavFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("error to file audio: %v", err)
	}

	return wavFile
}

func decodeAudio(wavFile *os.File) *wav.Stream {
	wavStream, err := wav.DecodeF32(wavFile)
	if err != nil {
		log.Fatalf("error to file audio: %v", err)
	}

	return wavStream
}

func setPlayer(target **audio.Player, context *audio.Context, wavStream *wav.Stream) {
	player, err := context.NewPlayerF32(wavStream)
	if err != nil {
		log.Fatalf("error to set player: %v", err)
	}

	*target = player
}

func audioPlayerFactory(a *Audio, audioType enums.AudioType) *audio.Player {
	switch audioType {
	case enums.SwooshAudio:
		return a.SwooshPlayer
	case enums.PointAudio:
		return a.PointPlayer
	case enums.DieAudio:
		return a.DiePlayer
	case enums.WingAudio:
		return a.WingPlayer
	default:
		log.Fatalf("invalid audio type")
		return nil
	}
}
