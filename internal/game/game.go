package game

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/scenes"
	"leandro.com/v2/internal/utils"
)

var (
	randomizeBackground bool = true
	background          *ebiten.Image
)

type Game struct {
	scene        enums.SceneType
	speed        float64
	currentScene scenes.Scene

	ground *entities.Ground

	assets *utils.Assets
	audio  *utils.Audio
}

func NewGame() (ebiten.Game, *Game) {
	g := &Game{}
	g.initGame()

	fmt.Println("Game Started!")
	return g, g
}

func (g *Game) initGame() {
	g.speed = constants.GAME_SPEED
	g.scene = enums.SceneMenu

	g.ground = entities.NewGround()
	g.assets = utils.NewAssets()
	g.audio = utils.NewAudio()

	g.ChangeScene(enums.SceneMenu)

	setScreenProperties(g)
	setGameProperties()
}

func (g *Game) GetAudio() *utils.Audio {
	return g.audio
}

func (g *Game) GetAssets() *utils.Assets {
	return g.assets
}

func (g *Game) GetGroundEntitie() *entities.Ground {
	return g.ground
}

func (g *Game) Update() error {
	updateRandomBackground(g)

	g.currentScene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBackground(screen)
	g.currentScene.Draw(screen)

	g.ground.Draw(screen)
}

func (g *Game) ChangeScene(sceneType enums.SceneType) {
	g.currentScene = scenes.SceneFactory(sceneType, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	bb := g.assets.BackgroundDay.Bounds()
	return bb.Dx(), bb.Dy()
}

func setScreenProperties(g *Game) {
	bb := g.assets.BackgroundDay.Bounds()

	sw := float64(bb.Dx()) * constants.SCALE_BACKGROUND
	sh := float64(bb.Dy()) * constants.SCALE_BACKGROUND

	ebiten.SetWindowSize(int(sw), int(sh))
	ebiten.SetWindowTitle(constants.GAME_TITLE)
	ebiten.SetWindowIcon([]image.Image{g.assets.Icon})
}

func setGameProperties() {
	ebiten.SetVsyncEnabled(false)
	ebiten.SetTPS(constants.TARGET_RATE)
}

func updateRandomBackground(g *Game) {
	if randomizeBackground {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if r.Intn(2) == 0 {
			background = g.assets.BackgroundDay
		} else {
			background = g.assets.BackgroundNight
		}
	}
	randomizeBackground = false
}

func drawBackground(screen *ebiten.Image) {
	opBg := &ebiten.DrawImageOptions{}
	opBg.GeoM.Scale(1, 1)
	screen.DrawImage(background, opBg)
}
