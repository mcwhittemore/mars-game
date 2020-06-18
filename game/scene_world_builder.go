package game

import (
	"app/characters"
	"app/maps"
	"app/sheet"

	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type WorldBuilder struct {
	mode    string
	cam     pixel.Vec
	MapOpts *maps.MapOpts
	Ground  *maps.Map
	Pos     pixel.Vec
	TileId  int
	path    string
}

func (g *WorldBuilder) Update(dt float64, mi characters.MindInput) {

	if g.mode == "input" {
		g.inputMode(mi)
	} else if g.mode == "normal" {
		g.normalMode(mi)
	}

	mi.KeepInView(g.Pos.Scaled(64), 128)

	imd := imdraw.New(nil)

	if g.mode == "input" {
		imd.Color = pixel.RGB(1, 0, 0)
	} else if g.mode == "normal" {
		imd.Color = pixel.RGB(0, 0, 1)
	}

	pos := g.Pos.Scaled(64).Sub(pixel.V(32, 32))
	imd.Push(pos, pos.Add(pixel.V(64, 64)))
	imd.Rectangle(2)

	mi.AddDraw(imd)
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
		g.save()
	} else if mi.JustPressed(pixelgl.KeyJ) {
		needsNewMap = true
		g.TileId++
		if g.TileId == 13 {
			g.TileId = 0
		}
	} else if mi.JustPressed(pixelgl.KeyK) {
		needsNewMap = true
		g.TileId--
		if g.TileId == -1 {
			g.TileId = 12
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
		g.cam = g.cam.Add(pixel.V(64, 0))
		g.Pos.X = 0
		g.addCol(false)
	}

	maxY := len(g.MapOpts.Grid)

	if int(g.Pos.Y) == maxY {
		g.addRow(false)
	}

	if g.Pos.Y < 0 {
		g.cam = g.cam.Add(pixel.V(0, 64))
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

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{X: 20, Y: 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	opts := &maps.MapOpts{
		Sheet:     groundSheet,
		Tiles:     []*maps.Tile{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}, {2, 0}, {4, 0}, {4, 2}, {2, 2}},
		TileTypes: []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		// 0: empty
		// 1: |_
		// 2: __
		// 3: _|
		// 4:  |
		// 5: ⎻|
		// 6: ⎻⎻
		// 7: |⎻
		// 8: |
		// 9:  ⎻| alt
		// 10: |⎻ alt
		// 11: |_ alt
		// 12: _| alt
		Grid: [][]int{
			{0},
		},
	}

	mapOne := maps.NewMap(opts)

	return &WorldBuilder{
		mode:    "normal",
		Ground:  mapOne,
		MapOpts: opts,
		cam:     pixel.V(0, 0),
		Pos:     pixel.V(0, 0),
		TileId:  0,
		path:    "",
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

func (g *WorldBuilder) save() {
	b, err := json.Marshal(g.MapOpts)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(g.path, b, 0644)
	if err != nil {
		panic(err)
	}
}

func (g *WorldBuilder) Enter(mi characters.MindInput) {
	g.path = os.Args[2]
	file, err := ioutil.ReadFile(g.path)
	sheet := g.MapOpts.Sheet
	if err == nil {
		err = json.Unmarshal(file, g.MapOpts)
		if err != nil {
			panic(err)
		}
		g.MapOpts.Sheet = sheet
		g.Ground = maps.NewMap(g.MapOpts)
	}
}

func (g *WorldBuilder) Exit(mi characters.MindInput) {
}
