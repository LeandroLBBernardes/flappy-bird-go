package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

const PIPE_COUNT = 3

type GameScene struct {
	gc GameContext

	isPaused   bool
	isGameOver bool

	ground  *entities.Ground
	counter *entities.Counter
	player  *entities.Player
	pipes   [PIPE_COUNT]*entities.Pipe

	audio  *utils.Audio
	assets *utils.Assets
}

func NewGameScene(gc GameContext) *GameScene {
	gs := &GameScene{}

	gs.gc = gc
	gs.ground = gc.GetGroundEntitie()
	gs.audio = gc.GetAudio()
	gs.assets = gc.GetAssets()

	screenWidth := gs.assets.BackgroundDay.Bounds().Dx()
	screenheight := gs.assets.BackgroundDay.Bounds().Dy()

	gs.counter = entities.NewCounter()
	gs.player = entities.NewPlayer(screenWidth, screenheight)

	for i := 0; i < PIPE_COUNT; i++ {
		gs.pipes[i] = entities.NewPipe(float64(screenWidth+i*constants.PIPE_SPACING), screenWidth)
	}

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

	gs.player.Update(gs.audio)
	gs.ground.Update()

	for _, pipe := range gs.pipes {
		pipe.Update()
	}
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	gs.player.Draw(screen)

	for _, pipe := range gs.pipes {
		pipe.Draw(screen)
	}
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
