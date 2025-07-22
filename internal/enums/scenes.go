package enums

type SceneType int

const (
	SceneMenu SceneType = iota
	SceneGame
	SceneGameOver
	ScenePause
)

func (s SceneType) String() string {
	switch s {
	case SceneMenu:
		return "SceneMenu"
	case SceneGame:
		return "SceneGame"
	case SceneGameOver:
		return "SceneGameOver"
	case ScenePause:
		return "ScenePause"
	default:
		return "Unknown"
	}
}
