package utils

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"leandro.com/v2/internal/enums"
)

const (
	SAMPLE_RATE = 44100

	SWOOSH_AUDIO_PATH = "assets/audio/swoosh.wav"
	POINT_AUDIO_PATH  = "assets/audio/point.wav"
	DIE_AUDIO_PATH    = "assets/audio/die.wav"
	WING_AUDIO_PATH   = "assets/audio/wing.wav"
	HIT_AUDIO_PATH    = "assets/audio/hit.wav"
)

type Audio struct {
	audioContext *audio.Context

	SwooshPlayer *audio.Player
	PointPlayer  *audio.Player
	DiePlayer    *audio.Player
	WingPlayer   *audio.Player
	HitPlayer    *audio.Player
}

func NewAudio(embeddedAssets embed.FS) *Audio {
	a := &Audio{}
	a.loadAudios(embeddedAssets)
	return a
}

func (a *Audio) PlayOnce(audioType enums.AudioType) {
	player := audioPlayerFactory(a, audioType)
	player.SetVolume(0.1)

	if err := player.Rewind(); err != nil {
		log.Fatalf("error to play audio: %v", err)
	}

	player.Play()
}

func (a *Audio) loadAudios(embeddedAssets embed.FS) {
	a.audioContext = audio.NewContext(SAMPLE_RATE)

	loadEmbeddedAudio(&a.SwooshPlayer, a.audioContext, embeddedAssets, SWOOSH_AUDIO_PATH)
	loadEmbeddedAudio(&a.PointPlayer, a.audioContext, embeddedAssets, POINT_AUDIO_PATH)
	loadEmbeddedAudio(&a.DiePlayer, a.audioContext, embeddedAssets, DIE_AUDIO_PATH)
	loadEmbeddedAudio(&a.WingPlayer, a.audioContext, embeddedAssets, WING_AUDIO_PATH)
	loadEmbeddedAudio(&a.HitPlayer, a.audioContext, embeddedAssets, HIT_AUDIO_PATH)
}

func loadEmbeddedAudio(target **audio.Player, context *audio.Context, embeddedAssets embed.FS, path string) {
	file, err := embeddedAssets.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading embedded audio file %s: %v", path, err)
	}

	wavStream, err := wav.DecodeF32(bytes.NewReader(file))
	if err != nil {
		log.Fatalf("error decoding audio %s: %v", path, err)
	}

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
	case enums.HitAudio:
		return a.HitPlayer
	default:
		log.Fatalf("invalid audio type")
		return nil
	}
}
