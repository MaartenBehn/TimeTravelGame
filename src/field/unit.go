package field

import (
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Unit struct {
	TimePos
	MovePattern
	FactionId int
	Action    Action
	Support   int
}

func NewUnit(pos TimePos, factionId int, pattern MovePattern) *Unit {
	return &Unit{
		FactionId:   factionId,
		TimePos:     pos,
		Action:      NewAction(),
		MovePattern: pattern,
	}
}

func (u *Unit) draw(img *ebiten.Image, fraction *util.Fraction) {

	w, h := fraction.Images["unit"].Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.GeoM{}
	op.GeoM.Translate(u.CalcPos().X-float64(w)/2, u.CalcPos().Y-float64(h)/2)
	img.DrawImage(fraction.Images["unit"], op)
}

func (t *Timeline) makeReadyUnits() {
	t.moveUnits = []*Unit{}
	t.supportUnits = []*Unit{}

	for _, unit := range t.Units {
		if unit.Action.Kind == actionMove {
			t.moveUnits = append(t.moveUnits, unit)
		} else if unit.Action.Kind == actionSupport {
			t.supportUnits = append(t.supportUnits, unit)
		}
	}
}

func (t *Timeline) GetUnitAtPos(pos TimePos) (int, *Unit) {
	for i, unit := range t.Units {
		if unit.SamePos(pos) {
			return i, unit
		}
	}
	return -1, nil
}

func (t *Timeline) AddUnitAtTile(pos TimePos, fraction *util.Fraction) *Unit {
	_, unit := t.GetUnitAtPos(pos)
	if unit == nil {
		id := util.GetFractionIndex(fraction)

		unit = NewUnit(pos, id, BasicMovePattern{Stride: 1})
		t.Units = append(t.Units, unit)
	}
	return unit
}

func (t *Timeline) RemoveUnitAtTile(tile *Tile) {
	t.RemoveUnitAtPos(tile.TimePos)
}
func (t *Timeline) RemoveUnitAtPos(pos TimePos) {
	i, unit := t.GetUnitAtPos(pos)
	if unit != nil {
		t.Units = append(t.Units[:i], t.Units[i+1:]...)
	}
}

func (t *Timeline) CopyUnit(unit *Unit) *Unit {
	fromUnit := *unit
	copyUnit := fromUnit
	t.Units = append(t.Units, &copyUnit)
	return &copyUnit
}
