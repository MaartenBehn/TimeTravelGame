package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

const fieldPadding = 30

type Timeline struct {
	FieldSize   int
	FieldBounds CardPos

	Fields       map[CardPos]*Field
	ActiveFields []CardPos

	U *UnitController
	S *Selector
}

func NewTimeline(fieldSize int) *Timeline {
	bounds := AxialPos{Q: float64(fieldSize), R: float64(fieldSize)}.MulFloat(tileSize * 2).ToCard()
	bounds = bounds.Add(CardPos{X: tileSize, Y: tileSize})

	timeline := &Timeline{
		FieldSize:    fieldSize,
		FieldBounds:  bounds,
		Fields:       map[CardPos]*Field{},
		ActiveFields: []CardPos{},
		U:            NewUnitController(),
		S:            NewSelector(),
	}

	return timeline
}

func (t *Timeline) makeReady() {
	for _, field := range t.Fields {
		field.makeReady()
	}
}

func (t *Timeline) AddField(pos CardPos) *Field {
	if t.Fields[pos] != nil {
		return t.Fields[pos]
	}

	f := NewField(t.FieldSize, t.FieldBounds)
	f.Pos = pos
	t.Fields[pos] = f
	t.ActiveFields = append(t.ActiveFields, pos)

	return f
}

func (t *Timeline) CopyField(pos CardPos, fromField *Field) *Field {
	copiedField := *fromField
	copiedField.Pos = pos
	copiedField.makeReady()
	copiedField.Update()
	t.Fields[pos] = &copiedField
	return &copiedField
}

func (t *Timeline) Get(pos CardPos) (*Tile, *Field) {
	var field *Field
	for _, f := range t.Fields {
		if pos.X >= f.Pos.X && pos.X < f.Pos.X+f.Bounds.X &&
			pos.Y >= f.Pos.Y && pos.Y < f.Pos.Y+f.Bounds.Y {
			field = f
			break
		}
	}
	if field == nil {
		return nil, nil
	}

	tile := field.GetCard(pos)
	return tile, field
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
	}

	t.S.Draw(img, cam, t.Fields[t.S.FieldPos])
}

func (t *Timeline) SubmitRound() {

	for i, pos := range t.ActiveFields {
		field := t.Fields[pos]
		newPos := pos.Add(CardPos{X: t.FieldBounds.X})
		newField := t.CopyField(newPos, field)
		t.U.CopyField(field, newField)
		t.U.makeReady()
		t.Update()

		t.ActiveFields = append(t.ActiveFields, newPos)
		t.ActiveFields = append(t.ActiveFields[:i], t.ActiveFields[i+1:]...)
	}

	t.U.SubmitRound(t.ActiveFields)
}
