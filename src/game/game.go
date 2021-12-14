package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventGameLoad, load)
}

type game struct {
	f   *field.Field
	cam *util.Camera
}

func load(data interface{}) {
	g := &game{
		f:   nil,
		cam: util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		update(g)
	})

	drawId := event.On(event.EventDraw, func(data interface{}) {
		draw(data.(*ebiten.Image), g)
	})

	loadMapId := event.On(event.EventGameLoadMap, func(data interface{}) {
		g.f = field.LoadField(data.(string))
	})

	submitRoundId := event.On(event.EventGameSubmitRound, func(data interface{}) {
		g.f.U.SubmitRound()
		g.f.Update()
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

		tile := g.f.GetCard(clickPos)
		return tile
	}

	if g.f != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile := getTile()
		if tile == nil {
			return
		}

		_, _, unit := g.f.U.GetUnitAtPos(tile.AxialPos)
		if unit != nil {
			g.f.S.Pos = unit.Pos
			g.f.S.Visible = true
		}

		g.f.Update()
	} else if g.f != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile := getTile()
		if tile == nil {
			return
		}

		_, _, unit := g.f.U.GetUnitAtPos(g.f.S.Pos)

		if unit != nil && tile.Visable {
			g.f.U.SetAction(unit, tile.AxialPos)
		}

		g.f.Update()
	}
}

func draw(screen *ebiten.Image, g *game) {
	if g.f != nil {
		g.f.Draw(screen, g.cam)
	}
}
