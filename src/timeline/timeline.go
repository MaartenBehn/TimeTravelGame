package timeline

import (
	"github.com/Stroby241/TimeTravelGame/src/field"
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

type Timeline struct {
	fields map[CardPos]*field.Field
}

func NewTimeline(startField *field.Field) *Timeline {

	timeline := &Timeline{
		fields: map[CardPos]*field.Field{},
	}

	return timeline
}
