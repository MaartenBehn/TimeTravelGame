package event

const EventUpdate EventId = 1
const EventDraw EventId = 2

const EventEditorLoad EventId = 10
const EventEditorUnload EventId = 11
const EventEditorNewMap EventId = 12
const EventEditorSaveMap EventId = 13
const EventEditorLoadMap EventId = 14

const EventUIEditorShow EventId = 20
const EventUIGameShow EventId = 21

const eventMax = 255

type EventId int
type ReciverId int

type event struct {
	id       EventId
	receiver []func(data interface{})
}

var events [eventMax]event

func Init() {
	for _, e := range events {
		e.receiver = []func(interface{}){}
	}
}

func Go(id EventId, data interface{}) {
	for _, r := range events[id].receiver {
		r(data)
	}
}

func On(id EventId, f func(data interface{})) ReciverId {
	for i, r := range events[id].receiver {
		if r == nil {
			events[id].receiver[i] = f
			return (ReciverId)(i)
		}
	}

	events[id].receiver = append(events[id].receiver, f)
	return (ReciverId)(len(events[id].receiver) - 1)
}

func UnOn(id EventId, rId ReciverId) {
	if (ReciverId)(len(events[id].receiver)) <= rId {
		return
	}
	events[id].receiver[id] = nil
}
