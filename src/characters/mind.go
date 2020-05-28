package characters

import (
	"github.com/faiface/pixel/pixelgl"
)

type MindInput interface {
	JustPressed(pixelgl.Button) bool
	Pressed(pixelgl.Button) bool
}

type MindFunc func(*Character, float64, MindInput)
