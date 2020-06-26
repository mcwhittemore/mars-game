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
	numItems := float64(5)
	padding := float64(5)
	w := float64((sheet.TileSize * numItems) + (padding * (numItems + 1)))
	h := float64(sheet.TileSize + (padding * 2))

	canvas := pixelgl.NewCanvas(pixel.R(0, 0, w, h))
	canvas.Clear(color.Black)

	hero := mi.GetCharacter("hero")

	imd := imdraw.New(nil)

	start := hero.InHands - int(numItems/2)
	for loc := float64(0); loc < numItems; loc++ {
		i := int(loc) + start
		spot := (loc * sheet.TileSize) + (padding * loc)

		pos := pixel.V(spot+padding, padding)
		if i >= 0 && i < len(hero.Items) {
			ci := hero.Items[i]
			item := items.NewItem(ci.Name, pos, "")
			sprite, im := item.GetSprite()
			sprite.Draw(canvas, im.Scaled(pixel.ZV, 2).Moved(pos.Add(pixel.V(24, 24))))
		}

		if i == hero.InHands {
			imd.Color = pixel.RGB(1, 0, 0)
		} else {

			imd.Color = pixel.RGB(1, 1, 1)
		}
		imd.Push(pos, pos.Add(pixel.V(sheet.TileSize, sheet.TileSize)))
		imd.Rectangle(2)
	}

	imd.Draw(canvas)

	bds := mi.GetWindowBounds()
	x := bds.W() / 2
	mi.AddCanvasStatic(canvas, pixel.V(x, h/2))
}
