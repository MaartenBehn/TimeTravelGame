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

type UnitController struct {
	Units []*Unit

	moveUnits    []*Unit
	supportUnits []*Unit
}

func NewUnitController() *UnitController {
	u := &UnitController{
		Units: []*Unit{},
	}
	u.makeReady()
	return u
}
func (u *UnitController) makeReady() {
	u.moveUnits = []*Unit{}
	u.supportUnits = []*Unit{}

	for _, unit := range u.Units {
		if unit.Action.Kind == actionMove {

			u.moveUnits = append(u.moveUnits, unit)

		} else if unit.Action.Kind == actionSupport {

			u.supportUnits = append(u.supportUnits, unit)
		}
	}
}

func (u *UnitController) GetUnitAtPos(fieldPos CardPos, pos AxialPos) (int, *Unit) {
	for i, unit := range u.Units {
		if unit.FieldPos == fieldPos && unit.Pos == pos {
			return i, unit
		}
	}
	return -1, nil
}

func (u *UnitController) AddUnitAtTile(field *Field, tile *Tile, fraction *Fraction) *Unit {
	_, unit := u.GetUnitAtPos(field.Pos, tile.AxialPos)
	if unit == nil && tile.Visable {
		id := getFractionIndex(fraction)

		unit = NewUnit(field.Pos, tile.AxialPos, id)
		u.Units = append(u.Units, unit)
	}
	return unit
}

func (u *UnitController) RemoveUnitAtTile(field *Field, tile *Tile) {
	u.RemoveUnitAtPos(field.Pos, tile.AxialPos)
}
func (u *UnitController) RemoveUnitAtPos(fieldPos CardPos, pos AxialPos) {
	i, unit := u.GetUnitAtPos(fieldPos, pos)
	if unit != nil {
		u.Units = append(u.Units[:i], u.Units[i+1:]...)
	}
}

func (u *UnitController) draw(img *ebiten.Image, f *Field) {
	for _, unit := range u.Units {
		if unit.FieldPos == f.Pos {
			unit.draw(img, &Fractions[unit.FactionId], f)
		}
	}

	for _, unit := range u.Units {
		if unit.FieldPos == f.Pos && (unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport) {
			tile := f.GetAxial(unit.Pos)
			totile := f.GetAxial(unit.Action.ToPos)
			DrawArrow(tile.Pos, totile.Pos, img, &Fractions[unit.FactionId])
		}
	}
}

func (u *UnitController) CopyField(fromField *Field, toField *Field) {
	for _, unit := range u.Units {
		if unit.FieldPos == fromField.Pos {
			copyUnit := unit.copyToField(toField)
			u.Units = append(u.Units, copyUnit)
		}
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
	op.GeoM.Translate(tile.Pos.X-float64(w)/2, tile.Pos.Y-float64(h)/2)
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
