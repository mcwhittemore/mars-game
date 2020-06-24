package items

type ItemState struct {
	Controller string
	Data       map[string]float64
}

func (is *ItemState) ChangeController(name string) {
	is.Controller = name
	for key, _ := range is.Data {
		delete(is.Data, key)
	}
}

func (is *ItemState) UsingController(name string) bool {
	return name == is.Controller
}

func (is *ItemState) Update(item *Item, dt float64, mi MindInput) {
	if is.UsingController("crop-grow") {
		ns := CropGrow(item, dt, mi)
		is.Controller = ns.Controller
		is.Data = ns.Data
	} else if is.UsingController("conveyor-belts") {
		ns := ControlConveyorBelt(item, dt, mi)
		is.Controller = ns.Controller
		is.Data = ns.Data
	}
}
