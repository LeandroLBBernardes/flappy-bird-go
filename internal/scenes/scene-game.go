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

	screenWidth  int
	screenheight int

	audio  *utils.Audio
	assets *utils.Assets
}

func NewGameScene(gc GameContext) *GameScene {
	gs := &GameScene{}

	gs.gc = gc
	gs.ground = gc.GetGroundEntitie()
	gs.audio = gc.GetAudio()
	gs.assets = gc.GetAssets()

	gs.screenWidth = gs.assets.BackgroundDay.Bounds().Dx()
	gs.screenheight = gs.assets.BackgroundDay.Bounds().Dy()

	gs.counter = entities.NewCounter()
	gs.player = entities.NewPlayer(gs.screenWidth, gs.screenheight)

	for i := 0; i < PIPE_COUNT; i++ {
		gs.pipes[i] = entities.NewPipe(float64(gs.screenWidth+i*constants.PIPE_SPACING), gs.screenWidth)
	}

	return gs
}

func (gs *GameScene) Update() {
	gs.gameOver()
	if gs.isGameOver {
		return
	}

	gs.pauseGame()
	if gs.isPaused {
		return
	}

	groundCollisionHeight := float64(gs.screenheight) - gs.ground.GetGroundHeight()

	gs.player.Update(gs.audio, groundCollisionHeight, gs)
	gs.ground.Update()

	for _, pipe := range gs.pipes {
		pipe.Update()

		pipeRectUp, pipeRectDown := pipe.GetCollisionRects(float64(gs.screenheight))

		if gs.player.GetCollisionRect().Overlaps(pipeRectUp) || gs.player.GetCollisionRect().Overlaps(pipeRectDown) {
			gs.audio.PlayOnce(enums.HitAudio)
			gs.audio.PlayOnce(enums.DieAudio)
			gs.SetGameOver()
			break
		}

		if !pipe.AlreadyPassed && gs.player.PosX+(gs.player.SpriteWidth/2) > pipe.PosX+float64(pipe.SpriteWidth)/2 {
			pipe.AlreadyPassed = true
			gs.counter.PlusValue()
			gs.audio.PlayOnce(enums.PointAudio)
		}
	}
}

func (gs *GameScene) Draw(screen *ebiten.Image) {
	gs.player.Draw(screen)

	for _, pipe := range gs.pipes {
		pipe.Draw(screen)
	}
	gs.counter.Draw(screen)

	if gs.isGameOver {
		utils.DrawCentralizedImage(gs.assets.GameOver, screen)
	}
}

func (gs *GameScene) pauseGame() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		gs.isPaused = !gs.isPaused
	}
}

func (gs *GameScene) gameOver() {
	if gs.isGameOver && (inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) {
		gs.gc.ChangeScene(enums.SceneMenu)
	}
}

func (gs *GameScene) SetGameOver() {
	gs.isGameOver = true
}
