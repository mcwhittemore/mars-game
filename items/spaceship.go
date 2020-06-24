package items

import (
	"app/sheet"

	"strings"

	"github.com/faiface/pixel"
)

type SpaceshipMode float64

const (
	Spaceship_Unloading SpaceshipMode = iota
	Spaceship_Loading
	Spaceship_Launching
	Spaceship_Waiting
	Spaceship_Landing
)

func ControlSpaceship(item *Item, dt float64, mi MindInput) ItemState {

	lp := mi.GetLocation("landing-pad").Center()

	mode := SpaceshipMode(item.State.Data["mode"])
	dur := item.State.Data["dur"]

	if mode == Spaceship_Unloading {
		item.Icon[1] = 0
		dzVerts := mi.GetLocation("from-earth-cb-1").Vertices()
		dz := dzVerts[1].Add(pixel.V(sheet.TileSize/2, 0))
		dur += dt
		if dur > 1 {
			dur = 0
			nextItem, key := spaceshipNextItem(item.State.Data)

			if nextItem == "" {
				mode = Spaceship_Loading
				dur = 0
			} else {
				item.State.Data[key]--
				if item.State.Data[key] == 0 {
					delete(item.State.Data, key)
				}
				t := NewItem(nextItem, dz, "")
				mi.AddItem(t)
			}
		}
	}

	if mode == Spaceship_Loading {
		dur += dt
		if dur > 5 {
			mode = Spaceship_Launching
			dur = 0
		}
	}

	if mode == Spaceship_Launching {
		y := (item.Pos.Y - lp.Y) + 5

		item.Pos = item.Pos.Add(pixel.V(0, y*2*dt))
		item.Icon[1] = 1

		if y > 40000 {
			mode = Spaceship_Waiting
			dur = 0
		}
	}

	if mode == Spaceship_Waiting {
		dur += dt
		if dur > 30 {
			mode = Spaceship_Landing
			dur = 0
		}
	}

	if mode == Spaceship_Landing {
		y := item.Pos.Y - lp.Y

		item.Pos = item.Pos.Sub(pixel.V(0, (y/2)*dt))
		if y <= 5 {
			item.Pos = lp
			dur = 0
			mode = Spaceship_Unloading
		}
	}

	item.State.Data["mode"] = float64(mode)
	item.State.Data["dur"] = dur

	return item.State
}

func spaceshipNextItem(data map[string]float64) (string, string) {
	for key, _ := range data {
		if strings.HasPrefix(key, "item_") {
			return strings.TrimPrefix(key, "item_"), key
		}
	}
	return "", ""
}
