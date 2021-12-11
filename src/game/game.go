package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
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
		submitRound(g)
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

	getTile := func() *gameMap.Tile {
		mat := *g.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := g.m.GetCard(clickPos)
		return tile
	}

	if g.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile := getTile()

		_, _, unit := g.m.UnitController.GetUnitAtPos(tile.AxialPos)
		if unit != nil {
			g.m.UnitController.SetSelector(unit.Pos)
		}

		g.m.Update()
	} else if g.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile := getTile()

		_, _, unit := g.m.UnitController.GetUnitAtPos(g.m.UnitController.SelectedUnit)

		if unit != nil && tile.Visable {
			unit.TargetPos = tile.AxialPos
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
	m := gameMap.Load(buffer)
	return m
}

func submitRound(g *game) {
	for _, chunk := range g.m.Chunks {
		for _, tile := range chunk.Tiles {
			tile.TargetOf = []*gameMap.Unit{}
		}
	}

	targetTiles := make([]*gameMap.Tile, 0)
	for _, units := range g.m.UnitController.Units {
		for _, unit := range units {
			tile, _ := g.m.GetAxial(unit.TargetPos)
			tile.TargetOf = append(tile.TargetOf, unit)

			isContained := false
			for _, targetTile := range targetTiles {
				if targetTile == tile {
					isContained = true
					break
				}
			}
			if !isContained {
				targetTiles = append(targetTiles, tile)
			}
		}
	}

	for _, tile := range targetTiles {
		if len(tile.TargetOf) == 1 {

			tile.TargetOf[0].Pos = tile.AxialPos

		} else if len(tile.TargetOf) > 1 {

			factions := make([]int, len(g.m.UnitController.Units))
			for _, u := range tile.TargetOf {
				factions[u.FactionId] += 1
			}

			winningFaction := -1
			winningFactionAmmount := 0
			for i, faction := range factions {
				if winningFactionAmmount < faction {
					winningFactionAmmount = faction
					winningFaction = i
				} else if winningFactionAmmount == faction {
					winningFaction = -1
				}
			}

			if winningFaction != -1 {
				for _, u := range tile.TargetOf {
					if u.FactionId == winningFaction {
						_, _, presentUnit := g.m.UnitController.GetUnitAtPos(tile.AxialPos)

						if presentUnit != nil && presentUnit != u {
							g.m.UnitController.RemoveUnitAtTile(tile)
						}

						u.Pos = tile.AxialPos
						break
					}
				}
			}
		}
	}

	for _, units := range g.m.UnitController.Units {
		for _, unit := range units {
			unit.TargetPos = unit.Pos
		}
	}

	g.m.Update()
}
