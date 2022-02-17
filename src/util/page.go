package util

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/hajimehoshi/ebiten/v2"
)

type Page struct {
	loaded bool

	loadEventId   event.EventId
	unloadEventId event.EventId

	loadId   event.ReciverId
	unloadId event.ReciverId
	updateId event.ReciverId
	drawId   event.ReciverId
}

func (p *Page) Init(loadEventId event.EventId, unloadEventId event.EventId) {
	p.loadEventId = loadEventId
	p.unloadEventId = unloadEventId

	p.loadId = event.On(p.loadEventId, p.Load)
}

func (p *Page) Load(data interface{}) {
	event.UnOn(p.loadEventId, p.loadId)

	p.updateId = event.On(event.EventUpdate, func(data interface{}) {
		p.Update()
	})
	p.drawId = event.On(event.EventDraw, func(data interface{}) {
		p.Draw(data.(*ebiten.Image))
	})

	p.unloadId = event.On(p.unloadEventId, p.Unload)
}

func (p *Page) Unload(data interface{}) {
	event.UnOn(p.unloadEventId, p.unloadId)

	event.UnOn(event.EventUpdate, p.updateId)
	event.UnOn(event.EventDraw, p.drawId)

	p.loadId = event.On(p.loadEventId, p.Load)
}

func (p *Page) Update() {

}

func (p *Page) Draw(screen *ebiten.Image) {

}
