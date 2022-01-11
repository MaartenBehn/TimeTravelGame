package field

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

var tileSize = 10.0 // Scaling factor for Tile Size.
var useTileImg = true
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
	TimePos

	vertices []ebiten.Vertex
}

func NewTile(pos TimePos) (tile Tile) {
	tile.TimePos = pos
	tile.makeReady()
	return tile
}

func (t *Tile) makeReady() {
	t.vertices = make([]ebiten.Vertex, len(tileVertices))
	for j, vertex := range tileVertices {
		vertex.DstX += float32(t.CalcTilePos().X)
		vertex.DstY += float32(t.CalcTilePos().Y)

		vertex.ColorA = 1
		t.vertices[j] = vertex
	}
	t.vertices[0].ColorA = 0.5
}

// draw draws the hexagon for the Tile
func (t Tile) draw(img *ebiten.Image, active bool) {
	if !t.Visable {
		return
	}

	if active {
		t.vertices[0].ColorR = 0.5
	} else {
		t.vertices[0].ColorR = 1
	}

	if useTileImg {
		w, h := tileImgae.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(t.CalcTilePos().X-float64(w)/2, t.CalcTilePos().Y-float64(h)/2)

		img.DrawImage(tileImgae, op)
	} else {
		op := &ebiten.DrawTrianglesOptions{}
		op.Address = ebiten.AddressUnsafe

		img.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
	}
}
