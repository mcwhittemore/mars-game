package characters

import (
	"app/sheet"

	"math"
	"time"

	"github.com/faiface/pixel"
)

func NewAlien(pos pixel.Vec) *Character {
	startRow := float64(23)

	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{X: 18, Y: 20}, pixel.Vec{X: 0, Y: 2}, 64)
	if err != nil {
		panic(err)
	}

	second := time.Tick(200 * time.Millisecond)

	dirPose := [4]string{"down", "side", "up", "left-side"}

	dir := 0 // 0 = down, 1 = right, 2 = up, 3 = left

	alien := NewCharacter(characterSheet, func(c *Character, dt float64, win MindInput) {
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

	alien.AddPose("down", []pixel.Vec{{1, startRow}, {2, startRow}, {3 + offsetH, startRow - offsetV}, {4 + offsetH, startRow}, {0, startRow}}, pixel.Vec{0, -200})
	alien.AddPose("side", []pixel.Vec{{1, startRow + 1}, {2, startRow + 1}, {3 + offsetH, startRow + 1 - offsetV}, {4 + offsetH, startRow + 1}, {0, startRow + 1}}, pixel.Vec{200, 0})
	alien.AddPose("up", []pixel.Vec{{1, startRow + 2}, {2, startRow + 2}, {3 + offsetH, startRow + 2 - offsetV}, {4 + offsetH, startRow + 2}, {0, startRow + 2}}, pixel.Vec{0, 200})

	alien.ChangePose("down")
	alien.Pos = pos

	return alien
}
