package main

type Scene int

const (
	SceneMenu Scene = iota
	SceneGame
	SceneGameOver
)

func (s Scene) String() string {
	switch s {
	case SceneMenu:
		return "SceneMenu"
	case SceneGame:
		return "SceneGame"
	case SceneGameOver:
		return "SceneGameOver"
	default:
		return "Unknown"
	}
}
