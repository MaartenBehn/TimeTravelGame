package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventGameLoad, load)
}

type game struct {
	t   *field.Timeline
	cam *util.Camera
}

func load(data interface{}) {
	g := &game{
		t:   nil,
		cam: util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10}),
	}

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		g.update()
	})

	drawId := event.On(event.EventDraw, func(data interface{}) {
		g.draw(data.(*ebiten.Image))
	})

	loadMapId := event.On(event.EventGameUILoadMap, func(data interface{}) {
		g.t = field.LoadTimeline(data.(string))

		users = []user{
			NewPlayer(0, 0, g.t, util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10})),
			//NewPlayer(1, 1, g.t, util.NewCamera(CardPos{0, 0}, CardPos{500, 500}, CardPos{1, 1}, CardPos{10, 10})),
			NewBasicAI(1, 1, g.t, g.cam),
		}

		playerIds = []int{0}
		aktiveUser = 0
	})

	uiSubmitId := event.On(event.EventGameUISubmitRound, func(data interface{}) {
		if users[aktiveUser].isPlayer() {
			event.Go(event.EventGameSubmitUser, aktiveUser)
		}
	})

	userSubmitId := event.On(event.EventGameSubmitUser, func(data interface{}) {
		if data.(int) == aktiveUser {

			aktiveUser++
			if aktiveUser >= len(users) {
				g.t.SubmitRound()
				aktiveUser = 0
			}
		}
	})

	var unloadId event.ReciverId
	event.On(event.EventGameUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)
		event.UnOn(event.EventGameUILoadMap, loadMapId)

		event.UnOn(event.EventGameUISubmitRound, uiSubmitId)
		event.UnOn(event.EventGameSubmitUser, userSubmitId)

		event.UnOn(event.EventGameUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageGame)
}

var users []user
var playerIds []int
var aktiveUser int

func (g *game) update() {
	if users == nil || users[aktiveUser] == nil {
		return
	}

	users[aktiveUser].update()
}

func (g *game) draw(screen *ebiten.Image) {
	if users == nil || users[aktiveUser] == nil {
		return
	}

	users[aktiveUser].draw(screen)
}
