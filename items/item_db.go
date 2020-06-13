package items

type Item struct {
	Id   int
	Type ItemType
	Name string
	Icon [2]int
}

type ItemType int

const (
	Crop ItemType = iota
)

var ItemsDB = make(map[int]Item)

func init() {

	ItemsDB[0] = Item{
		Id:   0,
		Type: Crop,
		Name: "Corn",
	}
}
