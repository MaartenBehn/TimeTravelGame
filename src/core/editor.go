package core

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

func init() {
	event.On(event.EventEditorNewMap, func(data interface{}) {
		newMap(data.(CardPos))
	})
	event.On(event.EventEditorSaveMap, func(data interface{}) {
		saveMap(data.(string))
	})
	event.On(event.EventEditorLoadMap, func(data interface{}) {
		loadMap(data.(string))
	})
}

func newMap(size CardPos) {
	g.m = NewMap(size)
}

func saveMap(name string) {
	if g.m == nil {
		return
	}
	saveMapBufferToFile(name, g.m.Save())
}

func loadMap(name string) {
	buffer := loadMapBufferFromFile(name)
	if buffer == nil {
		return
	}
	m := LoadMap(buffer)
	g.m = m
}
