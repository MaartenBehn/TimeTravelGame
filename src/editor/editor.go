package editor

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
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
	f    *field.Field
	cam  *util.Camera
	mode int
}

func load(data interface{}) {
	e := &editor{
		f:    nil,
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
		e.f = newField(data.(int))
	})
	saveMapId := event.On(event.EventEditorSaveMap, func(data interface{}) {
		saveField(data.(string), e)
	})
	loadMapId := event.On(event.EventEditorLoadMap, func(data interface{}) {
		e.f = loadField(data.(string))
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

	getTile := func() *field.Tile {
		mat := *e.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := e.f.GetCard(clickPos)
		return tile
	}

	if e.f != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile := getTile()
		if tile == nil {
			return
		}

		if e.mode == 0 {
			tile.Visable = true
		} else if e.mode == 1 && tile.Visable {
			e.f.U.AddUnitAtTile(tile, &field.Fractions[1])
		} else if e.mode == 2 {
			e.f.U.AddUnitAtTile(tile, &field.Fractions[0])
		} else if e.mode == 3 && tile.Visable {
			_, _, unit := e.f.U.GetUnitAtPos(tile.AxialPos)
			if unit != nil {
				e.f.U.SetSelector(unit.Pos)
			}
		}

		e.f.Update()
	} else if e.f != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile := getTile()
		if tile == nil {
			return
		}

		if e.mode == 0 {
			tile.Visable = false
			e.f.U.RemoveUnitAtTile(tile)

		} else if (e.mode == 1 || e.mode == 2) && tile.Visable {
			e.f.U.RemoveUnitAtTile(tile)
		} else if e.mode == 3 && tile.Visable {
			_, _, unit := e.f.U.GetUnitAtPos(e.f.U.SelectedUnit)

			if unit != nil && tile.Visable {
				e.f.U.SetAction(unit, tile.AxialPos)
			}
		}

		e.f.Update()
	}
}

func draw(screen *ebiten.Image, e *editor) {
	if e.f != nil {
		e.f.Draw(screen, e.cam)
	}
}

func newField(size int) *field.Field {
	m := field.NewField(size)
	return m
}

func saveField(name string, e *editor) {
	if e.m == nil {
		return
	}
	util.SaveMapBufferToFile(name, e.f.Save())
}

func loadField(name string) *field.Field {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}
	m := field.Load(buffer)
	return m
}
