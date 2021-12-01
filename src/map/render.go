package gameMap

import (
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// DrawTile draws the hexagon for the Tile
func (t Tile) DrawTile(img *ebiten.Image) {
	if !t.Visable {
		return
	}

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

func (m *Map) DrawMap(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()

	img.DrawImage(m.mapImage, op)
}
