package game002

import (
	"app/characters"
	"app/sheet"

	"math/rand"
	"time"

	"github.com/faiface/pixel"
)

func NewRando(startRow float64, startPos pixel.Vec) *characters.Character {
	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{18, 20}, pixel.Vec{0, 2}, 64)
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

	rando.AddPose("down", []pixel.Vec{{1, startRow}, {2, startRow}, {3 + offsetH, startRow - offsetV}, {4 + offsetH, startRow}, {0, startRow}}, pixel.Vec{0, -200})
	rando.AddPose("side", []pixel.Vec{{1, startRow + 1}, {2, startRow + 1}, {3 + offsetH, startRow + 1 - offsetV}, {4 + offsetH, startRow + 1}, {0, startRow + 1}}, pixel.Vec{200, 0})
	rando.AddPose("up", []pixel.Vec{{1, startRow + 2}, {2, startRow + 2}, {3 + offsetH, startRow + 2 - offsetV}, {4 + offsetH, startRow + 2}, {0, startRow + 2}}, pixel.Vec{0, 200})

	rando.ChangePose("down")

	return rando
}
