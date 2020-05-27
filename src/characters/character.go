package characters

import (
	"sheet"

	"github.com/faiface/pixel"
)

type Character struct {
	Pos   pixel.Vec
	pose  string
	sheet *sheet.Sheet
	poses map[string]*Pose
	mind  func(*Character, int64)
}

func NewCharacter(sheet *sheet.Sheet, pos pixel.Vec, mind func(*Character, int64)) *Character {
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
