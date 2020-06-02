package characters

import (
	"app/sheet"

	"github.com/faiface/pixel"
)

type Character struct {
	Pos   pixel.Vec
	pose  string
	sheet *sheet.Sheet
	poses map[string]*Pose
	mind  MindFunc
}

func NewCharacter(sheet *sheet.Sheet, pos pixel.Vec, mind MindFunc) *Character {
	poses := make(map[string]*Pose)
	return &Character{pos, "", sheet, poses, mind}
}

func (c *Character) ChangePose(name string) {
	if c.pose != "" && c.pose != name {
		pose, _ := c.GetPose()
		pose.Stop()
	}
	c.pose = name
}

func (c *Character) AddPose(name string, f []pixel.Vec, mv pixel.Vec) {

	frames := c.sheet.GetSprites(f)

	c.poses[name] = NewPose(frames, mv)
}

func (c *Character) GetPose() (*Pose, bool) {
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
	pose, _ := c.GetPose()
	bds := pose.Bounds()
	bds.Min = im.Project(bds.Min)
	bds.Max = im.Project(bds.Max)
	w := (bds.Max.X - bds.Min.X) / 2
	h := (bds.Max.Y - bds.Min.Y) / 2

	return pixel.R(pos.X-w, pos.Y-h, pos.X+w, pos.Y+h)
}

func (c *Character) Hits(rect pixel.Rect) bool {
	bds := c.PosBounds(c.Pos)
	return bds.Intersect(rect) != pixel.ZR
}

func (c *Character) DropNear(pos pixel.Vec, findCollision FindCollision) bool {
	nextPos, hadHit := adjustPosForCollison(c, pos, c.Pos, c.PosBounds, findCollision)
	c.Pos = nextPos
	return hadHit

}

func (c *Character) GetNextPos(dt float64) pixel.Vec {
	pose, isLeft := c.GetPose()
	mov := pose.GetMovement()
	flip := pixel.V(-1, 1)
	if isLeft {
		mov = mov.ScaledXY(flip)
	}

	return c.Pos.Add(mov.Scaled(dt))
}

func (c *Character) Update(dt float64, win MindInput) (*pixel.Sprite, pixel.Matrix) {
	c.mind(c, dt, win)

	pose, isLeft := c.GetPose()
	sprite := pose.GetSprite()

	matrix := c.sheet.IM()
	flip := pixel.V(-1, 1)
	faceLeft := pixel.IM.ScaledXY(pixel.ZV, flip)

	if isLeft {
		matrix = matrix.Chained(faceLeft)
	}

	return sprite, matrix.Moved(c.Pos)
}

func (c *Character) Step() {
	pose, _ := c.GetPose()
	pose.Step()
}

func (c *Character) Stop() {
	pose, _ := c.GetPose()
	pose.Stop()
}
