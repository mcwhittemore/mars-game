package game

import (
	"app/characters"
	"app/items"

	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type GameState struct {
	characters   map[string]*characters.CharacterData
	items        []*items.Item
	draws        []*imdraw.IMDraw
	texts        []*text.Text
	sceneManager *SceneManager
	gameTime     float64
	win          *pixelgl.Window
}

func NewGameState(win *pixelgl.Window) *GameState {
	sm := NewSceneManager()
	c := make(map[string]*characters.CharacterData, 0)
	return &GameState{
		characters:   c,
		sceneManager: sm,
		gameTime:     0,
		win:          win,
	}
}

/*
 * GAME LOOP
 */

func (gs *GameState) Update(dt float64) {
	for _, cd := range gs.characters {
		cd.Update(dt, gs)
	}

	for _, item := range gs.items {
		item.State.Update(item, dt, gs)
	}

	gs.sceneManager.Current.Update(dt, gs)
}

func (gs *GameState) Render(win *pixelgl.Window) {
	campos := gs.sceneManager.Current.GetCamera()
	cam := pixel.IM.Moved(gs.win.Bounds().Center().Sub(campos))
	gs.win.SetMatrix(cam)

	activeMap := gs.sceneManager.Current.GetMap()
	activeMap.Render.Draw(win)

	wallsBatch := items.ItemSheets[items.Wall_Sheet].GetBatch()
	//wallsBatch := pixel.NewBatch(&pixel.TrianglesData{}, items.GetSheet(items.Wall_Sheet))
	for _, item := range gs.items {
		sprite, im := item.GetSprite()
		matrix := im.Moved(item.Pos)

		if item.Sheet == items.Wall_Sheet {
			sprite.Draw(wallsBatch, matrix)
		} else {
			sprite.Draw(win, matrix)
		}
	}
	wallsBatch.Draw(win)

	for _, cd := range gs.characters {
		cd.Render(win)
	}

	for _, imd := range gs.draws {
		imd.Draw(win)
	}

	for _, text := range gs.texts {
		text.Draw(win, pixel.IM)
	}

	// keeps cap but sets len to 0
	gs.draws = gs.draws[:0]
	gs.texts = gs.texts[:0]
}

/*
 * DATA ACCESS AND STATE CONTROL
 */

func (gs *GameState) AddDraw(imd *imdraw.IMDraw) {
	gs.draws = append(gs.draws, imd)
}

func (gs *GameState) AddText(text *text.Text) {
	gs.texts = append(gs.texts, text)
}

func (gs *GameState) AddItem(item *items.Item) {
	gs.items = append(gs.items, item)
}

func (gs *GameState) GetItem(tp pixel.Vec) *items.Item {
	maxArea := float64(0)
	var match *items.Item
	for _, item := range gs.items {
		tr := item.PosBounds(tp)
		ir := item.PosBounds(item.Pos)

		ol := tr.Intersect(ir)
		if ol.Area() > maxArea {
			match = item
			maxArea = ol.Area()
		}
	}

	return match
}

func (gs *GameState) RemoveItem(t *items.Item) {
	for i, item := range gs.items {
		if item == t {
			end := len(gs.items) - 1
			gs.items[i] = gs.items[end]
			gs.items[end] = nil
			gs.items = gs.items[:end]
			return
		}
	}
}

func (gs *GameState) ListItems() []*items.Item {
	return gs.items
}

func (gs *GameState) ShowCharacter(name string, c *characters.Character) {
	gs.characters[name].Character = c
}

func (gs *GameState) HideCharacter(name string) {
	gs.characters[name].Character = nil
}

func (gs *GameState) AddCharacter(name string, cd *characters.CharacterData) {
	if _, ok := gs.characters[name]; ok {
		panic(fmt.Sprintf("Multiple characters with the name %s", name))
	}

	if cd == nil {
		gs.characters[name] = characters.NewCharacterData(name)
	} else {
		gs.characters[name] = cd
	}
}

func (gs *GameState) RemoveCharacter(name string) {
	gs.characters[name] = nil
}

func (gs *GameState) GetCharacter(name string) *characters.CharacterData {
	return gs.characters[name]
}

func (gs *GameState) AddScene(name string, sb SceneBuilder) {
	gs.sceneManager.AddScene(name, sb)
}

func (gs *GameState) ChangeScene(name string) {
	gs.sceneManager.ChangeScene(name, gs)
}

func (gs *GameState) GetLocation(name string) pixel.Rect {
	am := gs.sceneManager.Current.GetMap()
	return am.Locations[name]
}

func (gs *GameState) GetCollideRect(rect pixel.Rect, thing interface{}) (pixel.Rect, *characters.Character) {
	var out pixel.Rect

	for _, cd := range gs.characters {
		if cd.Character == nil {
			continue
		}

		if thing != interface{}(cd.Character) {
			out = cd.Character.PosBounds(cd.Character.Pos).Intersect(rect)
			if out != pixel.ZR {
				return out, cd.Character
			}
		}
	}

	return pixel.ZR, nil
}

func (gs *GameState) GetHeroPos() pixel.Vec {

	for _, cd := range gs.characters {
		if cd.Character == nil {
			continue
		}

		if cd.Name == "hero" {
			return cd.Character.Pos
		}
	}

	// this should never happen but it could...
	return pixel.ZV
}

func (gs *GameState) IsObstacle(pos pixel.Vec) bool {
	activeMap := gs.sceneManager.Current.GetMap()
	return activeMap.IsObstacle(pos)
}

func (gs *GameState) JustPressed(button pixelgl.Button) bool {
	return gs.win.JustPressed(button)
}

func (gs *GameState) Pressed(button pixelgl.Button) bool {
	return gs.win.Pressed(button)
}

func (gs *GameState) Typed() string {
	return gs.win.Typed()
}

func (gs *GameState) KeepInView(pos pixel.Vec, buffer float64) {
	bds := gs.win.Bounds()
	campos := gs.sceneManager.Current.GetCamera()
	cam := pixel.IM.Moved(gs.win.Bounds().Center().Sub(campos))

	viewBox := pixel.Rect{
		Min: cam.Unproject(bds.Min),
		Max: cam.Unproject(bds.Max),
	}

	edgeDir := [4]pixel.Vec{{X: -1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 0, Y: -1}}

	viewbox := viewBox.Edges()
	for i, edge := range viewbox {
		closest := edge.Closest(pos)
		dis := pixel.L(pos, closest).Len()
		if dis < buffer {
			mov := edgeDir[i].Scaled(buffer - dis)
			campos = campos.Add(mov)
			gs.sceneManager.Current.SetCamera(campos)
			break
		}
	}

}
