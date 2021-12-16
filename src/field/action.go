package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"math"
)

const (
	actionStay    = 1
	actionMove    = 2
	actionSupport = 3
)

type Action struct {
	Kind       int
	ToFieldPos CardPos
	ToPos      AxialPos
	Support    int // For Move Action
}

func NewAction() *Action {
	return &Action{
		Kind: actionStay,
	}
}

func (u *UnitController) SetAction(unit *Unit, fieldPos CardPos, pos AxialPos) {

	if unit.Action.Kind == actionMove {
		for i := len(u.moveUnits) - 1; i >= 0; i-- {
			if u.moveUnits[i] == unit {
				u.moveUnits = append(u.moveUnits[:i], u.moveUnits[i+1:]...)
			}
		}
	} else if unit.Action.Kind == actionSupport {

		if _, actionUnit := u.GetUnitAtPos(unit.Action.ToFieldPos, unit.Action.ToPos); actionUnit != nil && actionUnit.FactionId == unit.FactionId {
			actionUnit.Support--
		}

		for _, actionUnit := range u.moveUnits {
			if actionUnit.Action.ToPos == unit.Action.ToPos && actionUnit.FactionId == unit.FactionId {
				actionUnit.Action.Support--
			}
		}

		for i := len(u.supportUnits) - 1; i >= 0; i-- {
			if u.supportUnits[i] == unit {
				u.supportUnits = append(u.supportUnits[:i], u.supportUnits[i+1:]...)
			}
		}
	}

	// If pos is the same -> Stay
	if unit.FieldPos == fieldPos && unit.Pos == pos {
		unit.Action.Kind = actionStay
		unit.Action.ToFieldPos = CardPos{}
		unit.Action.ToPos = AxialPos{}
		return
	}

	// If is to an own Unit -> Support
	if _, actionUnit := u.GetUnitAtPos(fieldPos, pos); actionUnit != nil && actionUnit.FactionId == unit.FactionId {
		unit.Action.Kind = actionSupport
		unit.Action.ToFieldPos = fieldPos
		unit.Action.ToPos = pos
		unit.Action.Support = 0

		actionUnit.Support++
		u.supportUnits = append(u.supportUnits, unit)
		return
	}

	// If is to an own Move -> Support
	for _, actionUnit := range u.moveUnits {
		if actionUnit.Action.ToFieldPos == fieldPos && actionUnit.Action.ToPos == pos &&
			actionUnit.FactionId == unit.FactionId {

			unit.Action.Kind = actionSupport
			unit.Action.ToFieldPos = fieldPos
			unit.Action.ToPos = pos
			unit.Action.Support = 0

			actionUnit.Action.Support++
			u.supportUnits = append(u.supportUnits, unit)
			return
		}
	}

	// Else -> Move
	unit.Action.Kind = actionMove
	unit.Action.ToFieldPos = fieldPos
	unit.Action.ToPos = pos
	u.moveUnits = append(u.moveUnits, unit)
}

type targetPos struct {
	fieldPos  CardPos
	pos       AxialPos
	moveUnits []*Unit

	loopUnit *Unit
}

func (u *UnitController) SubmitRound(aktiveFields []CardPos) {

	targetPositons := []*targetPos{}

	for _, unit := range u.moveUnits {

		isAktive := false
		for _, pos := range aktiveFields {
			if unit.FieldPos == pos {
				isAktive = true
			}
		}
		if !isAktive {
			continue
		}

		found := false
		for i, positon := range targetPositons {
			if positon.fieldPos == unit.Action.ToFieldPos && positon.pos == unit.Action.ToPos {
				targetPositons[i].moveUnits = append(targetPositons[i].moveUnits, unit)
				found = true
			}
		}

		if !found {
			position := &targetPos{
				fieldPos:  unit.Action.ToFieldPos,
				pos:       unit.Action.ToPos,
				moveUnits: []*Unit{unit},
			}
			targetPositons = append(targetPositons, position)
		}
	}

	moveUnit := func(unit *Unit, fieldPos CardPos, pos AxialPos) {
		for j := len(u.supportUnits) - 1; j >= 0; j-- {
			if u.supportUnits[j].Action.ToFieldPos == unit.FieldPos && u.supportUnits[j].Action.ToPos == unit.Pos {

				u.SetAction(u.supportUnits[j], u.supportUnits[j].FieldPos, u.supportUnits[j].Pos)

			} else if u.supportUnits[j].Action.ToFieldPos == unit.Action.ToFieldPos && u.supportUnits[j].Action.ToPos == unit.Action.ToPos {

				u.SetAction(u.supportUnits[j], u.supportUnits[j].FieldPos, u.supportUnits[j].Pos)
			}
		}

		unit.FieldPos = fieldPos
		unit.Pos = pos

		u.SetAction(unit, unit.FieldPos, unit.Pos)
	}

	for len(targetPositons) > 0 {

		madeMove := false

		for i, positon := range targetPositons {
			positon.loopUnit = nil

			var winningUnit *Unit
			winningSupport := 0

			var loopWinningUnit *Unit
			loopWinningSupport := math.MaxInt32

			_, presentUnit := u.GetUnitAtPos(positon.fieldPos, positon.pos)
			if presentUnit != nil {
				winningSupport = presentUnit.Support + 1
				loopWinningSupport = presentUnit.Support
			}

			for _, unit := range positon.moveUnits {

				if (unit.Action.Support + 1) > winningSupport {
					winningUnit = unit
					winningSupport = unit.Action.Support + 1
				} else if (unit.Action.Support + 1) == winningSupport {
					winningUnit = nil
				}

				if (unit.Action.Support + 1) > loopWinningSupport {
					loopWinningUnit = unit
					loopWinningSupport = unit.Action.Support + 1
				} else if (unit.Action.Support + 1) == loopWinningSupport {
					loopWinningUnit = nil
				}
			}

			if winningUnit != nil {
				if presentUnit != nil {
					u.RemoveUnitAtPos(positon.fieldPos, positon.pos)
				}

				moveUnit(winningUnit, positon.fieldPos, positon.pos)

				targetPositons = append(targetPositons[:i], targetPositons[i+1:]...)

				madeMove = true
				break

			} else if loopWinningUnit != nil && presentUnit != nil {
				positon.loopUnit = loopWinningUnit
			}
		}

		for _, positon := range targetPositons {
			if positon.loopUnit != nil {

				loop := []*targetPos{positon}

				findingPos := true
				loopDone := false
				for findingPos {
					findingPos = false

					for _, testPosition := range targetPositons {
						if positon.loopUnit != nil && testPosition != positon &&
							loop[len(loop)-1].loopUnit.Pos == testPosition.pos {

							loop = append(loop, testPosition)
							findingPos = true

							if testPosition.loopUnit.Pos == loop[0].pos {
								loopDone = true
							}

							break
						}
					}
				}

				if loopDone {
					for _, pos := range loop {
						moveUnit(pos.loopUnit, pos.fieldPos, pos.pos)

						for i, targetPositon := range targetPositons {
							if targetPositon == pos {
								targetPositons = append(targetPositons[:i], targetPositons[i+1:]...)
								break
							}
						}
					}
					madeMove = true
					break
				}
			}
		}

		if !madeMove {
			break
		}
	}

}
