package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type basicAI userData

func NewBasicAI(id int, factionId int, t *field.Timeline, cam *util.Camera) *basicAI {
	return &basicAI{
		id:        id,
		factionId: factionId,
		t:         t,
		cam:       cam,
	}
}

func (ai *basicAI) isPlayer() bool { return false }

func (ai *basicAI) update() {

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

func (ai *basicAI) draw(screen *ebiten.Image) {
	if ai.t == nil {
		return
	}

	ai.t.Draw(screen, ai.cam)
}
