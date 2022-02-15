package game

import (
	"github.com/Stroby241/TimeTravelGame/src/field"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type user interface {
	isPlayer() bool
	getScore() int
	evaluate()
	update()
	draw(screen *ebiten.Image)
}

type userData struct {
	player    bool
	factionId int
	t         *field.Timeline
	cam       *util.Camera
	score     int
}

func NewUserData(player bool, factionId int, t *field.Timeline, cam *util.Camera) userData {
	return userData{
		player:    player,
		factionId: factionId,
		t:         t,
		cam:       cam,
		score:     0,
	}
}

func (u *userData) isPlayer() bool { return u.player }
func (u *userData) getScore() int  { return u.score }
func (u *userData) draw(screen *ebiten.Image) {
	if u.t == nil {
		return
	}

	u.t.Draw(screen, u.cam)
}
func (u *userData) evaluate() {
	u.score = EvaluateFaction(u.t, u.factionId)
}

func EvaluateFaction(t *field.Timeline, factionId int) int {
	score := 0
	for _, unit := range t.Units {
		field := t.Fields[unit.FieldPos]
		if field != nil && field.Active && unit.FactionId == factionId {
			score++
		}
	}
	return score
}
