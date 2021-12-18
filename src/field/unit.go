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

func (t *Timeline) GetUnitAtPos(fieldPos CardPos, pos AxialPos) (int, *Unit) {
	for i, unit := range t.Units {
		if unit.FieldPos == fieldPos && unit.Pos == pos  {
			return  i, unit
		}
	}
	return -1, nil
}

func (t *Timeline) AddUnitAtTile(field *Field, tile *Tile, fraction *Fraction) *Unit {
	_, unit := t.GetUnitAtPos(field.Pos, tile.AxialPos)
	if unit == nil && tile.Visable {
		id := getFractionIndex(fraction)

		unit = NewUnit(field.Pos ,tile.AxialPos, id)
		t.Units = append(t.Units, unit)
	}
	return unit
}

func (t *Timeline) RemoveUnitAtTile(field *Field, tile *Tile) {
	t.RemoveUnitAtPos(field.Pos, tile.AxialPos)
}
func (t *Timeline) RemoveUnitAtPos(fieldPos CardPos, pos AxialPos) {
	i, unit := t.GetUnitAtPos(fieldPos, pos)
	if unit != nil {
		t.Units = append(t.Units[:i], t.Units[i+1:]...)
	}
}

type Unit struct {
	FactionId int
	FieldPos  CardPos
	Pos       AxialPos
	Action    *Action
	Support   int
}

func NewUnit(fieldPos CardPos, pos AxialPos, factionId int) *Unit {
	return &Unit{
		FactionId: factionId,
		FieldPos:  fieldPos,
		Pos:       pos,
		Action:    NewAction(),
	}
}

func (u *Unit) draw(img *ebiten.Image, fraction *Fraction, f *Field) {

	w, h := fraction.Images["unit"].Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.GeoM{}
	tile := f.GetAxial(u.Pos)
	op.GeoM.Translate(f.Pos.X + tile.Pos.X  - float64(w)/2, f.Pos.Y + tile.Pos.Y - float64(h)/2)
	img.DrawImage(fraction.Images["unit"], op)
}

func (u *Unit) copyToField(field *Field) *Unit {
	copyUnit := NewUnit(field.Pos, u.Pos, u.FactionId)

	copyUnit.Action.Kind = u.Action.Kind
	copyUnit.Action.ToPos = u.Action.ToPos
	copyUnit.Action.Support = u.Action.Support
	copyUnit.Action.ToFieldPos = field.Pos

	copyUnit.Support = u.Support

	return copyUnit
}
