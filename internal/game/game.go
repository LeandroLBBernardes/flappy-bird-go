package game

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"leandro.com/v2/internal/constants"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/scenes"
	"leandro.com/v2/internal/utils"
)

type Game struct {
	scene enums.SceneType
	speed float64

	currentScene scenes.Scene

	ground *entities.Ground

	assets *utils.Assets
	audio  *utils.Audio

	embeddedAssets embed.FS
}

func NewGame(embeddedAssets embed.FS) (ebiten.Game, *Game) {
	g := &Game{}
	g.speed = constants.GAME_SPEED
	g.scene = enums.SceneMenu
	g.embeddedAssets = embeddedAssets
	g.assets = utils.NewAssets(g.embeddedAssets)
	g.audio = utils.NewAudio(g.embeddedAssets)
	g.ground = entities.NewGround(g.assets)
	g.ChangeScene(enums.SceneMenu)

	setScreenProperties(g)
	setGameProperties()
	return g, g
}

func (g *Game) GetAudio() *utils.Audio {
	return g.audio
}

func (g *Game) GetAssets() *utils.Assets {
	return g.assets
}

func (g *Game) GetEmbeddedAssets() embed.FS {
	return g.embeddedAssets
}

func (g *Game) GetGroundEntitie() *entities.Ground {
	return g.ground
}

func (g *Game) Update() error {
	g.currentScene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawBackground(screen)
	g.currentScene.Draw(screen)
	g.ground.Draw(screen)
}

func (g *Game) drawBackground(screen *ebiten.Image) {
	opBg := &ebiten.DrawImageOptions{}
	opBg.GeoM.Scale(1, 1)
	screen.DrawImage(g.assets.BackgroundDay, opBg)
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
