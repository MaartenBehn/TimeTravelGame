package util

import (
	"bytes"
	"fmt"
	"os"
)

const pathPreFix = "./saves/maps/"
const mapSaveFileSufix = ".mapsave"

func SaveBufferToFile(path string, buffer *bytes.Buffer) {
	f, err := os.Create(path)
	checkErr(err)

	_, err = f.Write(buffer.Bytes())
	checkErr(err)

	defer checkErr(f.Close())
}

func SaveMapBufferToFile(name string, buffer *bytes.Buffer) {
	SaveBufferToFile(pathPreFix+name+mapSaveFileSufix, buffer)
}

func LoadBufferFromFile(path string) *bytes.Buffer {
	buf, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("File %s not found.", path)
		return nil
	}
	return bytes.NewBuffer(buf)
}

func LoadMapBufferFromFile(name string) *bytes.Buffer {
	return LoadBufferFromFile(pathPreFix + name + mapSaveFileSufix)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
