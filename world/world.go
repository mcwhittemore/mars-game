package world

import (
	"app/characters"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type World struct {
	Maps   []*pixel.Batch
	NPCs   []*characters.Character
	Hero   *characters.Character
	CamPos pixel.Vec
	Win    *pixelgl.Window
}

func (w *World) JustPressed(but pixelgl.Button) bool {
	return w.Win.JustPressed(but)
}

func (w *World) Pressed(but pixelgl.Button) bool {
	return w.Win.Pressed(but)
}

func (w *World) GetCollideRect(rect pixel.Rect, thing interface{}) (pixel.Rect, *characters.Character) {
	var out pixel.Rect

	if thing != interface{}(w.Hero) {
		out = w.Hero.PosBounds(w.Hero.Pos).Intersect(rect)
		if out != pixel.ZR {
			return out, w.Hero
		}
	}

	for _, being := range w.NPCs {
		if thing != interface{}(being) {
			out = being.PosBounds(being.Pos).Intersect(rect)
			if out != pixel.ZR {
				return out, being
			}
		}
	}

	return pixel.ZR, nil
}

func (w *World) DrawHitBoxes() {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

	hb := w.Hero.PosBounds(w.Hero.Pos)
	imd.Push(hb.Min, hb.Max)
	imd.Rectangle(2)

	for _, being := range w.NPCs {
		bb := being.PosBounds(being.Pos)
		imd.Push(bb.Min, bb.Max)
		imd.Rectangle(2)
	}

	imd.Draw(w.Win)
}

func (w *World) KeepInView(pos pixel.Vec, mov pixel.Vec, buffer float64) {
	cam := w.GetCamPos()
	viewbox := w.GetViewBox().Edges()
	for _, edge := range viewbox {
		closest := edge.Closest(pos)
		dis := pixel.L(pos, closest).Len()
		if dis < buffer {
			w.SetCamPos(cam.Add(mov))
			break
		}
	}
}

func (w *World) GetViewBox() *pixel.Rect {
	bds := w.Win.Bounds()
	cam := pixel.IM.Moved(w.Win.Bounds().Center().Sub(w.CamPos))

	return &pixel.Rect{
		Min: cam.Unproject(bds.Min),
		Max: cam.Unproject(bds.Max),
	}

}

func (w *World) GetCamPos() pixel.Vec {
	return w.CamPos
}

func (w *World) SetCamPos(pos pixel.Vec) {
	w.CamPos = pos
}

func (w *World) Update(dt float64) {
	cam := pixel.IM.Moved(w.Win.Bounds().Center().Sub(w.CamPos))
	w.Win.SetMatrix(cam)

	activeMap := w.Maps[0]
	activeMap.Draw(w.Win)

	hs, hm := w.Hero.Update(dt, w)
	hs.Draw(w.Win, hm)

	for _, being := range w.NPCs {
		bs, bm := being.Update(dt, w)
		bs.Draw(w.Win, bm)
	}
}
