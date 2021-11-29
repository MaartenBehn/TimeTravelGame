package core

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

type Unit struct {
	tile *Tile
	Pos  AxialPos
}

type Faction struct {
	Units []*Unit
}

type UnitController struct {
	Factions []*Faction
}

func NewFaction() *Faction {
	return &Faction{
		Units: []*Unit{},
	}
}

func NewUnitController(factionAmmount int) *UnitController {
	u := &UnitController{
		Factions: make([]*Faction, factionAmmount),
	}
	for i := range u.Factions {
		u.Factions[i] = NewFaction()
	}
	return u
}

func (u *Unit) Update() {

}

func (u *UnitController) Update() {
	for _, faction := range u.Factions {
		for _, unit := range faction.Units {
			unit.Update()
		}
	}
}
