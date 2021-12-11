package gameMap

import (
	"bytes"
	"encoding/gob"
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
	return &b
}

func Load(b *bytes.Buffer) *Map {

	m := &Map{}
	d := gob.NewDecoder(b)
	checkErr(d.Decode(m))

	m.mapImage = ebiten.NewImage(int(m.Size.X), int(m.Size.X))
	for _, chunk := range m.Chunks {
		for i, tile := range chunk.Tiles {
			tile.chunk = chunk
			tile.makeReady()
			chunk.Tiles[i] = tile
		}
	}
	m.U.makeReady()

	m.Update()

	return m
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
