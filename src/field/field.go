package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Field struct {
	Size   int
	Bounds CardPos
	Pos    CardPos
	Tiles  []Tile
	Active bool

	image *ebiten.Image
}

func NewField(size int, bounds CardPos, pos CardPos) *Field {
	field := &Field{
		Size:   size,
		Bounds: bounds,
		Pos:    pos,
		Tiles:  make([]Tile, size*size),
	}

	for i, _ := range field.Tiles {
		q, r := reverseIndex(i, size)
		field.Tiles[i] = NewTile(TimePos{
			TilePos:     AxialPos{Q: float64(q), R: float64(r)},
			FieldPos:    pos,
			FieldBounds: bounds,
		})
	}

	field.makeReady()
	field.Update()

	return field
}

func (f *Field) makeReady() {
	for i, tile := range f.Tiles {
		tile.makeReady()
		f.Tiles[i] = tile
	}

	f.image = ebiten.NewImage(int(f.Bounds.X), int(f.Bounds.Y))
}

// Converts a 2D Tile index to 1D a index.
func index(q int, r int, size int) int {
	return q + r*size
}

// Converts a 1D Tile index to a 2D index.
func reverseIndex(i int, size int) (q int, r int) {
	return i % size, i / size
}

func (f *Field) GetAxial(pos AxialPos) *Tile {
	q := int(pos.Q)
	r := int(pos.R)
	i := index(q, r, f.Size)
	if i >= 100 || i < 0 ||
		q < 0 || r < 0 ||
		q >= f.Size || r >= f.Size {
		return nil
	}

	tile := &f.Tiles[i]

	return tile
}

func (f *Field) GetCard(pos CardPos) *Tile {
	pos = pos.Sub(f.Pos.Mul(f.Bounds).Add(CardPos{X: tileSize, Y: tileSize}))
	axialPos := pos.ToAxial().DivFloat(tileSize * 2).Round()
	return f.GetAxial(axialPos)
}

func (f *Field) Update() {
	f.image.Clear()

	for _, tile := range f.Tiles {
		tile.draw(f.image, f.Active)
	}
}

func (f *Field) Draw(img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	pos := f.Pos.Mul(f.Bounds)
	op.GeoM.Translate(pos.X, pos.Y)

	img.DrawImage(f.image, op)
}
