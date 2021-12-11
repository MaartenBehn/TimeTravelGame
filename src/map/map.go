package gameMap

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const chunkSizeQ = 10 // The ammount of Tiles per Chunk Q-axis.
const chunkSizeR = 10 // The ammount of Tiles per Chunk R-axis.

type Chunk struct {
	Pos      CardPos
	AxialPos AxialPos
	Tiles    []Tile
}

type Map struct {
	Size   CardPos
	Chunks map[AxialPos]*Chunk

	U *UnitController

	ArrowPos CardPos

	mapImage *ebiten.Image
}

// Converts a 2D Tile index to 1D a index.
func index(q int, r int) int {
	return q + r*chunkSizeQ
}

// Converts a 1D Tile index to a 2D index.
func reverseIndex(i int) (q int, r int) {
	return i % chunkSizeQ, i / chunkSizeQ
}

// NewChunk is the init function for Chunk
func NewChunk(pos AxialPos) *Chunk {
	chunk := &Chunk{
		AxialPos: pos,
		Pos:      pos.MulFloat(tileSize * chunkSizeQ * 2).ToCard(),
		Tiles:    make([]Tile, chunkSizeQ*chunkSizeR),
	}

	for i, _ := range chunk.Tiles {
		q, r := reverseIndex(i)
		chunk.Tiles[i] = NewTile(q, r, chunk)
	}

	return chunk
}

// NewMap is the init func for a new Map
func NewMap(size CardPos) *Map {
	return &Map{
		Size:   size,
		Chunks: map[AxialPos]*Chunk{},
		U:      NewUnitController(),

		mapImage: ebiten.NewImage(int(size.X), int(size.Y)),
	}
}

// GetChunk returns the Chunk for at the corresponding Pos.
// It creates a new Chunk when the Chunk is nill
func (m *Map) GetChunk(pos AxialPos) *Chunk {
	chunk := m.Chunks[pos]
	if chunk == nil {
		chunk = NewChunk(pos)
		m.Chunks[pos] = chunk
	}
	return chunk
}

func (m *Map) GetAxial(pos AxialPos) (*Tile, *Chunk) {
	roundPos := pos.Round()
	chunkPos := roundPos.Div(AxialPos{chunkSizeQ, chunkSizeR}).Trunc()

	if roundPos.Q < 0 && math.Mod(roundPos.Q, 10) != 0 {
		chunkPos.Q -= 1
	}

	if roundPos.R < 0 && math.Mod(roundPos.R, 10) != 0 {
		chunkPos.R -= 1
	}

	tilePos := roundPos.Sub(chunkPos.Mul(AxialPos{chunkSizeQ, chunkSizeR}))

	chunk := m.GetChunk(chunkPos)
	i := index(int(tilePos.Q), int(tilePos.R))
	if i >= 100 || i < 0 {
		fmt.Print("error")
	}

	tile := &chunk.Tiles[i]

	return tile, chunk
}

func (m *Map) GetCard(pos CardPos) (*Tile, *Chunk) {
	t, c := m.GetAxial(pos.ToAxial().DivFloat(tileSize * 2))
	return t, c
}

func (m *Map) Update() {
	m.mapImage.Clear()

	for _, chunk := range m.Chunks {
		chunk.draw(m.mapImage)
	}

	m.U.draw(m.mapImage, m)
}
