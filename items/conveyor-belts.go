package items

import (
	"app/sheet"

	"github.com/faiface/pixel"
)

func ControlConveyorBelt(item *Item, dt float64, mi MindInput) ItemState {
	time := mi.GetTime() * 2

	width := GetSheet(item.Sheet).GetWidth()

	sec := int(time)
	where := time - float64(sec)

	step := float64(1) / float64(width)

	frame := int(where / step)

	item.Icon[0] = float64(frame)

	right := pixel.V(1, 0)
	left := pixel.V(-1, 0)
	up := pixel.V(0, 1)
	down := pixel.V(0, -1)

	speed := float64(sheet.TileSize)

	mov := pixel.ZV

	var grab [2]int
	var buf [2]pixel.Vec

	switch item.Icon[1] {
	case 0:
		mov = right.Scaled(speed * dt)
		grab[0] = 3
		grab[1] = 2
		buf[0] = left
	case 1:
		mov = down.Scaled(speed * dt)
		grab[0] = 1
		grab[1] = 2
		buf[1] = up
	case 2:
		mov = left.Scaled(speed * dt)
		grab[0] = 0
		grab[1] = 1
		buf[1] = right
	case 3:
		mov = up.Scaled(speed * dt)
		grab[0] = 0
		grab[1] = 3
		buf[0] = down
	}

	itemList := mi.GetItems(item.PosBounds(item.Pos), func(t *Item) pixel.Rect {
		if t == item {
			return pixel.ZR
		}

		bds := t.PosBounds(t.Pos)
		return bds
	})

	for _, t := range itemList {
		t.Pos = t.Pos.Add(mov)
	}

	return item.State
}
