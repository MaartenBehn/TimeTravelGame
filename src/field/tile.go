package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

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
	Visable bool
}

type Tile struct {
	TileSettings
	AxialPos AxialPos
	Pos      CardPos

	vertices []ebiten.Vertex
}

func NewTile(q int, r int, field *Field) (tile Tile) {
	tile.AxialPos = AxialPos{Q: float64(q), R: float64(r)}
	tile.makeReady()
	return tile
}

func (t *Tile) makeReady() {
	t.Pos = t.AxialPos.MulFloat(tileSize * 2).ToCard()
	t.Pos = t.Pos.Add(CardPos{tileSize, tileSize})
	t.createVertices()
}

func (t *Tile) createVertices() {
	t.vertices = make([]ebiten.Vertex, len(tileVertices))
	for j, vertex := range tileVertices {
		vertex.DstX += float32(t.Pos.X)
		vertex.DstY += float32(t.Pos.Y)

		vertex.ColorA = 1
		t.vertices[j] = vertex
	}
	t.vertices[0].ColorA -= 0.5
}

// draw draws the hexagon for the Tile
func (t Tile) draw(img *ebiten.Image) {
	if !t.Visable {
		return
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	img.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}
