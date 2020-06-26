package ui

import (
	"app/characters"
	"app/items"
	"app/sheet"

	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func DrawHeroItemsUI(mi characters.MindInput) {
	canvas := pixelgl.NewCanvas(pixel.R(0, 0, 440, 42))
	canvas.Clear(color.Black)

	hero := mi.GetCharacter("hero")

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1, 0, 0)

	for i := hero.InHands - 2; i <= hero.InHands+2; i++ {
		pos := pixel.V(float64(((i+2)*64)+10), 10)
		if i >= 0 && i < len(hero.Items) {
			ci := hero.Items[i]
			item := items.NewItem(ci.Name, pos, "")
			sprite, im := item.GetSprite()
			sprite.Draw(canvas, im.Scaled(pixel.ZV, 2).Moved(pos.Add(pixel.V(24, 24))))
		}

		imd.Push(pos, pos.Add(pixel.V(sheet.TileSize, sheet.TileSize)))
		imd.Rectangle(2)

	}

	imd.Draw(canvas)

	bds := mi.GetWindowBounds()
	x := bds.W() / 2
	mi.AddCanvasStatic(canvas, pixel.V(x, 21))
}
