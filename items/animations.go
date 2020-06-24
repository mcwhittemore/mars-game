package items

func XAnimation(item *Item, dt float64, mi MindInput) ItemState {
	time := mi.GetTime() * 2
	width := GetSheet(item.Sheet).GetWidth()

	sec := int(time)
	where := time - float64(sec)

	step := float64(1) / float64(width)

	frame := int(where / step)

	item.Icon[0] = float64(frame)

	return item.State
}
