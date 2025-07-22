package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"leandro.com/v2/internal/entities"
	"leandro.com/v2/internal/enums"
	"leandro.com/v2/internal/utils"
)

type GameContext interface {
	ChangeScene(sceneType enums.SceneType)
	GetAudio() *utils.Audio
	GetAssets() *utils.Assets
	GetGroundEntitie() *entities.Ground
}

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}
