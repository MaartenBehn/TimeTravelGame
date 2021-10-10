package main

import (
	"fmt"
	. "github.com/TimeTravelGame/TimeTravelGame/math"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

const chunkSizeQ = 10 // The ammount of Tiles per Chunk Q-axis.
const chunkSizeR = 10 // The ammount of Tiles per Chunk R-axis.

var tileSize = 10.0                    // Scaling factor for Tile size.
var tileHeigth = 2 * tileSize          // From Tutorial not uesd at the moment
var tileWith = math.Sqrt(3) * tileSize // From Tutorial not uesd at the moment

// The local Vertex coodrs for a Tile
var tileVertices = [7]ebiten.Vertex{
	{DstX: 0, DstY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(-0.5 * tileSize), DstY: float32(tileSize), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(0.5 * tileSize), DstY: float32(tileSize), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(tileSize), DstY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(0.5 * tileSize), DstY: float32(-tileSize), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(-0.5 * tileSize), DstY: float32(-tileSize), ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: float32(-tileSize), DstY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
}

// The corresponding indices
var tileIndecies = []uint16{
	0, 1, 2,
	0, 2, 3,
	0, 3, 4,
	0, 4, 5,
	0, 5, 6,
	0, 6, 1,
}

type TileSettings struct {
	isWalkable bool
}

type Tile struct {
	TileSettings
	pos      CardPos
	chunk    *Chunk
	vertices []ebiten.Vertex
}

type Chunk struct {
	pos   CardPos
	tiles []Tile
}

type Map struct {
	chunks map[AxialPos]*Chunk
}

// Converts a 2D tile index to 1D a index.
func index(q int, r int) int {
	return q + r*chunkSizeQ
}

// Converts a 1D tile index to a 2D index.
func reverseIndex(i int) (q int, r int) {
	return i % chunkSizeQ, i / chunkSizeQ
}

func NewTile(q int, r int, chunk *Chunk) (tile Tile) {
	tile.pos = AxialPos{Q: float64(q), R: float64(r)}.MulFloat(tileSize * 2).ToCard()

	tile.vertices = make([]ebiten.Vertex, len(tileVertices))
	for j, vertex := range tileVertices {
		vertex.DstX += float32(tile.pos.X + chunk.pos.X)
		vertex.DstY += float32(tile.pos.Y + chunk.pos.Y)

		vertex.ColorA = float32(index(q, r)) / float32(len(chunk.tiles))
		tile.vertices[j] = vertex
	}
	tile.vertices[0].ColorA -= 0.1
	return tile
}

// NewChunk is the init function for Chunk
func NewChunk(pos AxialPos) *Chunk {
	chunk := &Chunk{
		pos:   pos.MulFloat(tileSize * chunkSizeQ * 2).ToCard(),
		tiles: make([]Tile, chunkSizeQ*chunkSizeR),
	}

	for i, _ := range chunk.tiles {
		q, r := reverseIndex(i)
		chunk.tiles[i] = NewTile(q, r, chunk)
	}

	return chunk
}

// NewMap is the init func for a new Map
func NewMap() *Map {
	return &Map{
		chunks: map[AxialPos]*Chunk{},
	}
}

// GetChunk returns the Chunk for at the corresponding pos.
// It creates a new Chunk when the Chunk is nill
func (m *Map) GetChunk(pos AxialPos) *Chunk {
	chunk := m.chunks[pos]
	if chunk == nil {
		chunk = NewChunk(pos)
		m.chunks[pos] = chunk
	}
	return chunk
}

func (m *Map) Get(pos AxialPos) (*Tile, *Chunk) {
	roundPos := pos.DivFloat(tileSize * 2).Round()
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

	tile := &chunk.tiles[i]

	return tile, chunk
}

// DrawTile draws the haxgon for the Tile
func (t Tile) DrawTile(screen *ebiten.Image) {

	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	screen.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

// DrawChunk draws the chunk
func (c *Chunk) DrawChunk(screen *ebiten.Image) {
	for _, tile := range c.tiles {
		tile.DrawTile(screen)
	}
}
