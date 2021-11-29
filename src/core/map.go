package core

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
	"math/rand"
)

const chunkSizeQ = 10 // The ammount of Tiles per Chunk Q-axis.
const chunkSizeR = 10 // The ammount of Tiles per Chunk R-axis.

var tileSize = 10.0                    // Scaling factor for Tile Size.
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
	IsWalkable bool
}

type Tile struct {
	TileSettings
	Pos      CardPos
	ChunkPos AxialPos

	chunk    *Chunk
	vertices []ebiten.Vertex
}

type Chunk struct {
	Pos      CardPos
	AxialPos AxialPos
	Tiles    []Tile
}

type Map struct {
	Size   CardPos
	Chunks map[AxialPos]*Chunk

	mapImage *ebiten.Image
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
	tile.Pos = AxialPos{Q: float64(q), R: float64(r)}.MulFloat(tileSize * 2).ToCard()
	tile.ChunkPos = chunk.AxialPos
	tile.chunk = chunk
	tile.createVertices()
	return tile
}

func (t *Tile) createVertices() {
	t.vertices = make([]ebiten.Vertex, len(tileVertices))
	for j, vertex := range tileVertices {
		vertex.DstX += float32(t.Pos.X + t.chunk.Pos.X)
		vertex.DstY += float32(t.Pos.Y + t.chunk.Pos.Y)

		vertex.ColorA = rand.Float32()
		t.vertices[j] = vertex
	}
	t.vertices[0].ColorA -= 0.1
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
		Size:     size,
		Chunks:   map[AxialPos]*Chunk{},
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

	tile := &chunk.Tiles[i]

	return tile, chunk
}

// DrawTile draws the hexagon for the Tile
func (t Tile) DrawTile(img *ebiten.Image) {

	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	img.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

// DrawChunk draws the Chunk
func (c *Chunk) DrawChunk(img *ebiten.Image) {
	for _, tile := range c.Tiles {
		tile.DrawTile(img)
	}
}

func (m *Map) DrawMap(img *ebiten.Image, cam *Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.matrix

	img.DrawImage(m.mapImage, op)
}

func (m *Map) Update() {
	m.mapImage.Clear()

	for _, chunk := range m.Chunks {
		chunk.DrawChunk(m.mapImage)
	}
}
