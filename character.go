package main

import (
	"fmt"
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
	fmt.Printf("Step %d ->", p.frame)
	if p.numFrames > 0 {
		p.frame = p.frame + 1

		if p.frame == p.numFrames {
			p.frame = 1
		}
	}
	fmt.Printf("%d\n", p.frame)
}

func (p *Pose) Stop() {
	p.frame = 0
}

type Character struct {
	Pos   pixel.Vec
	pose  string
	sheet *Sheet
	poses map[string]*Pose
	mind  func(*Character, int64)
}

func NewCharacter(sheet *Sheet, pos pixel.Vec, mind func(*Character, int64)) *Character {
	poses := make(map[string]*Pose)
	return &Character{pos, "", sheet, poses, mind}
}

func (c *Character) ChangePose(name string) {
	if c.pose != "" && c.pose != name {
		pose, _ := c.getPose()
		pose.Stop()
	}
	c.pose = name
}

func (c *Character) AddPose(name string, f []pixel.Vec, mv pixel.Vec) {

	frames := c.sheet.GetSprites(f)

	c.poses[name] = NewPose(frames, mv)
}

func (c *Character) getPose() (*Pose, bool) {
	if len(c.pose) < 5 {
		pose := c.poses[c.pose]
		return pose, false
	}
	pre := c.pose[0:5]
	post := c.pose[5:]
	if pre == "left-" {
		pose := c.poses[post]
		return pose, true
	}
	pose := c.poses[c.pose]
	return pose, false
}

func (c *Character) Update(dt int64) (*pixel.Sprite, pixel.Matrix) {
	c.mind(c, dt)

	pose, isLeft := c.getPose()
	sprite := pose.GetSprite()

	mov := pose.GetMovement()

	matrix := c.sheet.IM()
	flip := pixel.V(-1, 1)
	faceLeft := pixel.IM.ScaledXY(pixel.ZV, flip)

	if isLeft {
		mov = mov.ScaledXY(flip)
		matrix = matrix.Chained(faceLeft)
	}

	c.Pos = c.Pos.Add(mov.Scaled(float64(dt)))

	return sprite, matrix.Moved(c.Pos)
}

func (c *Character) Step() {
	pose, _ := c.getPose()
	pose.Step()
}

func (c *Character) Stop() {
	pose, _ := c.getPose()
	pose.Stop()
}
