package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Field struct {
	Size  int
	Tiles []Tile

	image *ebiten.Image

	Pos CardPos
}

func NewField(size int) *Field {
	field := &Field{
		Size:  size,
		Tiles: make([]Tile, size*size),
	}

	for i, _ := range field.Tiles {
		q, r := reverseIndex(i, size)
		field.Tiles[i] = NewTile(q, r, field)
	}

	field.makeReady()
	field.Update()

	return field
}

func (f *Field) makeReady() {
	pos := AxialPos{Q: float64(f.Size), R: float64(f.Size)}.MulFloat(tileSize * 2).ToCard()
	f.image = ebiten.NewImage(int(pos.X), int(pos.Y))
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
	if i >= 100 || i < 0 || q >= f.Size || r >= f.Size {
		return nil
	}

	tile := &f.Tiles[i]

	return tile
}

func (f *Field) GetCard(pos CardPos) *Tile {
	return f.GetAxial(pos.ToAxial().DivFloat(tileSize * 2).Round())
}

func (f *Field) Update() {
	f.image.Clear()

	for _, tile := range f.Tiles {
		tile.draw(f.image)
	}
}

func (f *Field) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()

	img.DrawImage(f.image, op)
}
