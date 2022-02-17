package editor

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventEditorLoad, load)
}

type editor struct {
	t    *field.Timeline
	cam  *util.Camera
	mode int
}

func load(data interface{}) {
	e := &editor{
		t:    nil,
		cam:  util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}, -1),
		mode: 0,
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		update(e)
	})
	drawId := event.On(event.EventDraw, func(data interface{}) {
		draw(data.(*ebiten.Image), e)
	})
	newMapId := event.On(event.EventEditorUINewMap, func(data interface{}) {
		e.t = field.NewTimeline(data.(int))
		e.t.AddField(CardPos{})
	})
	saveMapId := event.On(event.EventEditorUISaveMap, func(data interface{}) {
		e.t.Save(data.(string))
	})
	loadMapId := event.On(event.EventEditorUILoadMap, func(data interface{}) {
		e.t = field.Load(data.(string))
	})
	modeId := event.On(event.EventEditorUISetMode, func(data interface{}) {
		e.mode = data.(int)
	})

	var unloadId event.ReciverId
	unloadId = event.On(event.EventEditorUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)
		event.UnOn(event.EventEditorUINewMap, newMapId)
		event.UnOn(event.EventEditorUISaveMap, saveMapId)
		event.UnOn(event.EventEditorUILoadMap, loadMapId)
		event.UnOn(event.EventEditorUISetMode, modeId)

		event.UnOn(event.EventEditorUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageMapEditor)
}

func update(e *editor) {
	e.cam.UpdateInput()

	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	getTile := func() (*field.Tile, *field.Field) {
		mat := *e.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, field := e.t.Get(clickPos)
		return tile, field
	}

	if e.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile, f := getTile()
		if tile == nil {
			return
		}

		if e.mode == 0 {
			tile.Visable = true
		} else if e.mode == 1 && tile.Visable {
			e.t.AddUnitAtTile(tile.TimePos, &util.Fractions[1])
		} else if e.mode == 2 {
			e.t.AddUnitAtTile(tile.TimePos, &util.Fractions[0])
		} else if e.mode == 3 && tile.Visable {
			_, unit := e.t.GetUnitAtPos(tile.TimePos)
			if unit != nil {
				e.t.S.FieldPos = f.Pos
				e.t.S.TilePos = unit.TilePos
				e.t.S.Visible = true
			}
		}

		e.t.Update()
	} else if e.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile, _ := getTile()
		if tile == nil {
			return
		}

		if e.mode == 0 {
			tile.Visable = false
			e.t.RemoveUnitAtTile(tile)

		} else if (e.mode == 1 || e.mode == 2) && tile.Visable {
			e.t.RemoveUnitAtTile(tile)
		} else if e.mode == 3 && tile.Visable && e.t.S.Visible {
			_, unit := e.t.GetUnitAtPos(e.t.S.TimePos)

			if unit != nil && tile.Visable {
				e.t.SetAction(unit, tile.TimePos)
			}
		}

		e.t.Update()
	}
}

func draw(screen *ebiten.Image, e *editor) {
	if e.t == nil {
		return
	}

	e.t.Draw(screen, e.cam)
}
