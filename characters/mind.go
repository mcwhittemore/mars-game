package characters

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type MindInput interface {
	JustPressed(pixelgl.Button) bool
	Pressed(pixelgl.Button) bool
	GetCollideRect(pixel.Rect, interface{}) (pixel.Rect, *Character)
	KeepInView(pixel.Vec, pixel.Vec, float64)
}

type MindFunc func(*Character, float64, MindInput)
