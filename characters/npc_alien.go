package characters

import (
	"math"
	"time"

	"github.com/faiface/pixel"
)

func NewAlien(pos pixel.Vec) *Character {
	startRow := float64(23)

	second := time.Tick(200 * time.Millisecond)

	dirPose := [4]string{"down", "side", "up", "left-side"}

	dir := 0 // 0 = down, 1 = right, 2 = up, 3 = left

	alien := NewCharacter(characterSheet, func(cd *CharacterData, dt float64, win MindInput) {
		c := cd.Character
		hp := win.GetHeroPos()
		dist := pixel.L(hp, c.Pos).Len()
		dx := math.Abs(hp.X - c.Pos.X)
		dy := math.Abs(hp.Y - c.Pos.Y)
		mov := c.GetMovement()
		nextPos := c.Pos.Add(mov.Scaled(dt))
		select {
		case <-second:
			if win.IsObstacle(nextPos) == false {
				c.Step()
			} else {
				c.Stop()
			}
		default:
		}

		if dist > 5000 {
			c.Stop()
		} else {
			mov := c.GetMovement()
			nextPos := c.Pos.Add(mov.Scaled(dt))
			if win.IsObstacle(nextPos) == false {
				c.Pos = nextPos
				return
			}

			var scores [4]int

			if dx < dy {
				scores[1]++
				scores[3]++
			} else {
				scores[0]++
				scores[2]++
			}

			hs := dx / (hp.X - c.Pos.X)
			vs := dy / (hp.Y - c.Pos.Y)
			scores[0] += int(vs)
			scores[2] += int(vs * -1)
			scores[1] += int(hs * -1)
			scores[3] += int(hs)

			scores[dir] = -5
			for i := 0; i < 4; i++ {
				c.ChangePose(dirPose[i])
				c.Step()
				mov := c.GetMovement()
				nextPos := c.Pos.Add(mov.Scaled(dt))
				if win.IsObstacle(nextPos) == false {
					scores[i] += 5
				}
			}

			dir = 0
			for i := 1; i < 4; i++ {
				if scores[i] > scores[dir] {
					dir = i
				}
			}

			c.ChangePose(dirPose[dir])
		}

	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	alien.AddPose("down", []pixel.Vec{{X: 1, Y: startRow}, {X: 2, Y: startRow}, {X: 3 + offsetH, Y: startRow - offsetV}, {X: 4 + offsetH, Y: startRow}, {X: 0, Y: startRow}}, pixel.Vec{X: 0, Y: -200})
	alien.AddPose("side", []pixel.Vec{{X: 1, Y: startRow + 1}, {X: 2, Y: startRow + 1}, {X: 3 + offsetH, Y: startRow + 1 - offsetV}, {X: 4 + offsetH, Y: startRow + 1}, {X: 0, Y: startRow + 1}}, pixel.Vec{X: 200, Y: 0})
	alien.AddPose("up", []pixel.Vec{{X: 1, Y: startRow + 2}, {X: 2, Y: startRow + 2}, {X: 3 + offsetH, Y: startRow + 2 - offsetV}, {X: 4 + offsetH, Y: startRow + 2}, {X: 0, Y: startRow + 2}}, pixel.Vec{X: 0, Y: 200})

	alien.ChangePose("down")
	alien.Pos = pos

	return alien
}
