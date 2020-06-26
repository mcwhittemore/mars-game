package ui

import (
	"app/characters"
	"app/fonts"
	"app/items"
	"app/sheet"

	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func eTagHeroItemUI(mi characters.MindInput) string {
	return "try"
}

func drawHeroItemsUI(mi characters.MindInput) (*pixelgl.Canvas, pixel.Vec) {
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
		count := 0
		if i >= 0 && i < len(hero.Items) {
			ci := hero.Items[i]
			item := items.NewItem(ci.Name, pos, "")
			count = ci.Count
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
		txt := fonts.NewText(fmt.Sprintf("%d", count), pos.Add(pixel.V(3, 3)))
		txt.Draw(canvas, pixel.IM)
	}

	imd.Draw(canvas)

	bds := mi.GetWindowBounds()
	x := bds.W() / 2

	return canvas, pixel.V(x, h/2)
}
