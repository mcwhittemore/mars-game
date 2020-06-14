package characters

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewHeroDefault(pos pixel.Vec) *Character {

	second := time.Tick(200 * time.Millisecond)
	hero := NewCharacter(characterSheet, func(cd *CharacterData, dt float64, mi MindInput) {
		c := cd.Character
		if mi.JustPressed(pixelgl.KeyD) {
			c.ChangePose("side")
		} else if mi.JustPressed(pixelgl.KeyA) {
			c.ChangePose("left-side")
		} else if mi.JustPressed(pixelgl.KeyS) {
			c.ChangePose("down")
		} else if mi.JustPressed(pixelgl.KeyW) {
			c.ChangePose("up")
		} else if mi.JustPressed(pixelgl.KeyJ) {
			dir := c.GetDirection()
			tp := c.Pos.Add(dir.Scaled(64))
			fmt.Println(dir, tp)
			item := mi.GetItem(tp)
			if item != nil && item.CanPickUp() {
				mi.RemoveItem(item)
				cd.AddItem(item.Name)
			}
		}

		if mi.Pressed(pixelgl.KeyD) || mi.Pressed(pixelgl.KeyA) || mi.Pressed(pixelgl.KeyS) || mi.Pressed(pixelgl.KeyW) {
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

		if mi.IsObstacle(nextPos) == false {
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
