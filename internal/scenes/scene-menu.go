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
	drawCentralizedImage(ms.assets.Menu, screen)
}

func drawCentralizedImage(image *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	imageWidth := float64(image.Bounds().Dx())
	imageHeight := float64(image.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())
	posX := (screenWidth - imageWidth) / 2
	posY := (screenHeight - imageHeight) / 2

	op.GeoM.Translate(posX, posY)
	op.GeoM.Scale(1, 1)

	screen.DrawImage(image, op)
}
