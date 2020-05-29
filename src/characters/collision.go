package characters

import (
	"math"

	"github.com/faiface/pixel"
)

func adjustPosForCollison(target interface{}, nextPos pixel.Vec, startPos pixel.Vec, getBounds func(pixel.Vec) pixel.Rect, findCollision func(pixel.Rect, interface{}) pixel.Rect) (pixel.Vec, bool) {
	charBox := getBounds(nextPos)

	hitrect := findCollision(charBox, target)

	hadHit := false

	ct := 0
	lastHitRect := hitrect
	for hitrect != pixel.ZR {
		ct++
		hadHit = true
		ln := pixel.L(charBox.Center(), hitrect.Center())
		rev := ln.IntersectRect(hitrect)
		rev.X = math.Ceil(rev.X)
		rev.Y = math.Ceil(rev.Y)

		nextPos = nextPos.Sub(rev)

		charBox = getBounds(nextPos)
		lastHitRect = hitrect
		hitrect = findCollision(charBox, target)

		if lastHitRect == hitrect {
			nextPos = startPos
			break
		}
	}

	return nextPos, hadHit
}
