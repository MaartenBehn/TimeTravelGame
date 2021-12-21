package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

const (
	actionStay    = 1
	actionMove    = 2
	actionSupport = 3
)

type Action struct {
	TimePos
	Kind    int
	Support int // For Move Action
}

func NewAction() Action {
	return Action{
		Kind: actionStay,
	}
}

func (t *Timeline) SetAction(unit *Unit, pos TimePos) {

	if unit.Action.Kind == actionMove {
		for i := len(t.moveUnits) - 1; i >= 0; i-- {
			if t.moveUnits[i] == unit {
				t.moveUnits = append(t.moveUnits[:i], t.moveUnits[i+1:]...)
			}
		}
	} else if unit.Action.Kind == actionSupport {

		if _, actionUnit := t.GetUnitAtPos(unit.Action.TimePos); actionUnit != nil && actionUnit.FactionId == unit.FactionId {
			actionUnit.Support--
		}

		for _, actionUnit := range t.moveUnits {
			if actionUnit.Action.TilePos == unit.Action.TilePos && actionUnit.FactionId == unit.FactionId {
				actionUnit.Action.Support--
			}
		}

		for i := len(t.supportUnits) - 1; i >= 0; i-- {
			if t.supportUnits[i] == unit {
				t.supportUnits = append(t.supportUnits[:i], t.supportUnits[i+1:]...)
			}
		}
	}

	// If TimePos is the same -> Stay
	if unit.SamePos(pos) {
		unit.Action.Kind = actionStay
		unit.Action.FieldPos = CardPos{}
		unit.Action.TilePos = AxialPos{}
		return
	}

	// If is to an own Unit -> Support
	if _, actionUnit := t.GetUnitAtPos(pos); actionUnit != nil && actionUnit.FactionId == unit.FactionId {
		unit.Action.Kind = actionSupport
		unit.Action.TimePos = pos
		unit.Action.Support = 0

		actionUnit.Support++
		t.supportUnits = append(t.supportUnits, unit)
		return
	}

	// If is to an own Move -> Support
	for _, actionUnit := range t.moveUnits {
		if actionUnit.Action.SamePos(pos) && actionUnit.FactionId == unit.FactionId {

			unit.Action.Kind = actionSupport
			unit.Action.TimePos = pos
			unit.Action.Support = 0

			actionUnit.Action.Support++
			t.supportUnits = append(t.supportUnits, unit)
			return
		}
	}

	// Else -> Move
	unit.Action.Kind = actionMove
	unit.Action.TimePos = pos
	t.moveUnits = append(t.moveUnits, unit)
}

type targetPos struct {
	TimePos
	moveUnits   []*Unit
	presentUnit *Unit
	winningUnit *Unit
}

var fieldY float64

func (t *Timeline) SubmitRound() {
	t.makeReady()

	var targetPositions []*targetPos
	for _, unit := range t.moveUnits {

		// Find all aktive Units
		isAktive := false
		for _, pos := range t.ActiveFields {
			if unit.FieldPos == pos {
				isAktive = true
			}
		}
		if !isAktive {
			continue
		}

		// Check if there is already a target Position
		found := false
		for i, positon := range targetPositions {
			if positon.SamePos(unit.Action.TimePos) {
				targetPositions[i].moveUnits = append(targetPositions[i].moveUnits, unit)
				found = true
			}
		}

		// If not add target Position
		if !found {
			position := &targetPos{
				TimePos:   unit.Action.TimePos,
				moveUnits: []*Unit{unit},
			}

			_, position.presentUnit = t.GetUnitAtPos(position.TimePos)
			if position.presentUnit != nil {
				position.moveUnits = append(position.moveUnits, position.presentUnit)
			}

			targetPositions = append(targetPositions, position)
		}
	}

	// Find Winning Units for target Positions
	changeHappend := true
	for changeHappend {
		changeHappend = false
		for _, position := range targetPositions {

			oldWinningUnit := position.winningUnit
			winningAmount := 0
			for _, unit := range position.moveUnits {

				amount := unit.Action.Support + 1
				if amount > winningAmount {
					position.winningUnit = unit
					winningAmount = amount
				} else if amount == winningAmount {
					position.winningUnit = nil
				}
			}

			if position.presentUnit != nil && position.winningUnit == nil {
				position.winningUnit = position.presentUnit
			}

			if position.winningUnit != oldWinningUnit {
				changeHappend = true
			}
		}
	}

	var newFieldPositions []CardPos
	var fieldPositions []CardPos

	// Copy winning Units
	for _, position := range targetPositions {
		if position.winningUnit == nil {
			continue
		}

		copyUnit := t.CopyUnit(position.winningUnit)

		isAktive := false
		for _, pos := range t.ActiveFields {
			if pos == position.FieldPos {
				isAktive = true
			}
		}
		if isAktive {

			copyUnit.FieldPos = position.FieldPos.Add(CardPos{X: t.FieldBounds.X})

		} else {
			var newFieldPos CardPos
			contained := false
			for i, pos := range fieldPositions {
				if pos == position.FieldPos {
					contained = true
					newFieldPos = newFieldPositions[i]
					break
				}
			}

			if !contained {
				fieldPositions = append(fieldPositions, position.FieldPos)

				fieldY += t.FieldBounds.Y
				newFieldPos = CardPos{X: position.FieldPos.X, Y: fieldY}
				newFieldPositions = append(newFieldPositions, newFieldPos)
			}

			copyUnit.FieldPos = newFieldPos
		}

		copyUnit.TilePos = position.TilePos
		copyUnit.Action = NewAction()
	}

	// Copy all not moved Units
	for _, unit := range t.Units {
		isAktive := false
		for _, pos := range t.ActiveFields {
			if unit.FieldPos == pos {
				isAktive = true
				break
			}
		}

		isWinning := false
		for _, position := range targetPositions {
			if unit == position.winningUnit {
				isWinning = true
				break
			}
		}

		newField := -1
		for i, pos := range fieldPositions {
			if pos == unit.FieldPos {
				newField = i
				break
			}
		}

		if isAktive && !isWinning {
			copyUnit := t.CopyUnit(unit)
			copyUnit.FieldPos = unit.FieldPos.Add(CardPos{X: t.FieldBounds.X})
			if copyUnit.Action.FieldPos == unit.FieldPos {
				copyUnit.Action.FieldPos = copyUnit.FieldPos
			}
		}

		if newField != -1 && !isAktive && !isWinning {
			copyUnit := t.CopyUnit(unit)
			copyUnit.FieldPos = newFieldPositions[newField]
			if copyUnit.Action.FieldPos == unit.FieldPos {
				copyUnit.Action.FieldPos = copyUnit.FieldPos
			}
		}
	}

	// Copy Aktive Fields
	for i := len(t.ActiveFields) - 1; i >= 0; i-- {
		field := t.Fields[t.ActiveFields[i]]
		newPos := t.ActiveFields[i].Add(CardPos{X: t.FieldBounds.X})
		copyField := t.CopyField(newPos, field)

		t.ActiveFields = append(t.ActiveFields, newPos)
		copyField.Active = true

		t.ActiveFields = append(t.ActiveFields[:i], t.ActiveFields[i+1:]...)
		field.Active = false
	}

	// Copy new Fields
	for i, pos := range fieldPositions {
		field := t.Fields[pos]
		newPos := newFieldPositions[i]
		copyField := t.CopyField(newPos, field)

		copyField.Active = true
		t.ActiveFields = append(t.ActiveFields, newPos)
	}

	t.Update()
}
