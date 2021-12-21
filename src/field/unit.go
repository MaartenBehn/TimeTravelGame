package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
)

type Fraction struct {
	name       string
	color      color.Color
	colorLigth color.Color

	Images map[string]*ebiten.Image
}

var Fractions = []Fraction{
	{
		name:  "red",
		color: colornames.Red,
	},
	{
		name:  "blue",
		color: colornames.Blue,
	},
}

func getFractionIndex(f *Fraction) int {
	for i, fraction := range Fractions {
		if fraction.name == f.name {
			return i
		}
	}
	return -1
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

func (t *Timeline) AddUnitAtTile(pos TimePos, fraction *Fraction) *Unit {
	_, unit := t.GetUnitAtPos(pos)
	if unit == nil {
		id := getFractionIndex(fraction)

		unit = NewUnit(pos, id)
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

type Unit struct {
	TimePos
	FactionId int
	Action    *Action
	Support   int
}

func NewUnit(pos TimePos, factionId int) *Unit {
	return &Unit{
		FactionId: factionId,
		TimePos:   pos,
		Action:    NewAction(),
	}
}

func (u *Unit) draw(img *ebiten.Image, fraction *Fraction) {

	w, h := fraction.Images["unit"].Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.GeoM{}
	op.GeoM.Translate(u.CalcPos().X-float64(w)/2, u.CalcPos().Y-float64(h)/2)
	img.DrawImage(fraction.Images["unit"], op)
}

func (u *Unit) copyToField(fieldPos CardPos) *Unit {
	fromUnit := *u
	copyUnit := fromUnit
	copyUnit.FieldPos = fieldPos
	return &copyUnit
}
