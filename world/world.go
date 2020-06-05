package world

import (
	"app/characters"
	"app/maps"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type World struct {
	Maps   []*maps.Map
	NPCs   []*characters.Character
	Hero   *characters.Character
	CamPos pixel.Vec
	Win    *pixelgl.Window
}

func (w *World) GetHeroPos() pixel.Vec {
	return w.Hero.Pos
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

func (w *World) IsObstacle(pos pixel.Vec) bool {
	activeMap := w.Maps[0]
	return activeMap.IsObstacle(pos)
}

func (w *World) KeepInView(pos pixel.Vec, mov pixel.Vec, buffer float64) {
	bds := w.Win.Bounds()
	cam := pixel.IM.Moved(w.Win.Bounds().Center().Sub(w.CamPos))

	viewBox := pixel.Rect{
		Min: cam.Unproject(bds.Min),
		Max: cam.Unproject(bds.Max),
	}

	viewbox := viewBox.Edges()
	for _, edge := range viewbox {
		closest := edge.Closest(pos)
		dis := pixel.L(pos, closest).Len()
		if dis < buffer {
			w.CamPos = w.CamPos.Add(mov)
			break
		}
	}
}

func (w *World) Update(dt float64) {
	cam := pixel.IM.Moved(w.Win.Bounds().Center().Sub(w.CamPos))
	w.Win.SetMatrix(cam)

	activeMap := w.Maps[0]
	activeMap.Render.Draw(w.Win)

	hs, hm := w.Hero.Update(dt, w)
	hs.Draw(w.Win, hm)

	for _, being := range w.NPCs {
		bs, bm := being.Update(dt, w)
		bs.Draw(w.Win, bm)
	}
}
