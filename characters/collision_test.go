package characters

import (
	"testing"

	"github.com/faiface/pixel"
)

func TestAdjustPosOnCollide(t *testing.T) {

	target := "target"
	nextPos := pixel.V(10, 10)
	startPos := pixel.V(8, 8)

	getBounds := func(min pixel.Vec) pixel.Rect {
		max := min.Add(pixel.V(32, 32))
		return pixel.Rect{min, max}
	}

	findCollision := func(bds pixel.Rect, target interface{}) pixel.Rect {
		couldHit := pixel.R(41, 41, 50, 50)
		return couldHit.Intersect(bds)
	}

	nextPos, hadHit := adjustPosForCollison(target, nextPos, startPos, getBounds, findCollision)

	if hadHit == false {
		t.Fatalf("Expected a hit but did not get one")
	}

	if nextPos == pixel.V(8, 8) {
		t.Fatalf("Returned position provided due to an infinite loop")
	}

	expectedPos := pixel.V(9, 10)

	if expectedPos != nextPos {
		t.Fatalf("Received %v but expected %v", nextPos, expectedPos)
	}
}
