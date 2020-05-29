package characters

import (
	"sheet"

	"github.com/faiface/pixel"
)

type Character struct {
	Pos      pixel.Vec
	Collided bool
	pose     string
	sheet    *sheet.Sheet
	poses    map[string]*Pose
	mind     MindFunc
}

func NewCharacter(sheet *sheet.Sheet, pos pixel.Vec, mind MindFunc) *Character {
	poses := make(map[string]*Pose)
	return &Character{pos, false, "", sheet, poses, mind}
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

func (c *Character) PosBounds(pos pixel.Vec) pixel.Rect {
	im := c.sheet.IM()
	pose, _ := c.getPose()
	bds := pose.Bounds()
	bds.Min = im.Project(bds.Min)
	bds.Max = im.Project(bds.Max)

	return bds.Moved(pos)
}

func (c *Character) Update(dt float64, win MindInput) (*pixel.Sprite, pixel.Matrix) {
	c.mind(c, dt, win)

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

	nextPos := c.Pos.Add(mov.Scaled(dt))

	nextPos, hadHit := adjustPosForCollison(c, nextPos, c.Pos, c.PosBounds, win.GetCollideRect)
	if hadHit {
		c.Collided = true
	}

	c.Pos = nextPos

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
