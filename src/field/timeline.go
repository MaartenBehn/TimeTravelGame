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

	Units []*Unit

	moveUnits    []*Unit
	supportUnits []*Unit

	S *Selector

	image *ebiten.Image
}

func NewTimeline(fieldSize int) *Timeline {
	bounds := AxialPos{Q: float64(fieldSize), R: float64(fieldSize)}.MulFloat(tileSize * 2).ToCard()
	bounds = bounds.Add(CardPos{X: tileSize, Y: tileSize})

	timeline := &Timeline{
		FieldSize:    fieldSize,
		FieldBounds:  bounds,
		Fields:       map[CardPos]*Field{},
		ActiveFields: []CardPos{},

		Units: []*Unit{},

		S: NewSelector(),
	}

	timeline.makeReady()

	return timeline
}

func (t *Timeline) makeReady() {
	for _, field := range t.Fields {
		field.makeReady()
	}
	t.makeReadyUnits()
	t.createImage()
}

func (t *Timeline) createImage() {
	size := CardPos{X: 10, Y: 10}
	for pos := range t.Fields {
		newSize := pos.Add(t.FieldBounds)
		if newSize.X >= size.X {
			size.X = newSize.X
		}
		if newSize.Y >= size.Y {
			size.Y = newSize.Y
		}
	}

	if t.image == nil {
		t.image = ebiten.NewImage(int(size.X), int(size.Y))
		return
	}

	w, h := t.image.Size()
	if w != int(size.X) || h != int(size.Y) {
		t.image = ebiten.NewImage(int(size.X), int(size.Y))
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

	t.makeReady()

	return f
}

func (t *Timeline) CopyField(toPos CardPos, fromField *Field) *Field {
	copiedField := *fromField
	copiedField.Pos = toPos
	copiedField.makeReady()
	copiedField.Update()

	t.Fields[toPos] = &copiedField

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
	t.image.Clear()

	for _, field := range t.Fields {
		field.Update()
		field.Draw(t.image)

		for _, unit := range t.Units {
			if unit.FieldPos == field.Pos {
				unit.draw(t.image, &Fractions[unit.FactionId])
			}
		}
	}

	for _, field := range t.Fields {
		for _, unit := range t.Units {
			if unit.FieldPos == field.Pos && (unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport) {
				tile := field.GetAxial(unit.TilePos)

				toField := t.Fields[unit.Action.FieldPos]
				totile := toField.GetAxial(unit.Action.TilePos)
				DrawArrow(field.Pos.Add(tile.CalcPos()), toField.Pos.Add(totile.CalcPos()), t.image, &Fractions[unit.FactionId])
			}
		}
	}
}

func (t *Timeline) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()
	img.DrawImage(t.image, op)

	t.S.Draw(img, cam, t.Fields[t.S.FieldPos])
}

func (t *Timeline) SubmitRound() {

	/*
		for i, pos := range t.ActiveFields {
			field := t.Fields[pos]
			newPos := pos.Add(CardPos{X: t.FieldBounds.X})
			t.CopyField(newPos, field)

			t.ActiveFields = append(t.ActiveFields, newPos)
			t.ActiveFields = append(t.ActiveFields[:i], t.ActiveFields[i+1:]...)
		}
	*/

	t.makeReady()

	t.SubmitRoundUnits()

	t.Update()

}
