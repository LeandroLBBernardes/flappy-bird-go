package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

type GameScene struct {
	gc GameContext

	isPaused   bool
	isGameOver bool

	ground  *entities.Ground
	counter *entities.Counter
	audio   *utils.Audio
	assets  *utils.Assets
}

func NewGameScene(gc GameContext) *GameScene {
	gs := &GameScene{}
	gs.gc = gc
	gs.ground = gc.GetGroundEntitie()
	gs.audio = gc.GetAudio()
	gs.assets = gc.GetAssets()
	gs.counter = entities.NewCounter()

	return gs
}

func (gs *GameScene) Update() {
	if gs.isGameOver {
		gs.gc.ChangeScene(enums.SceneMenu)
	}

	gs.pauseGame()

	if gs.isPaused {
		return
	}

	gs.gameOver()
	gs.ground.Update()
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	gs.counter.Draw(screen)
}

func (gs *GameScene) pauseGame() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gs.isPaused = !gs.isPaused
	}
}

func (gs *GameScene) gameOver() {
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		gs.isGameOver = true
	}
}
