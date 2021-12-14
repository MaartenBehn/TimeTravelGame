package field

import (
	"bytes"
	"encoding/gob"
)

var versions = []string{
	"0.1",
}

func (f *Field) Save() *bytes.Buffer {

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(*f); err != nil {
		panic(err)
	}
	return &b
}

func Load(b *bytes.Buffer) *Field {

	f := &Field{}
	d := gob.NewDecoder(b)
	checkErr(d.Decode(f))

	return f
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
