package characters

import (
	"github.com/faiface/pixel"
)

type Pose struct {
	frames    []*pixel.Sprite
	numFrames int
	frame     int
	movement  pixel.Vec
}

func NewPose(fs []*pixel.Sprite, mv pixel.Vec) *Pose {
	return &Pose{fs, len(fs), 0, mv}
}

func (p *Pose) Bounds() pixel.Rect {
	return p.GetSprite().Frame()
}

func (p *Pose) GetMovement() pixel.Vec {
	if p.frame == 0 {
		return pixel.ZV
	}

	return p.movement
}

func (p *Pose) GetSprite() *pixel.Sprite {
	return p.frames[p.frame]
}

func (p *Pose) Step() {
	if p.numFrames > 0 {
		p.frame = p.frame + 1

		if p.frame == p.numFrames {
			p.frame = 1
		}
	}
}

func (p *Pose) Stop() {
	p.frame = 0
}
