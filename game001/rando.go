package game001

import (
	"app/characters"
	"app/sheet"

	"math/rand"
	"time"

	"github.com/faiface/pixel"
)

func NewRando(startRow float64, startPos pixel.Vec) *characters.Character {
	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{X: 18, Y: 20}, pixel.Vec{X: 0, Y: 2}, 64)
	if err != nil {
		panic(err)
	}

	second := time.Tick(200 * time.Millisecond)

	moves := []string{"down", "up", "side", "left-side"}

	safe := pixel.R(188, 200, 388, 400)
	rando := characters.NewCharacter(characterSheet, startPos, func(c *characters.Character, dt float64, win characters.MindInput) {
		select {
		case <-second:
			c.ChangePose(moves[rand.Int()%4])
			if rand.Int()%4 == 0 {
				c.Stop()
			} else {
				c.Step()
			}
		default:
		}

		nextPos := c.GetNextPos(dt)

		isSafe := c.Hits(safe)
		if isSafe {
			c.Pos = nextPos
			return
		}

		selfbox := c.PosBounds(nextPos)
		_, subject := win.GetCollideRect(selfbox, interface{}(c))
		if subject == nil {
			c.Pos = nextPos
		}
	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	rando.AddPose("down", []pixel.Vec{{X: 1, Y: startRow}, {X: 2, Y: startRow}, {X: 3 + offsetH, Y: startRow - offsetV}, {X: 4 + offsetH, Y: startRow}, {X: 0, Y: startRow}}, pixel.Vec{X: 0, Y: -200})
	rando.AddPose("side", []pixel.Vec{{X: 1, Y: startRow + 1}, {X: 2, Y: startRow + 1}, {X: 3 + offsetH, Y: startRow + 1 - offsetV}, {X: 4 + offsetH, Y: startRow + 1}, {X: 0, Y: startRow + 1}}, pixel.Vec{X: 200, Y: 0})
	rando.AddPose("up", []pixel.Vec{{X: 1, Y: startRow + 2}, {X: 2, Y: startRow + 2}, {X: 3 + offsetH, Y: startRow + 2 - offsetV}, {X: 4 + offsetH, Y: startRow + 2}, {X: 0, Y: startRow + 2}}, pixel.Vec{X: 0, Y: 200})

	rando.ChangePose("down")

	return rando
}
