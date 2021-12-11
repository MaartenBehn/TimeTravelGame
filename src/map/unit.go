package gameMap

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"golang.org/x/image/colornames"
	"image/color"
)

type Fraction struct {
	name  string
	color color.Color
}

var (
	FractionRed = Fraction{
		name:  "red",
		color: colornames.Red,
	}
	FractionBlue = Fraction{
		name:  "blue",
		color: colornames.Blue,
	}
)

type UnitController struct {
	fractions []*Fraction
	Units     map[int][]*Unit

	SelectedUnit AxialPos
}

func NewUnitController() *UnitController {
	u := &UnitController{
		Units: map[int][]*Unit{},
	}
	u.makeReady()
	return u
}
func (u *UnitController) makeReady() {
	u.fractions = []*Fraction{&FractionRed, &FractionBlue}
}

func (u *UnitController) getFractionIndex(f *Fraction) int {
	for i, fraction := range u.fractions {
		if fraction == f {
			return i
		}
	}
	return -1
}

func (u *UnitController) GetUnitAtPos(pos AxialPos) (*Fraction, int, *Unit) {
	for f, units := range u.Units {
		for i, unit := range units {
			if unit.Pos == pos {
				return u.fractions[f], i, unit
			}
		}
	}
	return nil, -1, nil
}

func (u *UnitController) AddUnitAtTile(tile *Tile, fraction *Fraction) *Unit {
	_, _, unit := u.GetUnitAtPos(tile.AxialPos)
	if unit == nil && tile.Visable {
		id := u.getFractionIndex(fraction)

		unit = NewUnit(tile.AxialPos, id)
		u.Units[id] = append(u.Units[id], unit)
	}
	return unit
}
func (u *UnitController) AddUnitAtPos(pos AxialPos, fraction *Fraction, m *Map) *Unit {
	tile, _ := m.GetAxial(pos)
	if tile == nil {
		return nil
	}
	return u.AddUnitAtTile(tile, fraction)
}

func (u *UnitController) RemoveUnitAtTile(tile *Tile) {
	f, i, unit := u.GetUnitAtPos(tile.AxialPos)
	if unit != nil {
		j := u.getFractionIndex(f)
		u.Units[j] = append(u.Units[j][:i], u.Units[j][i+1:]...)
	}
}
func (u *UnitController) RemoveUnitAtPos(pos AxialPos, m *Map) {
	tile, _ := m.GetAxial(pos)
	if tile == nil {
		return
	}
	u.RemoveUnitAtTile(tile)
}

func (u *UnitController) SetSelector(pos AxialPos) {
	_, _, unit := u.GetUnitAtPos(pos)
	if unit != nil {
		u.SelectedUnit = unit.Pos
	}
}

type Unit struct {
	FactionId int
	Pos       AxialPos
	TargetPos AxialPos
}

func NewUnit(pos AxialPos, factionId int) *Unit {
	return &Unit{
		FactionId: factionId,
		Pos:       pos,
		TargetPos: pos,
	}
}
