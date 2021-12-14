package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	gameMap "github.com/Stroby241/TimeTravelGame/src/map"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventGameLoad, load)
}

type game struct {
	m   *gameMap.Map
	cam *util.Camera
}

func load(data interface{}) {
	g := &game{
		m:   nil,
		cam: util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		update(g)
	})

	drawId := event.On(event.EventDraw, func(data interface{}) {
		draw(data.(*ebiten.Image), g)
	})

	loadMapId := event.On(event.EventGameLoadMap, func(data interface{}) {
		g.m = loadMap(data.(string))
	})

	submitRoundId := event.On(event.EventGameSubmitRound, func(data interface{}) {
		g.m.U.SubmitRound()
		g.m.Update()
	})

	var unloadId event.ReciverId
	event.On(event.EventGameUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)
		event.UnOn(event.EventGameLoadMap, loadMapId)
		event.UnOn(event.EventGameSubmitRound, submitRoundId)

		event.UnOn(event.EventGameUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageGame)
}

func update(g *game) {
	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	getTile := func() *field.Tile {
		mat := *g.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := g.m.GetCard(clickPos)
		return tile
	}

	if g.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile := getTile()
		if tile == nil {
			return
		}

		_, _, unit := g.m.U.GetUnitAtPos(tile.AxialPos)
		if unit != nil {
			g.m.U.SetSelector(unit.Pos)
		}

		g.m.Update()
	} else if g.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile := getTile()
		if tile == nil {
			return
		}

		_, _, unit := g.m.U.GetUnitAtPos(g.m.U.SelectedUnit)

		if unit != nil && tile.Visable {
			g.m.U.SetAction(unit, tile.AxialPos)
		}

		g.m.Update()
	}
}

func draw(screen *ebiten.Image, g *game) {
	if g.m != nil {
		g.m.Draw(screen, g.cam)
	}
}

func loadMap(name string) *gameMap.Map {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}
	m := field.Load(buffer)
	return m
}
