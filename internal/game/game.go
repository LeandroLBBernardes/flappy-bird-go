package game

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

var (
	randomizeBackground bool = true
	background          *ebiten.Image
)

const (
	GAME_TITLE       = "Flappy Bird"
	SCALE_BACKGROUND = 1.5
	TARGET_RATE      = 60
)

type Game struct {
	scene enums.Scene
	speed float64

	ground *entities.Ground
	assets *utils.Assets
	audio  *utils.Audio
}

func (g *Game) Speed() float64 {
	return g.speed
}

func (g *Game) initGame() {
	g.speed = 0.75
	g.scene = enums.SceneMenu

	g.ground = entities.NewGround()
	g.assets = utils.NewAssets()
	g.audio = utils.NewAudio()

	setScreenProperties(g)
	setGameProperties()
}

func NewGame() (ebiten.Game, *Game) {
	g := &Game{}
	g.initGame()

	fmt.Println("Game Started!")
	fmt.Println("Scene: ", g.scene)
	return g, g
}

func (g *Game) Update() error {
	updateRandomBackground(g)

	switch g.scene {
	case enums.SceneMenu:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if err := g.audio.SwooshPlayer.Rewind(); err != nil {
				return err
			}
			g.audio.SwooshPlayer.Play()
			g.scene = enums.SceneGame
			fmt.Println("Scene: ", g.scene)
		}
	case enums.SceneGame:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if err := g.audio.DiePlayer.Rewind(); err != nil {
				return err
			}
			g.audio.DiePlayer.Play()
			g.scene = enums.SceneGameOver
			fmt.Println("Scene: ", g.scene)
		}
	case enums.SceneGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if err := g.audio.PointPlayer.Rewind(); err != nil {
				return err
			}
			g.audio.PointPlayer.Play()
			g.scene = enums.SceneMenu
			fmt.Println("Scene: ", g.scene)
			randomizeBackground = true
		}
	}

	g.ground.Update(g.speed, g.scene)
	return nil
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

func (g *Game) Draw(screen *ebiten.Image) {
	//lembrando que os canos devem ser desenhados antes do background
	//player deve ser desenhado depois

	//draw background
	opBg := &ebiten.DrawImageOptions{}
	opBg.GeoM.Scale(1, 1)
	screen.DrawImage(background, opBg)

	//draw scenes
	switch g.scene {
	case enums.SceneMenu:
		drawCentralizedImage(g.assets.Menu, screen)
	case enums.SceneGameOver:
		drawCentralizedImage(g.assets.GameOver, screen)
	}

	g.ground.Draw(screen)

	// draw game infos
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	bb := g.assets.BackgroundDay.Bounds()
	return bb.Dx(), bb.Dy()
}

// fazer uma forma dinamica para escalar esse background de acordo
// com o tamanho da tela do usuario
// talvez fazer dinamico para quando o usuario mudar de um monitor para outro
func setScreenProperties(g *Game) {
	bb := g.assets.BackgroundDay.Bounds()

	sw := float64(bb.Dx()) * SCALE_BACKGROUND
	sh := float64(bb.Dy()) * SCALE_BACKGROUND

	ebiten.SetWindowSize(int(sw), int(sh))
	ebiten.SetWindowTitle(GAME_TITLE)
	ebiten.SetWindowIcon([]image.Image{g.assets.Icon})
}

func setGameProperties() {
	ebiten.SetVsyncEnabled(false)
	ebiten.SetTPS(TARGET_RATE)
}
