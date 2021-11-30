package event

const EventUpdate EventId = 1
const EventCamUpdate EventId = 2

const EventEditorNewMap EventId = 10
const EventEditorSaveMap EventId = 11

const eventMax = 255

type EventId int
type reciverId int

type event struct {
	id       EventId
	receiver []func(data interface{})
}

var events [eventMax]event

func init() {
	for _, e := range events {
		e.receiver = []func(interface{}){}
	}
}

func Go(id EventId, data interface{}) {
	for _, r := range events[id].receiver {
		r(data)
	}
}

func On(id EventId, f func(data interface{})) reciverId {
	for i, r := range events[id].receiver {
		if r == nil {
			events[id].receiver[i] = f
			return (reciverId)(i)
		}
	}

	events[id].receiver = append(events[id].receiver, f)
	return (reciverId)(len(events[id].receiver) - 1)
}

func UnOn(id EventId, rId reciverId) {
	if (reciverId)(len(events[id].receiver)) <= rId {
		return
	}
	events[id].receiver[id] = nil
}
