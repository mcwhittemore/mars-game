package characters

import (
	"sheet"

	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewHero() *Character {
	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{18, 20}, pixel.Vec{0, 2}, 64)
	if err != nil {
		panic(err)
	}

	safe := pixel.R(188, 200, 388, 400)
	second := time.Tick(200 * time.Millisecond)
	hero := NewCharacter(characterSheet, safe.Center(), func(c *Character, dt float64, win MindInput) {
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

		pose, isLeft := c.getPose()
		mov := pose.GetMovement()
		flip := pixel.V(-1, 1)
		if isLeft {
			mov = mov.ScaledXY(flip)
		}

		c.Pos = c.Pos.Add(mov.Scaled(dt))

		isSafe := c.Hits(safe)
		if isSafe {
			return
		}

		selfbox := c.PosBounds(c.Pos)
		_, subject := win.GetCollideRect(selfbox, interface{}(c))

		for subject != nil {
			subject.DropNear(safe.Center(), win.GetCollideRect)
			_, subject = win.GetCollideRect(selfbox, interface{}(c))
		}
	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	hero.AddPose("down", []pixel.Vec{{1, 95}, {2, 95}, {3 + offsetH, 95 - offsetV}, {4 + offsetH, 95}, {0, 95}}, pixel.Vec{0, -200})
	hero.AddPose("side", []pixel.Vec{{1, 96}, {2, 96}, {3 + offsetH, 96 - offsetV}, {4 + offsetH, 96}, {0, 96}}, pixel.Vec{200, 0})
	hero.AddPose("up", []pixel.Vec{{1, 97}, {2, 97}, {3 + offsetH, 97 - offsetV}, {4 + offsetH, 97}, {0, 97}}, pixel.Vec{0, 200})

	hero.ChangePose("down")

	return hero
}
