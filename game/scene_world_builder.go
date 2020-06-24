package game

import (
	"app/characters"
	"app/fonts"
	"app/items"
	"app/maps"
	"app/sheet"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type WorldBuilder struct {
	mode          string
	cam           pixel.Vec
	MapOpts       *maps.MapOpts
	Ground        *maps.Map
	TileSheet     *sheet.TileSheet
	Pos           pixel.Vec
	TileId        int
	path          string
	structurePath string
	locationName  string
	buildingNames []string
	buildingIdx   int
}

func (g *WorldBuilder) Update(dt float64, mi characters.MindInput) {

	if g.mode == "input" {
		g.inputMode(mi)
	} else if g.mode == "normal" {
		g.normalMode(mi)
	} else if g.mode == "location" {
		g.locationMode(mi)
	} else if g.mode == "new-location" {
		g.newLocationMode(mi)
	} else if g.mode == "remove" {
		g.removeMode(mi)
	} else if g.mode == "building" {
		g.buildingMode(mi)
	}

	mi.KeepInView(g.Pos.Scaled(sheet.TileSize), 128)

	imd := imdraw.New(nil)

	if g.mode == "input" {
		imd.Color = pixel.RGB(1, 0, 0)
	} else if g.mode == "normal" {
		imd.Color = pixel.RGB(0, 0, 1)
	} else if g.mode == "location" || g.mode == "new-location" {
		imd.Color = pixel.RGB(0, 1, 0)
	}

	pos := g.Pos.Scaled(sheet.TileSize).Sub(pixel.V(sheet.TileSize/2, sheet.TileSize/2))
	abt := fmt.Sprintf("x: %d/%d\ny: %d/%d", int(g.Pos.X), len(g.MapOpts.Grid[0]), int(g.Pos.Y), len(g.MapOpts.Grid))
	mi.AddText(fonts.NewText(abt, pos.Add(pixel.V(0, sheet.TileSize*.8))))
	imd.Push(pos, pos.Add(pixel.V(sheet.TileSize, sheet.TileSize)))
	imd.Rectangle(2)

	imd.Color = pixel.RGB(1, 1, 0)
	for name, loc := range g.MapOpts.Locations {
		lMin := loc.Min.Scaled(sheet.TileSize).Sub(pixel.V(sheet.TileSize/2, sheet.TileSize/2))
		lMax := loc.Max.Scaled(sheet.TileSize).Sub(pixel.V(sheet.TileSize/2, sheet.TileSize/2))

		txt := fonts.NewText(name, lMin.Add(pixel.V(2, 2)))
		mi.AddText(txt)
		imd.Push(lMin, lMax)
		imd.Rectangle(2)
	}

	mi.AddDraw(imd)
}

func (g *WorldBuilder) removeBuilding(mi characters.MindInput) {
	pos := g.Pos.Scaled(sheet.TileSize)
	items := mi.GetItems(pixel.Rect{
		Min: pos,
		Max: pos.Add(pixel.V(1, 1)),
	}, nil)
	if len(items) != 0 {
		mi.RemoveItem(items[0])
	}
}

func (g *WorldBuilder) addBuilding(mi characters.MindInput) {
	g.removeBuilding(mi)
	name := g.buildingNames[g.buildingIdx]
	pos := g.Pos.Scaled(sheet.TileSize)
	item := items.NewItem(name, pos, "")
	mi.AddItem(item)
}

func (g *WorldBuilder) buildingMode(mi characters.MindInput) {
	moveUpdate := false
	pos := g.Pos
	if mi.JustPressed(pixelgl.KeyEscape) {
		g.mode = "normal"
		g.save(mi)
	} else if mi.JustPressed(pixelgl.KeyA) {
		pos = g.Pos.Add(pixel.V(-1, 0))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyW) {
		pos = g.Pos.Add(pixel.V(0, 1))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyD) {
		pos = g.Pos.Add(pixel.V(1, 0))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyS) {
		pos = g.Pos.Add(pixel.V(0, -1))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyEnter) {
		g.addBuilding(mi)
	} else if mi.JustPressed(pixelgl.KeyJ) {
		g.buildingIdx++
		if g.buildingIdx == len(g.buildingNames) {
			g.buildingIdx = 0
		}
		g.addBuilding(mi)
	} else if mi.JustPressed(pixelgl.KeyK) {
		g.buildingIdx--
		if g.buildingIdx == -1 {
			g.buildingIdx = len(g.buildingNames) - 1
		}
		g.addBuilding(mi)
	} else if mi.JustPressed(pixelgl.KeyDelete) || mi.JustPressed(pixelgl.KeyBackspace) {
		g.removeBuilding(mi)
	}

	if moveUpdate {
		maxX := len(g.MapOpts.Grid[0])
		maxY := len(g.MapOpts.Grid)

		if pos.X >= 0 && int(pos.X) < maxX && pos.Y >= 0 && int(pos.Y) < maxY {
			g.Pos = pos
		}
	}
}

func (g *WorldBuilder) newLocationMode(mi characters.MindInput) {
	if mi.JustPressed(pixelgl.KeyEscape) {
		g.mode = "normal"
		g.locationName = ""
		return
	}
	g.locationName += mi.Typed()

	if mi.JustPressed(pixelgl.KeyDelete) || mi.JustPressed(pixelgl.KeyBackspace) {
		if len(g.locationName) > 0 {
			g.locationName = g.locationName[:len(g.locationName)-1]
		}
	}

	if mi.JustPressed(pixelgl.KeyEnter) {
		for i, v := range g.locationName {
			if string(v) == "\n" {
				g.locationName = g.locationName[0 : i-1]
				break
			}
		}

		g.MapOpts.Locations[g.locationName] = pixel.Rect{
			Min: g.Pos,
			Max: g.Pos.Add(pixel.V(1, 1)),
		}
		g.mode = "location"
	} else {
		pos := g.Pos.Scaled(sheet.TileSize).Sub(pixel.V(sheet.TileSize/2, sheet.TileSize/2))
		mi.AddText(fonts.NewText(g.locationName, pos.Add(pixel.V(2, 2))))
	}
}

func (g *WorldBuilder) locationMode(mi characters.MindInput) {
	pos := g.Pos
	moveUpdate := false
	if mi.JustPressed(pixelgl.KeyA) {
		pos = g.Pos.Add(pixel.V(-1, 0))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyW) {
		pos = g.Pos.Add(pixel.V(0, 1))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyD) {
		pos = g.Pos.Add(pixel.V(1, 0))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyS) {
		pos = g.Pos.Add(pixel.V(0, -1))
		moveUpdate = true
	} else if mi.JustPressed(pixelgl.KeyEscape) {
		delete(g.MapOpts.Locations, g.locationName)
		g.locationName = ""
		g.mode = "normal"
		return
	} else if mi.JustPressed(pixelgl.KeyEnter) {
		g.mode = "normal"
		g.Ground = maps.NewMap(g.MapOpts)
		g.locationName = ""
		g.save(mi)
		return
	}

	if moveUpdate {
		maxX := len(g.MapOpts.Grid[0])
		maxY := len(g.MapOpts.Grid)

		if pos.X >= 0 && int(pos.X) < maxX && pos.Y >= 0 && int(pos.Y) < maxY {
			g.Pos = pos
		}

		min := g.Pos
		max := min.Add(pixel.V(1, 1))

		box := g.MapOpts.Locations[g.locationName]

		if box.Min.X > min.X {
			box.Min.X = min.X
		}
		if box.Max.X < max.X {
			box.Max.X = max.X
		}

		if box.Min.Y > min.Y {
			box.Min.Y = min.Y
		}
		if box.Max.Y < max.Y {
			box.Max.Y = max.Y
		}

		g.MapOpts.Locations[g.locationName] = box
	}

}

func (g *WorldBuilder) removeMode(mi characters.MindInput) {
	if mi.JustPressed(pixelgl.KeyL) {
		pos := g.Pos.Scaled(sheet.TileSize)
		loc, _ := g.Ground.GetLocationAt(pos)
		delete(g.MapOpts.Locations, loc)
		g.Ground = maps.NewMap(g.MapOpts)
		g.save(mi)
		g.mode = "normal"
	} else if mi.JustPressed(pixelgl.KeyEscape) {
		g.mode = "normal"
	}
}

func (g *WorldBuilder) normalMode(mi characters.MindInput) {
	pos := g.Pos
	if mi.JustPressed(pixelgl.KeyA) {
		pos = g.Pos.Add(pixel.V(-1, 0))
	} else if mi.JustPressed(pixelgl.KeyW) {
		pos = g.Pos.Add(pixel.V(0, 1))
	} else if mi.JustPressed(pixelgl.KeyD) {
		pos = g.Pos.Add(pixel.V(1, 0))
	} else if mi.JustPressed(pixelgl.KeyS) {
		pos = g.Pos.Add(pixel.V(0, -1))
	} else if mi.JustPressed(pixelgl.KeyI) {
		g.mode = "input"
	} else if mi.JustPressed(pixelgl.KeyL) {
		g.mode = "new-location"
	} else if mi.JustPressed(pixelgl.KeyR) {
		g.mode = "remove"
	} else if mi.JustPressed(pixelgl.KeyB) {
		g.mode = "building"
	}

	maxX := len(g.MapOpts.Grid[0])
	maxY := len(g.MapOpts.Grid)

	if pos.X >= 0 && int(pos.X) < maxX && pos.Y >= 0 && int(pos.Y) < maxY {
		g.Pos = pos
	}
}

func (g *WorldBuilder) inputMode(mi characters.MindInput) {
	needsNewMap := false

	if mi.JustPressed(pixelgl.KeyEscape) {
		g.mode = "normal"
		g.save(mi)
	} else if mi.JustPressed(pixelgl.KeyJ) {
		needsNewMap = true
		g.TileId++
		if g.TileId == len(g.TileSheet.TileTypes) {
			g.TileId = 0
		}
	} else if mi.JustPressed(pixelgl.KeyK) {
		needsNewMap = true
		g.TileId--
		if g.TileId == -1 {
			g.TileId = len(g.TileSheet.TileTypes) - 1
		}
	} else if mi.JustPressed(pixelgl.KeyA) {
		needsNewMap = true
		g.Pos = g.Pos.Add(pixel.V(-1, 0))
	} else if mi.JustPressed(pixelgl.KeyW) {
		needsNewMap = true
		g.Pos = g.Pos.Add(pixel.V(0, 1))
	} else if mi.JustPressed(pixelgl.KeyD) {
		needsNewMap = true
		g.Pos = g.Pos.Add(pixel.V(1, 0))
	} else if mi.JustPressed(pixelgl.KeyS) {
		needsNewMap = true
		g.Pos = g.Pos.Add(pixel.V(0, -1))
	}

	g.setMinMax()
	if needsNewMap {
		x := int(g.Pos.X)
		y := int(g.Pos.Y)

		maxY := len(g.MapOpts.Grid) - 1

		g.MapOpts.Grid[maxY-y][x] = g.TileId
		g.Ground = maps.NewMap(g.MapOpts)
	}
}

func (g *WorldBuilder) setMinMax() {

	maxX := len(g.MapOpts.Grid[0])

	if int(g.Pos.X) == maxX {
		g.addCol(true)
	}

	if g.Pos.X < 0 {
		g.cam = g.cam.Add(pixel.V(sheet.TileSize, 0))
		g.Pos.X = 0
		g.addCol(false)
	}

	maxY := len(g.MapOpts.Grid)

	if int(g.Pos.Y) == maxY {
		g.addRow(false)
	}

	if g.Pos.Y < 0 {
		g.cam = g.cam.Add(pixel.V(0, sheet.TileSize))
		g.Pos.Y = 0
		g.addRow(true)
	}
}

func (g *WorldBuilder) addCol(toEnd bool) {
	// adding a col means expanding a row
	for i, row := range g.MapOpts.Grid {
		if toEnd {
			g.MapOpts.Grid[i] = append(row, 0)
		} else {
			g.MapOpts.Grid[i] = append([]int{0}, row...)
		}
	}
}

func (g *WorldBuilder) addRow(toEnd bool) {
	l := len(g.MapOpts.Grid[0])
	row := make([]int, l)

	if toEnd {
		g.MapOpts.Grid = append(g.MapOpts.Grid, row)
	} else {
		g.MapOpts.Grid = append([][]int{row}, g.MapOpts.Grid...)
	}
}

func NewWorldBuilder() Scene {

	opts := &maps.MapOpts{
		Sheet: "ground-tile-sheet",
		Grid: [][]int{
			{0},
		},
	}

	buildingNames := items.GetClassNames("Cinder Block")
	buildingNames = append(buildingNames, items.GetClassNames("Conveyor Belt")...)

	mapOne := maps.NewMap(opts)

	return &WorldBuilder{
		mode:          "normal",
		Ground:        mapOne,
		MapOpts:       opts,
		TileSheet:     sheet.GetTileSheet(opts.Sheet),
		cam:           pixel.V(0, 0),
		Pos:           pixel.V(0, 0),
		TileId:        0,
		path:          "",
		structurePath: "",
		buildingNames: buildingNames,
	}
}

func (g *WorldBuilder) GetMap() *maps.Map {
	return g.Ground
}

func (g *WorldBuilder) GetCamera() pixel.Vec {
	return g.cam
}

func (g *WorldBuilder) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *WorldBuilder) save(mi characters.MindInput) {
	b, err := json.Marshal(g.MapOpts)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(g.path, b, 0644)
	if err != nil {
		panic(err)
	}

	items := mi.ListItems()
	b, err = json.Marshal(items)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(g.structurePath, b, 0644)
	if err != nil {
		panic(err)
	}

}

func (g *WorldBuilder) Enter(mi characters.MindInput) {
	g.path = os.Args[2]
	g.structurePath = os.Args[3]

	file, err := ioutil.ReadFile(g.path)
	sheet := g.MapOpts.Sheet

	if err == nil {
		err = json.Unmarshal(file, g.MapOpts)
		if err != nil {
			panic(err)
		}
		g.MapOpts.Sheet = sheet
		if g.MapOpts.Locations == nil {
			g.MapOpts.Locations = make(map[string]pixel.Rect)
		}
		g.Ground = maps.NewMap(g.MapOpts)
	}

	file, err = ioutil.ReadFile(g.structurePath)
	if err == nil {
		items := make([]*items.Item, 0)
		err = json.Unmarshal(file, &items)
		if err != nil {
			panic(err)
		}

		for _, item := range items {
			mi.AddItem(item)
		}
	}

}

func (g *WorldBuilder) Exit(mi characters.MindInput) {
}
