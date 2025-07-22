package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/utils"
)

type GameScene struct {
	gc GameContext

	paused bool

	ground *entities.Ground
	audio  *utils.Audio
	assets *utils.Assets
}

func NewGameScene(gc GameContext) *GameScene {
	gs := &GameScene{}
	gs.gc = gc
	gs.ground = gc.GetGroundEntitie()
	gs.audio = gc.GetAudio()
	gs.assets = gc.GetAssets()

	return gs
}

func (gs *GameScene) Update() {
	gs.pauseGame()

	if gs.paused {
		return
	}

	gs.ground.Update()
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	if gs.paused {
		gs.drawPauseButton(screen)
	}
}

func (gs *GameScene) drawPauseButton(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	imageWidth := float64(gs.assets.Pause.Bounds().Dx())
	imageHeight := float64(gs.assets.Pause.Bounds().Dy())
	screenWidth := float64(screen.Bounds().Dx())
	screenHeight := float64(screen.Bounds().Dy())
	posX := (screenWidth - (imageWidth * 0.05)) / 2
	posY := (screenHeight - (imageHeight * 0.05)) / 2

	op.GeoM.Translate(posX, posY)
	op.GeoM.Scale(0.05, 0.05)

	screen.DrawImage(gs.assets.Pause, op)
}

func (gs *GameScene) pauseGame() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gs.togglePause()
	}
}

func (gs *GameScene) togglePause() {
	gs.paused = !gs.paused
}
