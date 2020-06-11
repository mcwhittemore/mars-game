package game

import (
	"app/characters"
	"app/maps"

	"github.com/faiface/pixel"
)

type Scene interface {
	Enter(mi characters.MindInput)
	Exit(mi characters.MindInput)
	Update(dt float64, mind characters.MindInput)
	GetCamera() pixel.Vec
	SetCamera(pixel.Vec)
	GetMap() *maps.Map
}

type SceneBuilder func() Scene

type SceneManager struct {
	Current Scene
	scenes  map[string]SceneBuilder
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		Current: nil,
		scenes:  make(map[string]SceneBuilder, 0),
	}
}

func (s *SceneManager) AddScene(name string, sb SceneBuilder) {
	s.scenes[name] = sb
}

func (s *SceneManager) ChangeScene(name string, mi characters.MindInput) {
	if s.Current != nil {
		s.Current.Exit(mi)
	}
	s.Current = s.scenes[name]()
	s.Current.Enter(mi)
}
