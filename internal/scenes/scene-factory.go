package scenes

import (
	"log"

	"leandro.com/v2/internal/enums"
)

func SceneFactory(sceneType enums.SceneType, gc GameContext) Scene {
	switch sceneType {
	case enums.SceneMenu:
		return NewMenuScene(gc)
	case enums.SceneGame:
		return NewGameScene(gc)
	default:
		log.Fatalf("invalid scene type")
		return nil
	}
}
