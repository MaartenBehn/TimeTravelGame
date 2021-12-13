package gameMap

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

type Chunk struct {
	ChunkSize int

	Pos      CardPos
	AxialPos AxialPos
	Tiles    []Tile
}

// NewChunk is the init function for Chunk
func NewChunk(pos AxialPos, size int) *Chunk {
	chunk := &Chunk{
		ChunkSize: size,
		AxialPos:  pos,
		Pos:       pos.MulFloat(tileSize * float64(size) * 2).ToCard(),
		Tiles:     make([]Tile, size*size),
	}

	for i, _ := range chunk.Tiles {
		q, r := reverseIndex(i, size)
		chunk.Tiles[i] = NewTile(q, r, chunk)
	}

	return chunk
}

func (c *Chunk) GetAxial(pos AxialPos) *Tile {
	roundPos := pos.Round()
	chunkPos := getChunkPosFromAxial(roundPos, c.ChunkSize)

	if chunkPos != c.AxialPos {
		return nil
	}

	tilePos := roundPos.Sub(chunkPos.Mul(AxialPos{float64(c.ChunkSize), float64(c.ChunkSize)}))

	i := index(int(tilePos.Q), int(tilePos.R), c.ChunkSize)
	if i >= 100 || i < 0 {
		fmt.Print("error")
	}

	tile := &c.Tiles[i]

	return tile
}
