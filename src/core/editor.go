package core

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

func init() {
	event.On(event.EventCamUpdate, func(data interface{}) {
		update()
	})
	event.On(event.EventEditorNewMap, func(data interface{}) {
		newMap(data.(CardPos))
	})
	event.On(event.EventEditorSaveMap, func(data interface{}) {
		saveMap(data.(string))
	})
}

func update() {

}

func newMap(size CardPos) {
	g.m = NewMap(size)
}

func saveMap(name string) {
	saveMapBufferToFile(name, g.m.Save())
}
