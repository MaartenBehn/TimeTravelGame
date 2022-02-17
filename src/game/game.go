package game

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/field"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

func Init() {
	event.On(event.EventGameLoad, load)
}

type game struct {
	t    *field.Timeline
	cam  *util.Camera
	done bool
}

func load(data interface{}) {
	g := &game{
		t:    field.Load(data.(string)),
		done: false,
	}

	users = []user{
		NewPlayer(0, g.t),
		NewPlayer(1, g.t),
	}

	event.On(event.EventGameCurrentUser, func(data interface{}) {
		aktiveUser = data.(int)
	})
	event.Go(event.EventGameCurrentUser, 0)

	updateId := event.On(event.EventUpdate, func(data interface{}) {
		if g.done {
			return
		}

		g.update()
	})

	drawId := event.On(event.EventDraw, func(data interface{}) {
		g.draw(data.(*ebiten.Image))
	})

	uiSubmitId := event.On(event.EventGameUISubmit, func(data interface{}) {
		if users[aktiveUser].isPlayer() {
			event.Go(event.EventGameSubmitUser, aktiveUser)
		}
	})

	userSubmitId := event.On(event.EventGameSubmitUser, func(data interface{}) {
		if g.done {
			return
		}

		if data.(int) == aktiveUser {

			if aktiveUser >= len(users)-1 {
				g.t.SubmitRound()
				g.t.Update()

				var won = -1
				for i, u := range users {
					u.evaluate()

					if won == -1 && u.getScore() > 0 {
						won = i
					} else if won >= 0 && u.getScore() > 0 {
						won = -2
						break
					}
				}

				if won >= 0 {
					event.Go(event.EventGameWon, won)
					g.done = true
				}

				event.Go(event.EventGameCurrentUser, 0)
			} else {
				event.Go(event.EventGameCurrentUser, aktiveUser+1)
			}
		}
	})

	var unloadId event.ReciverId
	event.On(event.EventGameUnload, func(data interface{}) {
		event.UnOn(event.EventUpdate, updateId)
		event.UnOn(event.EventDraw, drawId)

		event.UnOn(event.EventGameUISubmit, uiSubmitId)
		event.UnOn(event.EventGameSubmitUser, userSubmitId)

		event.UnOn(event.EventGameUnload, unloadId)

		event.Go(event.EventUIShowPanel, ui.PageStart)
	})

	event.Go(event.EventUIShowPanel, ui.PageGame)
}

var users []user
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
