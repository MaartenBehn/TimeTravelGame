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

	fieldY float64
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

		S: NewSelector(bounds),
	}

	timeline.MakeReadyUI()
	timeline.Update()

	return timeline
}

func (t *Timeline) MakeReadyUI() {
	for _, field := range t.Fields {
		field.makeReadyUI()
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

	f.makeReadyUI()

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

	copiedField.makeReadyUI()

	t.Fields[toPos] = &copiedField

	return &copiedField
}

func (t *Timeline) Get(pos CardPos) (*Tile, *Field) {
	var field *Field
	for _, f := range t.Fields {

		fieldPos := f.Pos.Mul(f.Bounds)
		if pos.X >= fieldPos.X && pos.X < fieldPos.X+f.Bounds.X &&
			pos.Y >= fieldPos.Y && pos.Y < fieldPos.Y+f.Bounds.Y {
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
	if util.Debug {
		fmt.Println("Timeline Update")
	}

	//Update Image
	size := CardPos{X: 10, Y: 10}
	for pos := range t.Fields {
		newSize := pos.Mul(t.FieldBounds).Add(t.FieldBounds)
		if newSize.X >= size.X {
			size.X = newSize.X
		}
		if newSize.Y >= size.Y {
			size.Y = newSize.Y
		}
	}

	if t.image == nil {
		t.image = ebiten.NewImage(int(size.X), int(size.Y))
	} else if w, h := t.image.Size(); w != int(size.X) || h != int(size.Y) {
		t.image = ebiten.NewImage(int(size.X), int(size.Y))
	} else {
		t.image.Clear()
	}

	// Draw Field
	for _, field := range t.Fields {
		field.Update()
		field.Draw(t.image)
	}

	// Draw Units
	for _, unit := range t.Units {
		unit.draw(t.image, &Fractions[unit.FactionId])
	}

	_, selectedUnit := t.GetUnitAtPos(t.S.TimePos)

	// Draw Arrows
	for _, unit := range t.Units {
		if unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport {
			DrawArrow(unit.CalcPos(), unit.Action.CalcPos(), t.image, &Fractions[unit.FactionId], selectedUnit == unit)
		}
	}
}

func (t *Timeline) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()
	img.DrawImage(t.image, op)
	t.S.Draw(img, cam)

	// Debug
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

		_, selectedUnit := t.GetUnitAtPos(t.S.TimePos)
		if selectedUnit != nil {
			moves := selectedUnit.GetPositions(selectedUnit.TimePos, t)
			for _, move := range moves {
				txt := fmt.Sprintf("\n\nMovable")
				x, y := cam.GetMatrix().Apply(move.CalcPos().X, move.CalcPos().Y)
				text.Draw(img, txt, debugFont, int(x), int(y), colornames.Green)
			}
		}
	}
}
