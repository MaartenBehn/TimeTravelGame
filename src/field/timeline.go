package field

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
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
	size := CardPos{X: 1000, Y: 1000}
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

	f := NewField(t.FieldSize, t.FieldBounds, pos)
	t.Fields[pos] = f

	f.Active = true
	t.ActiveFields = append(t.ActiveFields, pos)

	t.makeReady()

	return f
}

func (t *Timeline) CopyField(toPos CardPos, fromField *Field) *Field {
	field := *fromField
	copiedField := field
	copiedField.Pos = toPos

	copiedField.Tiles = make([]Tile, len(field.Tiles))
	for i, tile := range field.Tiles {
		tile.FieldPos = toPos
		copiedField.Tiles[i] = tile
	}

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

	// Draw Field
	for _, field := range t.Fields {
		field.Update()
		field.Draw(t.image)
	}

	// Draw Units
	for _, unit := range t.Units {
		unit.draw(t.image, &Fractions[unit.FactionId])
	}

	// Draw Arrows
	for _, unit := range t.Units {
		if unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport {
			DrawArrow(unit.CalcPos(), unit.Action.CalcPos(), t.image, &Fractions[unit.FactionId])
		}
	}
}

func (t *Timeline) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()
	img.DrawImage(t.image, op)

	t.S.Draw(img, cam)

	if util.Debug {
		for _, field := range t.Fields {
			for _, tile := range field.Tiles {
				if !tile.Visable {
					continue
				}

				x, y := cam.GetMatrix().Apply(tile.CalcPos().X, tile.CalcPos().Y)
				text.Draw(img, tile.TimePos.ToString(), debugFont, int(x), int(y), colornames.Green)
			}
		}

		for _, unit := range t.Units {
			txt := fmt.Sprintf("\nKind: %d", unit.Action.Kind)
			x, y := cam.GetMatrix().Apply(unit.CalcPos().X, unit.CalcPos().Y)
			text.Draw(img, txt, debugFont, int(x), int(y), colornames.Green)
		}
	}
}
