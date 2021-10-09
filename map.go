package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

const chunkSizeQ = 10 // The ammount of Tiles per Chunk Q-axis.
const chunkSizeR = 10 // The ammount of Tiles per Chunk R-axis.

var tileSize = 100.0                   // Scaling factor for Tile size.
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

type Tile struct {
	pos      CardPos
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
func index(x int, y int) int {
	return y + x*chunkSizeQ
}

// Converts a 1D tile index to a 2D index.
func reverseIndex(i int) (x int, y int) {
	return i / chunkSizeQ, i % chunkSizeQ
}

// NewChunk is the init function for Chunk
func NewChunk(pos AxialPos) *Chunk {
	chunk := &Chunk{
		pos:   AxialtoCard(AxialPos{pos.q * tileSize * chunkSizeQ * 2, pos.r * tileSize * chunkSizeR * 2}),
		tiles: make([]Tile, chunkSizeQ*chunkSizeR),
	}

	for i, tile := range chunk.tiles {
		x, y := reverseIndex(i)
		tile.pos = AxialtoCard(AxialPos{float64(x) * tileSize * 2, float64(y) * tileSize * 2})

		tile.vertices = make([]ebiten.Vertex, len(tileVertices))
		for j, vertex := range tileVertices {
			vertex.DstX += float32(tile.pos.x + chunk.pos.x)
			vertex.DstY += float32(tile.pos.y + chunk.pos.y)

			vertex.ColorA = float32(i) / float32(len(chunk.tiles))
			tile.vertices[j] = vertex
		}
		tile.vertices[0].ColorA -= 0.1

		chunk.tiles[i] = tile
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
