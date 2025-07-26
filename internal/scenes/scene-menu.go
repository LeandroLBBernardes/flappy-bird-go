package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

type MenuScene struct {
	gc     GameContext
	audio  *utils.Audio
	assets *utils.Assets
}

func NewMenuScene(gc GameContext) *MenuScene {
	ms := &MenuScene{}
	ms.gc = gc
	ms.audio = gc.GetAudio()
	ms.assets = gc.GetAssets()

	return ms
}

func (ms *MenuScene) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		ms.audio.PlayOnce(enums.SwooshAudio)
		ms.gc.ChangeScene(enums.SceneGame)
	}
}

func (ms *MenuScene) Draw(screen *ebiten.Image) {
	utils.DrawCentralizedImage(ms.assets.Menu, screen)
}
