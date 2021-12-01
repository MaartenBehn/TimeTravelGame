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
	event.On(event.EventEditorLoad, editorLoad)
}

type editor struct {
	m   *gameMap.Map
	cam *util.Camera
}

func editorLoad(data interface{}) {
	e := &editor{
		m:   nil,
		cam: util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		editorUpdate(e.m, e.cam)
	})
	drawId := event.On(event.EventDraw, func(data interface{}) {
		editorDraw(data.(*ebiten.Image), e.m, e.cam)
	})

	newMapId := event.On(event.EventEditorNewMap, func(data interface{}) {
		e.m = editorNewMap(data.(CardPos))
	})
	saveMapId := event.On(event.EventEditorSaveMap, func(data interface{}) {
		editorSaveMap(data.(string), e.m)
	})
	loadMapId := event.On(event.EventEditorLoadMap, func(data interface{}) {
		e.m = editorLoadMap(data.(string))
	})

	var unloadId event.ReciverId
	unloadId = event.On(event.EventEditorUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)
		event.UnOn(event.EventEditorNewMap, newMapId)
		event.UnOn(event.EventEditorSaveMap, saveMapId)
		event.UnOn(event.EventEditorLoadMap, loadMapId)

		event.UnOn(event.EventEditorUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageMapEditor)
}

func editorUpdate(m *gameMap.Map, cam *util.Camera) {
	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	if m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mapPos := CardPos{}
		mat := *cam.GetMatrix()
		mat.Invert()

		mapPos.X, mapPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := m.Get(mapPos.ToAxial())
		tile.Visable = true
		m.Update()
	} else if m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mapPos := CardPos{}
		mat := *cam.GetMatrix()
		mat.Invert()

		mapPos.X, mapPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := m.Get(mapPos.ToAxial())
		tile.Visable = false
		m.Update()
	}
}

func editorDraw(screen *ebiten.Image, m *gameMap.Map, cam *util.Camera) {
	if m != nil {
		m.DrawMap(screen, cam)
	}
}

func editorNewMap(size CardPos) *gameMap.Map {
	m := gameMap.NewMap(size)
	return m
}

func editorSaveMap(name string, m *gameMap.Map) {
	if m == nil {
		return
	}
	util.SaveMapBufferToFile(name, m.Save())
}

func editorLoadMap(name string) *gameMap.Map {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}
	m := gameMap.LoadMap(buffer)
	return m
}
