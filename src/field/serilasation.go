package field

import (
	"bytes"
	"encoding/gob"
	"github.com/Stroby241/TimeTravelGame/src/util"
)

func (t *Timeline) Save(name string) {

	var buffer bytes.Buffer
	e := gob.NewEncoder(&buffer)
	if err := e.Encode(*t); err != nil {
		panic(err)
	}
	util.SaveMapBufferToFile(name, &buffer)
}

func LoadTimeline(name string) *Timeline {
	buffer := util.LoadMapBufferFromFile(name)
	if buffer == nil {
		return nil
	}

	t := &Timeline{}
	d := gob.NewDecoder(buffer)
	checkErr(d.Decode(t))

	t.makeReady()
	t.Update()

	return t
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
