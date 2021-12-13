package gameMap

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type Map struct {
	ChunkSize int
	Chunks    map[AxialPos]*Chunk

	U *UnitController
	T *Timeline

	mapImage *ebiten.Image
}

// Converts a 2D Tile index to 1D a index.
func index(q int, r int, size int) int {
	return q + r*size
}

// Converts a 1D Tile index to a 2D index.
func reverseIndex(i int, size int) (q int, r int) {
	return i % size, i / size
}

func getChunkPosFromAxial(pos AxialPos, size int) AxialPos {
	roundPos := pos.Round()
	chunkPos := roundPos.Div(AxialPos{float64(size), float64(size)}).Trunc()

	if roundPos.Q < 0 && math.Mod(roundPos.Q, 10) != 0 {
		chunkPos.Q -= 1
	}

	if roundPos.R < 0 && math.Mod(roundPos.R, 10) != 0 {
		chunkPos.R -= 1
	}
	return chunkPos
}

// NewMap is the init func for a new Map
func NewMap(chunkSize int) *Map {
	return &Map{
		ChunkSize: chunkSize,
		Chunks:    map[AxialPos]*Chunk{},
		U:         NewUnitController(),

		mapImage: ebiten.NewImage(1000, 1000),
	}
}

// GetChunk returns the Chunk for at the corresponding Pos.
func (m *Map) GetChunk(pos AxialPos) *Chunk {
	return m.Chunks[pos]
}

func (m *Map) CreateChunk(pos AxialPos) *Chunk {
	chunk := m.GetChunk(pos)
	if chunk != nil {
		return chunk
	}
	chunk = NewChunk(pos, m.ChunkSize)
	m.Chunks[pos] = chunk
	return chunk
}

func (m *Map) GetAxial(pos AxialPos) (*Tile, *Chunk) {
	roundPos := pos.Round()
	chunkPos := getChunkPosFromAxial(roundPos, m.ChunkSize)
	chunk := m.GetChunk(chunkPos)
	if chunk == nil {
		return nil, nil
	}
	tile := chunk.GetAxial(roundPos)
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

// draw draws the Chunk
func (c *Chunk) draw(img *ebiten.Image) {
	for _, tile := range c.Tiles {
		tile.draw(img)
	}
}

func (m *Map) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()

	img.DrawImage(m.mapImage, op)

	m.drawSelector(img, cam)
}

const selectorBlinkIntervall = 1

var (
	selectorVisable bool
	selectorTime    time.Time
)

func (u *UnitController) SetSelector(pos AxialPos) {
	_, _, unit := u.GetUnitAtPos(pos)
	if unit != nil {
		u.SelectedUnit = unit.Pos
	}
}

func (m *Map) drawSelector(img *ebiten.Image, cam *util.Camera) {
	if time.Since(selectorTime).Seconds() >= selectorBlinkIntervall {
		selectorVisable = !selectorVisable
		selectorTime = time.Now()
	}

	if selectorVisable {
		op := &ebiten.DrawImageOptions{}
		tile, _ := m.GetAxial(m.U.SelectedUnit)

		size := selectorImgMask.Bounds().Size()

		op.GeoM.Translate(tile.Pos.X-float64(size.X)/2, tile.Pos.Y-float64(size.Y)/2)
		op.GeoM.Concat(*cam.GetMatrix())

		img.DrawImage(selectorImgMask, op)
	}
}
