package editor

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/map"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventEditorLoad, load)
}

type editor struct {
	m    *gameMap.Map
	cam  *util.Camera
	mode int
}

func load(data interface{}) {
	e := &editor{
		m:    nil,
		cam:  util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
		mode: 0,
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		update(e)
	})
	drawId := event.On(event.EventDraw, func(data interface{}) {
		draw(data.(*ebiten.Image), e)
	})
	newMapId := event.On(event.EventEditorNewMap, func(data interface{}) {
		e.m = newMap(data.(CardPos))
	})
	saveMapId := event.On(event.EventEditorSaveMap, func(data interface{}) {
		saveMap(data.(string), e)
	})
	loadMapId := event.On(event.EventEditorLoadMap, func(data interface{}) {
		e.m = loadMap(data.(string))
	})
	modeId := event.On(event.EventEditorSetMode, func(data interface{}) {
		e.mode = data.(int)
	})

	var unloadId event.ReciverId
	unloadId = event.On(event.EventEditorUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)
		event.UnOn(event.EventEditorNewMap, newMapId)
		event.UnOn(event.EventEditorSaveMap, saveMapId)
		event.UnOn(event.EventEditorLoadMap, loadMapId)
		event.UnOn(event.EventEditorSetMode, modeId)

		event.UnOn(event.EventEditorUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageMapEditor)
}

func update(e *editor) {
	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	getTile := func() *gameMap.Tile {
		mat := *e.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := e.m.GetCard(clickPos)
		return tile
	}

	if e.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile := getTile()

		if e.mode == 0 {
			tile.Visable = true
		} else if e.mode == 1 && tile.Visable {
			e.m.U.AddUnitAtTile(tile, &gameMap.FractionBlue)
		} else if e.mode == 2 {
			e.m.U.AddUnitAtTile(tile, &gameMap.FractionRed)
		} else if e.mode == 3 && tile.Visable {
			_, _, unit := e.m.U.GetUnitAtPos(tile.AxialPos)
			if unit != nil {
				e.m.U.SetSelector(unit.Pos)
			}
		}

		e.m.Update()
	} else if e.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile := getTile()

		if e.mode == 0 {
			tile.Visable = false
			e.m.U.RemoveUnitAtTile(tile)

		} else if (e.mode == 1 || e.mode == 2) && tile.Visable {
			e.m.U.RemoveUnitAtTile(tile)
		} else if e.mode == 3 && tile.Visable {
			_, _, unit := e.m.U.GetUnitAtPos(e.m.U.SelectedUnit)

			if unit != nil && tile.Visable {
				e.m.U.SetAction(unit, tile.AxialPos)
			}
		}

		e.m.Update()
	}
}

func draw(screen *ebiten.Image, e *editor) {
	if e.m != nil {
		e.m.Draw(screen, e.cam)
	}
}

func newMap(size CardPos) *gameMap.Map {
	m := gameMap.NewMap(size)
	return m
}

func saveMap(name string, e *editor) {
	if e.m == nil {
		return
	}
	util.SaveMapBufferToFile(name, e.m.Save())
}

func loadMap(name string) *gameMap.Map {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}
	m := gameMap.Load(buffer)
	return m
}
