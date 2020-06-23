package items

func CropGrow(item *Item, dt float64, mi MindInput) ItemState {

	age := item.State.Data["age"]
	stage := item.State.Data["stage"]
	controller := "crop-grow"

	age += dt
	if age > 30 {
		age = 0
		item.Icon[0]--
		stage++
	}

	if stage == 3 {
		controller = ""
	}

	data := make(map[string]float64)
	data["age"] = age
	data["stage"] = stage

	return ItemState{
		controller: controller,
		Data:       data,
	}
}
