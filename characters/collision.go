package characters

import (
	"github.com/faiface/pixel"
)

type FindCollision func(pixel.Rect, interface{}) (pixel.Rect, *Character)

func adjustPosForCollison(target interface{}, nextPos pixel.Vec, startPos pixel.Vec, getBounds func(pixel.Vec) pixel.Rect, findCollision FindCollision) (pixel.Vec, bool) {

	diff := nextPos.Sub(startPos)
	mvAngle := diff.Angle()
	step := pixel.Unit(mvAngle)

	charBox := getBounds(nextPos)

	hitrect, _ := findCollision(charBox, target)

	hadHit := false

	lastHitRect := hitrect
	for hitrect != pixel.ZR {
		hadHit = true

		nextPos = nextPos.Add(step)

		charBox = getBounds(nextPos)
		lastHitRect = hitrect
		hitrect, _ = findCollision(charBox, target)

		if lastHitRect == hitrect {
			nextPos = startPos
			break
		}
	}

	return nextPos, hadHit
}
