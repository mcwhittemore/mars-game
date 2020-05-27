package characters

import (
	"sheet"

	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewRando(win *pixelgl.Window, startRow float64, startPos pixel.Vec) *Character {
	characterSheet, err := sheet.NewSheet("characters.png", pixel.Vec{18, 20}, pixel.Vec{0, 2}, 64)
	if err != nil {
		panic(err)
	}

	second := time.Tick(200 * time.Millisecond)

	moves := []string{"down", "up", "side", "left-side"}

	rando := NewCharacter(characterSheet, startPos, func(c *Character, dt int64) {
		select {
		case <-second:
			c.ChangePose(moves[rand.Int()%4])
			if rand.Int()%2 == 0 {
				c.Step()
			} else {
				c.Stop()
			}
		default:
		}
	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	rando.AddPose("down", []pixel.Vec{{1, startRow}, {2, startRow}, {3 + offsetH, startRow - offsetV}, {4 + offsetH, startRow}, {0, startRow}}, pixel.Vec{0, -1})
	rando.AddPose("side", []pixel.Vec{{1, startRow + 1}, {2, startRow + 1}, {3 + offsetH, startRow + 1 - offsetV}, {4 + offsetH, startRow + 1}, {0, startRow + 1}}, pixel.Vec{1, 0})
	rando.AddPose("up", []pixel.Vec{{1, startRow + 2}, {2, startRow + 2}, {3 + offsetH, startRow + 2 - offsetV}, {4 + offsetH, startRow + 2}, {0, startRow + 2}}, pixel.Vec{0, 1})

	rando.ChangePose("down")

	return rando
}
