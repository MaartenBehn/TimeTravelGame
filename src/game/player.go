package game

import (
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type player struct {
	userData
}

func NewPlayer(factionId int, t *field.Timeline) *player {
	return &player{
		userData: NewUserData(
			true,
			factionId,
			t,
			util.NewCamera(
				CardPos{0, 0},
				CardPos{50000, 50000},
				CardPos{1, 1},
				CardPos{10, 10},
				factionId)),
	}
}

func (p *player) update() {
	p.cam.UpdateInput()

	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	getTile := func() (*field.Tile, *field.Field) {
		mat := *p.cam.GetMatrix()
		mat.Invert()

		clickPos := CardPos{}
		clickPos.X, clickPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, field := p.t.Get(clickPos)
		return tile, field
	}

	if p.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		tile, _ := getTile()
		if tile == nil {
			return
		}

		_, unit := p.t.GetUnitAtPos(tile.TimePos)
		if unit != nil {
			p.t.S.TimePos = unit.TimePos
			p.t.S.Visible = true
			p.t.Update()
		}
	} else if p.t != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		tile, _ := getTile()
		if tile == nil {
			return
		}

		_, unit := p.t.GetUnitAtPos(p.t.S.TimePos)
		field := p.t.Fields[unit.FieldPos]

		if unit != nil &&
			unit.FactionId == p.factionId &&
			tile.Visable && field.Active {
			p.t.SetAction(unit, tile.TimePos)
			p.t.Update()
		}
	}
}
