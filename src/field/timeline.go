package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Timeline struct {
	Fields map[CardPos]*Field

	U *UnitController
	S *Selector
}

func NewTimeline() *Timeline {

	timeline := &Timeline{
		Fields: map[CardPos]*Field{},
		U:      NewUnitController(),
		S:      NewSelector(),
	}

	return timeline
}

func (t *Timeline) Update() {
	for _, field := range t.Fields {
		field.Update()
		t.U.draw(field.image, field)
	}
}

func (t *Timeline) Draw(img *ebiten.Image, cam *util.Camera) {
	for _, field := range t.Fields {
		field.Draw(img, cam)
		t.S.Draw(img, cam, field)
	}
}
