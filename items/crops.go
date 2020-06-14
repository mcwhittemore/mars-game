package items

func NewMindCropGrow() MindFunc {

	age := float64(0)

	stage := 0

	return func(item *Item, dt float64, mi MindInput) {
		age += dt
		if age > 30 {
			age = 0
			item.Icon[0]--
			stage++
		}

		if stage == 3 {
			item.Mind = nil
		}
	}
}
