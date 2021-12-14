package field

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten"
	"time"
)

const selectorBlinkIntervall = 1

var (
	selectorVisable bool
	selectorTime    time.Time
)

type Field struct {
	Size  int
	Tiles []Tile

	U          *UnitController
	fieldImage *ebiten.Image
}

func NewField(size int) *Field {
	field := &Field{
		Size:  size,
		Tiles: make([]Tile, size*size),
		U:     NewUnitController(),
	}

	for i, _ := range field.Tiles {
		q, r := reverseIndex(i, size)
		field.Tiles[i] = NewTile(q, r, field)
	}

	return field
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
	i := index(int(pos.Q), int(pos.R), f.Size)
	if i >= 100 || i < 0 {
		fmt.Print("error")
	}

	tile := &f.Tiles[i]

	return tile
}

func (f *Field) Update() {
	f.fieldImage.Clear()
	f.draw(f.fieldImage)
	f.U.draw(f.fieldImage, f)
}

// draw draws the Chunk
func (f *Field) draw(img *ebiten.Image) {
	for _, tile := range f.Tiles {
		tile.draw(img)
	}
}
