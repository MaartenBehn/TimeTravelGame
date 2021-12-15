package field

import (
	"bytes"
	"encoding/gob"
	"github.com/Stroby241/TimeTravelGame/src/util"
)

func (f *Field) Save(name string) {

	var buffer bytes.Buffer
	e := gob.NewEncoder(&buffer)
	if err := e.Encode(*f); err != nil {
		panic(err)
	}
	util.SaveMapBufferToFile(name, &buffer)
}

func LoadField(name string) *Field {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}

	f := &Field{}
	d := gob.NewDecoder(buffer)
	checkErr(d.Decode(f))

	f.makeReady()

	for i, tile := range f.Tiles {
		tile.makeReady(f)
		f.Tiles[i] = tile
	}

	f.Update()

	return f
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
