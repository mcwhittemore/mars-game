package world

import (
	"characters"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type World struct {
	maps   []*pixel.Batch
	npcs   []*characters.Character
	hero   *characters.Character
	camPos pixel.Vec
	win    *pixelgl.Window
}

func (w *World) JustPressed(but pixelgl.Button) bool {
	return w.win.JustPressed(but)
}

func (w *World) Pressed(but pixelgl.Button) bool {
	return w.win.Pressed(but)
}

func (w *World) GetCollideRect(rect pixel.Rect, thing interface{}) pixel.Rect {
	out := pixel.ZR

	if thing != interface{}(w.hero) {
		out = w.hero.PosBounds(w.hero.Pos).Intersect(rect)
		if out != pixel.ZR {
			return out
		}
	}

	for _, being := range w.npcs {
		if thing != interface{}(being) {
			out = being.PosBounds(being.Pos).Intersect(rect)
			if out != pixel.ZR {
				return out
			}
		}
	}

	return out
}

func (w *World) DrawHitBoxes() {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

	hb := w.hero.PosBounds(w.hero.Pos)
	imd.Push(hb.Min, hb.Max)
	imd.Rectangle(2)

	for _, being := range w.npcs {
		bb := being.PosBounds(being.Pos)
		imd.Push(bb.Min, bb.Max)
		imd.Rectangle(2)
	}

	imd.Draw(w.win)
}

func (w *World) Update(dt float64) {
	cam := pixel.IM.Moved(w.win.Bounds().Center().Sub(w.camPos))
	w.win.SetMatrix(cam)

	activeMap := w.maps[0]
	activeMap.Draw(w.win)

	hs, hm := w.hero.Update(dt, w)
	hs.Draw(w.win, hm)

	for _, being := range w.npcs {
		bs, bm := being.Update(dt, w)
		bs.Draw(w.win, bm)
	}
}
