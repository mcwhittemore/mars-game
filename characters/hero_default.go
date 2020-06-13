package characters

import (
	"app/sheet"

	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewHeroDefault(pos pixel.Vec) *Character {
	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{X: 18, Y: 20}, pixel.Vec{X: 0, Y: 2}, 64)
	if err != nil {
		panic(err)
	}

	second := time.Tick(200 * time.Millisecond)
	hero := NewCharacter(characterSheet, func(c *Character, dt float64, win MindInput) {
		if win.JustPressed(pixelgl.KeyD) {
			c.ChangePose("side")
		} else if win.JustPressed(pixelgl.KeyA) {
			c.ChangePose("left-side")
		} else if win.JustPressed(pixelgl.KeyS) {
			c.ChangePose("down")
		} else if win.JustPressed(pixelgl.KeyW) {
			c.ChangePose("up")
		}

		if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyW) {
			select {
			case <-second:
				c.Step()
			default:
			}
		} else {
			c.Stop()
		}

		pose, isLeft := c.GetPose()
		mov := pose.GetMovement()
		flip := pixel.V(-1, 1)
		if isLeft {
			mov = mov.ScaledXY(flip)
		}

		nextPos := c.Pos.Add(mov.Scaled(dt))

		if win.IsObstacle(nextPos) == false {
			c.Pos = nextPos
		}
	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	hero.AddPose("down", []pixel.Vec{{X: 1, Y: 95}, {X: 2, Y: 95}, {X: 3 + offsetH, Y: 95 - offsetV}, {X: 4 + offsetH, Y: 95}, {X: 0, Y: 95}}, pixel.Vec{X: 0, Y: -200})
	hero.AddPose("side", []pixel.Vec{{X: 1, Y: 96}, {X: 2, Y: 96}, {X: 3 + offsetH, Y: 96 - offsetV}, {X: 4 + offsetH, Y: 96}, {X: 0, Y: 96}}, pixel.Vec{X: 200, Y: 0})
	hero.AddPose("up", []pixel.Vec{{X: 1, Y: 97}, {X: 2, Y: 97}, {X: 3 + offsetH, Y: 97 - offsetV}, {X: 4 + offsetH, Y: 97}, {X: 0, Y: 97}}, pixel.Vec{X: 0, Y: 200})

	hero.ChangePose("down")
	hero.Pos = pos

	return hero
}
