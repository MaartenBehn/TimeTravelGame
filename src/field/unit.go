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
	Units [][]*Unit

	moveUnits    []*Unit
	supportUnits []*Unit
}

func NewUnitController() *UnitController {
	u := &UnitController{
		Units: make([][]*Unit, len(Fractions)),
	}
	u.makeReady()
	return u
}
func (u *UnitController) makeReady() {
	for _, units := range u.Units {
		for _, unit := range units {
			if unit.Action.Kind == actionMove {

				u.moveUnits = append(u.moveUnits, unit)

			} else if unit.Action.Kind == actionSupport {

				u.supportUnits = append(u.supportUnits, unit)
			}
		}
	}
}

func (u *UnitController) GetUnitAtPos(pos AxialPos) (*Fraction, int, *Unit) {
	for f, units := range u.Units {
		for i, unit := range units {
			if unit.Pos == pos {
				return &Fractions[f], i, unit
			}
		}
	}
	return nil, -1, nil
}

func (u *UnitController) AddUnitAtTile(tile *Tile, fraction *Fraction) *Unit {
	_, _, unit := u.GetUnitAtPos(tile.AxialPos)
	if unit == nil && tile.Visable {
		id := getFractionIndex(fraction)

		unit = NewUnit(tile.AxialPos, id)
		u.Units[id] = append(u.Units[id], unit)
	}
	return unit
}
func (u *UnitController) AddUnitAtPos(pos AxialPos, fraction *Fraction, f *Field) *Unit {
	tile := f.GetAxial(pos)
	if tile == nil {
		return nil
	}
	return u.AddUnitAtTile(tile, fraction)
}

func (u *UnitController) RemoveUnitAtTile(tile *Tile) {
	u.RemoveUnitAtPos(tile.AxialPos)
}
func (u *UnitController) RemoveUnitAtPos(pos AxialPos) {
	f, i, unit := u.GetUnitAtPos(pos)
	if unit != nil {
		j := getFractionIndex(f)
		u.Units[j] = append(u.Units[j][:i], u.Units[j][i+1:]...)
	}
}

func (u *UnitController) draw(img *ebiten.Image, f *Field) {
	for i, units := range u.Units {
		for _, unit := range units {
			unit.draw(img, &Fractions[i], f)
		}
	}

	for _, units := range u.Units {
		for _, unit := range units {
			if unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport {
				tile := f.GetAxial(unit.Pos)
				totile := f.GetAxial(*unit.Action.ToPos)
				DrawArrow(tile.Pos, totile.Pos, img, &Fractions[unit.FactionId])
			}
		}
	}
}

type Unit struct {
	FactionId int
	Pos       AxialPos
	Action    *Action
	Support   int
}

func NewUnit(pos AxialPos, factionId int) *Unit {
	return &Unit{
		FactionId: factionId,
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
