package core

import (
	"bytes"
	"os"
)

const pathPreFix = "./saves/maps/"
const mapSaveFileSufix = ".mapsave"

func saveBufferToFile(path string, buffer *bytes.Buffer) {
	f, err := os.Create(path)
	checkErr(err)
	defer checkErr(f.Close())

	_, err = f.Write(buffer.Bytes())
	checkErr(err)
}

func saveMapBufferToFile(name string, buffer *bytes.Buffer) {
	saveBufferToFile(pathPreFix+name+mapSaveFileSufix, buffer)
}

func loadBufferFromFile(path string) *bytes.Buffer {
	buf, err := os.ReadFile(path)
	checkErr(err)
	return bytes.NewBuffer(buf)
}

func loadMapBufferFromFile(name string) *bytes.Buffer {
	return loadBufferFromFile(pathPreFix + name + mapSaveFileSufix)
}
