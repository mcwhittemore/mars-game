package items

type ItemState struct {
	controller string
	Data       map[string]float64
}

func (is *ItemState) ChangeController(name string) {
	is.controller = name
	for key, _ := range is.Data {
		delete(is.Data, key)
	}
}

func (is *ItemState) UsingController(name string) bool {
	return name == is.controller
}

func (is *ItemState) Update(item *Item, dt float64, mi MindInput) {
	if is.UsingController("crop-grow") {
		ns := CropGrow(item, dt, mi)
		is.controller = ns.controller
		is.Data = ns.Data
	}
}
