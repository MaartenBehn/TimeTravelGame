package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

var versions = []string{
	"0.1",
}

func (m *Map) Save() *bytes.Buffer {

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(*m); err != nil {
		panic(err)
	}
	fmt.Println("Encoded Struct ", b)
	return &b
}

func LoadMap(b *bytes.Buffer) *Map {

	m := &Map{}
	d := gob.NewDecoder(b)
	checkErr(d.Decode(m))

	m.mapImage = ebiten.NewImage(int(m.Size.X), int(m.Size.X))
	for _, chunk := range m.Chunks {
		for i, tile := range chunk.Tiles {
			tile.chunk = chunk
			tile.createVertices()
			chunk.Tiles[i] = tile
		}
	}
	m.Update()

	return m
}
