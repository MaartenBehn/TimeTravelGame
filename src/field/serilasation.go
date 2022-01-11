package field

import (
	"bytes"
	"encoding/gob"
	"github.com/Stroby241/TimeTravelGame/src/util"
)

func (t *Timeline) ToBuffer() *bytes.Buffer {
	var buffer bytes.Buffer
	e := gob.NewEncoder(&buffer)
	err := e.Encode(*t)
	checkErr(err)

	return &buffer
}

func FromBuffer(buffer *bytes.Buffer) *Timeline {
	t := &Timeline{}
	d := gob.NewDecoder(buffer)
	checkErr(d.Decode(t))
	t.makeReadyUnits()
	return t
}

func (t *Timeline) Save(name string) {
	buffer := t.ToBuffer()
	util.SaveMapBufferToFile(name, buffer)
}

func Load(name string) *Timeline {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}
	t := FromBuffer(buffer)
	t.MakeReadyUI()
	t.Update()
	return t
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
