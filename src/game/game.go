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
	t   *field.Timeline
	cam *util.Camera
}

func load(data interface{}) {
	g := &game{
		t:   nil,
		cam: util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		update(g)
	})

	drawId := event.On(event.EventDraw, func(data interface{}) {
		draw(data.(*ebiten.Image), g)
	})

	loadMapId := event.On(event.EventGameLoadMap, func(data interface{}) {
		g.t = field.LoadTimeline(data.(string))
	})

	submitRoundId := event.On(event.EventGameSubmitRound, func(data interface{}) {
		g.t.SubmitRound()
		g.t.Update()
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

	getTile := func() (*field.Tile, *field.Field) {
		mat := *g.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, field := g.t.Get(clickPos)
		return tile, field
	}

	if g.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile, field := getTile()
		if tile == nil {
			return
		}

		_, unit := g.t.GetUnitAtPos(field.Pos, tile.AxialPos)
		if unit != nil {
			g.t.S.FieldPos = field.Pos
			g.t.S.Pos = unit.Pos
			g.t.S.Visible = true
		}

		g.t.Update()
	} else if g.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile, field := getTile()
		if tile == nil {
			return
		}

		_, unit := g.t.GetUnitAtPos(g.t.S.FieldPos, g.t.S.Pos)

		if unit != nil && tile.Visable {
			g.t.SetAction(unit, field.Pos, tile.AxialPos)
		}

		g.t.Update()
	}
}

func draw(screen *ebiten.Image, g *game) {
	if g.t != nil {
		g.t.Draw(screen, g.cam)
	}
}
