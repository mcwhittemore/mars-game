package game

type Line []Pos

func (l Line) Contains(p Pos) bool {
	for _, a := range l {
		if a.Eq(p) {
			return true
		}
	}

	return false
}
