package characters

import (
	"app/sheet"

	"testing"

	"github.com/faiface/pixel"
)

func TestPosBounds(t *testing.T) {
	cs, err := sheet.NewSheet("../app/characters.png", pixel.V(32, 32), pixel.ZV, 32)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	char := NewCharacter(cs, pixel.ZV, func(c *Character, dt float64, md MindInput) {})
	char.AddPose("only", []pixel.Vec{{X: 0, Y: 0}}, pixel.ZV)
	char.ChangePose("only")

	received := char.PosBounds(pixel.ZV)

	expected := pixel.R(0, 0, 32, 32)

	if received != expected {
		t.Fatalf("Expected %v, Received %v", expected, received)
	}
}

func TestPosBoundsWithEqualSheetScaling(t *testing.T) {
	cs, err := sheet.NewSheet("../app/characters.png", pixel.V(20, 20), pixel.ZV, 32)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	char := NewCharacter(cs, pixel.ZV, func(c *Character, dt float64, md MindInput) {})
	char.AddPose("only", []pixel.Vec{{X: 0, Y: 0}}, pixel.ZV)
	char.ChangePose("only")

	received := char.PosBounds(pixel.ZV)

	expected := pixel.R(0, 0, 32, 32)

	if received != expected {
		t.Fatalf("Expected %v, Received %v", expected, received)
	}
}

func TestPosBoundsWithSheetScaling(t *testing.T) {
	cs, err := sheet.NewSheet("../app/characters.png", pixel.V(10, 20), pixel.ZV, 32)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	char := NewCharacter(cs, pixel.ZV, func(c *Character, dt float64, md MindInput) {})
	char.AddPose("only", []pixel.Vec{{X: 0, Y: 0}}, pixel.ZV)
	char.ChangePose("only")

	received := char.PosBounds(pixel.ZV)

	// this is 16 because the sprite is scaled uniformly
	// and the Y access grows twice as fast as the X access
	// due to the input spite being a narrow rectangle
	expected := pixel.R(0, 0, 16, 32)

	if received != expected {
		t.Fatalf("Expected %v, Received %v", expected, received)
	}
}
