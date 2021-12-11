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
	moveUnits    []*Unit
	supportUnits []*Unit
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

	for _, units := range u.Units {
		for _, unit := range units {
			if unit.Action.Kind == actionMove {

				u.moveUnits = append(u.moveUnits, unit)

			} else if unit.Action.Kind == actionSupport {

				u.supportUnits = append(u.supportUnits, unit)

				if _, _, actionUnit := u.GetUnitAtPos(*unit.Action.ToPos); actionUnit != nil && actionUnit.FactionId == unit.FactionId {
					actionUnit.supportUnits++
				}

				for _, actionUnit := range u.moveUnits {
					if *actionUnit.Action.ToPos == *unit.Action.ToPos && actionUnit.FactionId == unit.FactionId {
						actionUnit.Action.supportUnits++
						break
					}
				}
			}
		}
	}
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
	u.RemoveUnitAtPos(tile.AxialPos)
}
func (u *UnitController) RemoveUnitAtPos(pos AxialPos) {
	f, i, unit := u.GetUnitAtPos(pos)
	if unit != nil {
		j := u.getFractionIndex(f)
		u.Units[j] = append(u.Units[j][:i], u.Units[j][i+1:]...)
	}
}

func (u *UnitController) SetSelector(pos AxialPos) {
	_, _, unit := u.GetUnitAtPos(pos)
	if unit != nil {
		u.SelectedUnit = unit.Pos
	}
}

type Unit struct {
	FactionId    int
	Pos          AxialPos
	Action       *Action
	supportUnits int
}

func NewUnit(pos AxialPos, factionId int) *Unit {
	return &Unit{
		FactionId: factionId,
		Pos:       pos,
		Action:    NewAction(),
	}
}
