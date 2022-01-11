package game

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"math"
	"math/rand"
)

type AI0 struct {
	userData
}

func NewAI0(id int, factionId int, t *field.Timeline, cam *util.Camera) *AI0 {
	return &AI0{
		userData: NewUserData(false, id, factionId, t, cam),
	}
}
func (ai *AI0) update() {
	for _, unit := range ai.t.Units {
		if unit.FactionId == ai.factionId {
			isAktive := false
			for _, activeField := range ai.t.ActiveFields {
				if unit.FieldPos == activeField {
					isAktive = true
				}
			}

			if isAktive {
				done := false
				for !done {
					pos := unit.TimePos
					pos.TilePos = pos.TilePos.Add(AxialDirections[rand.Intn(len(AxialDirections)-1)])

					tile, _ := ai.t.Get(pos.CalcPos())
					if tile != nil && tile.Visable {
						ai.t.SetAction(unit, pos)
						done = true
					}
				}
			}
		}
	}

	event.Go(event.EventGameSubmitUser, ai.id)
}

type AI1 struct {
	userData
}

func NewAI1(id int, factionId int, t *field.Timeline, cam *util.Camera) *AI1 {
	return &AI1{
		userData: NewUserData(false, id, factionId, t, cam),
	}
}
func (ai *AI1) update() {
	for _, unit := range ai.t.Units {
		if unit.FactionId == ai.factionId {
			isAktive := false
			for _, activeField := range ai.t.ActiveFields {
				if unit.FieldPos == activeField {
					isAktive = true
				}
			}

			if isAktive {
				done := false
				for !done {
					pos := unit.TimePos
					pos.TilePos = pos.TilePos.Add(AxialDirections[rand.Intn(len(AxialDirections)-1)])
					pos.FieldPos = pos.FieldPos.Add(CardPos{
						X: float64(rand.Intn(2) - 1),
						Y: float64(rand.Intn(2) - 1),
					})

					tile, _ := ai.t.Get(pos.CalcPos())
					if tile != nil && tile.Visable {
						ai.t.SetAction(unit, pos)
						done = true
					}
				}
			}
		}
	}

	event.Go(event.EventGameSubmitUser, ai.id)
}

type AI2 struct {
	userData
}

func NewAI2(id int, factionId int, t *field.Timeline, cam *util.Camera) *AI2 {
	return &AI2{
		userData: NewUserData(false, id, factionId, t, cam),
	}
}

func (ai *AI2) update() {
	var units []*field.Unit

	for _, unit := range ai.t.Units {
		if unit.FactionId == ai.factionId {
			isAktive := false
			for _, activeField := range ai.t.ActiveFields {
				if unit.FieldPos == activeField {
					isAktive = true
				}
			}

			if isAktive {
				units = append(units, unit)
			}
		}
	}

	var possibleMoves = make([][]field.TimePos, len(units))
	var totalCombinations = 1
	var size = make([]int, len(units))
	var counter = make([]int, len(units))

	for i, unit := range units {
		moves := unit.MovePattern.GetPositions(unit.TimePos, ai.t)
		for _, move := range moves {
			tile, _ := ai.t.Get(move.CalcPos())

			if tile != nil && tile.Visable {
				possibleMoves[i] = append(possibleMoves[i], move)
			}
		}

		s := len(possibleMoves[i])
		size[i] = s
		totalCombinations *= s
	}

	var bestEvaluation = math.MinInt
	var bestCounter = make([]int, len(counter))
	for i := 0; i < totalCombinations; i++ {

		fmt.Printf("\rCalculating possible move %d of %d moves. Counter: ", i, totalCombinations)
		fmt.Print(counter)

		buffer := ai.t.ToBuffer()
		tCopy := field.FromBuffer(buffer)

		for i := 0; i < len(units); i++ {
			_, copyUnit := tCopy.GetUnitAtPos(units[i].TimePos)
			tCopy.SetAction(copyUnit, possibleMoves[i][counter[i]])
		}

		tCopy.SubmitRound()
		score := EvaluateFaction(tCopy, ai.factionId)
		score -= EvaluateFaction(tCopy, 0) * 2

		if score > bestEvaluation {
			for j, val := range counter {
				bestCounter[j] = val
			}
			bestEvaluation = score
		}

		for incIndex := len(possibleMoves) - 1; incIndex >= 0; incIndex-- {
			if counter[incIndex]+1 < size[incIndex] {
				counter[incIndex]++
				break
			}
			counter[incIndex] = 0
		}
	}

	for i, unit := range units {
		ai.t.SetAction(unit, possibleMoves[i][bestCounter[i]])
	}
	fmt.Printf(" done\n")

	event.Go(event.EventGameSubmitUser, ai.id)
}

type AI3 struct {
	userData
	tries int
}

func NewAI3(id int, factionId int, t *field.Timeline, cam *util.Camera, tries int) *AI3 {
	return &AI3{
		userData: NewUserData(false, id, factionId, t, cam),
		tries:    tries,
	}
}

func (ai *AI3) update() {
	var units []*field.Unit

	for _, unit := range ai.t.Units {
		if unit.FactionId == ai.factionId {
			isAktive := false
			for _, activeField := range ai.t.ActiveFields {
				if unit.FieldPos == activeField {
					isAktive = true
				}
			}

			if isAktive {
				units = append(units, unit)
			}
		}
	}

	var possibleMoves = make([][]field.TimePos, len(units))
	var totalCombinations = 1
	var size = make([]int, len(units))
	var counter = make([]int, len(units))

	for i, unit := range units {
		moves := unit.MovePattern.GetPositions(unit.TimePos, ai.t)
		for _, move := range moves {
			tile, _ := ai.t.Get(move.CalcPos())

			if tile != nil && tile.Visable {
				possibleMoves[i] = append(possibleMoves[i], move)
			}
		}

		s := len(possibleMoves[i])
		size[i] = s
		totalCombinations *= s
	}

	var bestEvaluation = math.MinInt
	var bestCounter = make([]int, len(counter))

	rounds := totalCombinations
	if rounds > ai.tries {
		rounds = ai.tries
	}
	for i := 0; i < rounds; i++ {

		fmt.Printf("\rCalculating possible move %d of %d moves. Counter: ", i, rounds)
		fmt.Print(counter)

		buffer := ai.t.ToBuffer()
		tCopy := field.FromBuffer(buffer)

		for i := 0; i < len(units); i++ {
			_, copyUnit := tCopy.GetUnitAtPos(units[i].TimePos)
			tCopy.SetAction(copyUnit, possibleMoves[i][counter[i]])
		}

		tCopy.SubmitRound()
		score := EvaluateFaction(tCopy, ai.factionId)
		score -= EvaluateFaction(tCopy, 0) * 5

		if score > bestEvaluation {
			for j, val := range counter {
				bestCounter[j] = val
			}
			bestEvaluation = score
		}

		if rounds == totalCombinations {
			for incIndex := len(possibleMoves) - 1; incIndex >= 0; incIndex-- {
				if counter[incIndex]+1 < size[incIndex] {
					counter[incIndex]++
					break
				}
				counter[incIndex] = 0
			}
		} else {
			for i, _ := range counter {
				counter[i] = rand.Intn(size[i] - 1)
			}
		}
	}

	for i, unit := range units {
		ai.t.SetAction(unit, possibleMoves[i][bestCounter[i]])
	}
	fmt.Printf(" done\n")

	event.Go(event.EventGameSubmitUser, ai.id)
}
