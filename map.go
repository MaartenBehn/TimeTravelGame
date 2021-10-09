package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

const chunkSizeX = 10
const chunkSizeY = 10

var tileSize = 10.0
var tileHeigth = 2 * tileSize
var tileWith = math.Sqrt(3) * tileSize

var tileVertices = [7]ebiten.Vertex{
	{DstX:   0, 						DstY: 0, 					ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32(-0.5 * tileSize), 	DstY: float32(tileSize),	ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32( 0.5 * tileSize), 	DstY: float32(tileSize),	ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32(tileSize), 		DstY: 0,					ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32( 0.5 * tileSize), 	DstY: float32(-tileSize),	ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32(-0.5 * tileSize), 	DstY: float32(-tileSize),	ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX:   float32(-tileSize), 		DstY: 0,					ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
}

var tileIndecies = []uint16 {
	0, 1, 2,
	0, 2, 3,
	0, 3, 4,
	0, 4, 5,
	0, 5, 6,
	0, 6, 1,
}

type Tile struct {
	pos CardPos
	vertices []ebiten.Vertex
}

type Chunk struct {
	pos AxialPos
	tiles []Tile
}

type Map struct {
	chunks map[AxialPos]*Chunk
}

func index(x int, y int) int {
	return y + x * chunkSizeX
}

func reverseIndex(i int) (x int, y int) {
	return i / chunkSizeX, i % chunkSizeX
}

func NewChunk(pos AxialPos) *Chunk {
	chunk := &Chunk{
		pos:   pos,
		tiles: make([]Tile, chunkSizeX*chunkSizeY),
	}

	for i, tile := range chunk.tiles {
		x, y := reverseIndex(i)
		tile.pos = AxialtoCard(AxialPos{float64(x) * tileSize * 2, float64(y) * tileSize * 2})

		tile.vertices = make([]ebiten.Vertex, len(tileVertices))
		for j, vertex := range tileVertices {
			vertex.DstX += float32(tile.pos.x)
			vertex.DstY += float32(tile.pos.y)

			vertex.ColorA = float32(i) / float32(len(chunk.tiles))
			tile.vertices[j] = vertex
		}

		chunk.tiles[i] = tile
	}

	return chunk
}

func NewMap() *Map {
	return &Map{
		chunks: map[AxialPos]*Chunk{},
	}
}

func (m *Map)GetChunk(pos AxialPos) *Chunk{
	chunk := m.chunks[pos]
	if chunk == nil {
		chunk = NewChunk(pos)
		m.chunks[pos] = chunk
	}
	return chunk
}

func (t Tile)DrawTile(screen *ebiten.Image){

	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	screen.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

func (c *Chunk)DrawChunk(screen *ebiten.Image){
	for _, tile := range c.tiles {
		tile.DrawTile(screen)
	}
}

/*

/  \  /  \  /  \  /  \
\  /  \  /  \  /  \  /
/  \  /  \  /  \  /  \
\  /  \  /  \  /  \  /


*/
