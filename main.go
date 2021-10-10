package main

import (
	"fmt"
	. "github.com/TimeTravelGame/TimeTravelGame/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

const (
	screenWidth  = 1024
	screenHeight = 800
)

var (
	emptyImage = ebiten.NewImage(3, 3)
)

func init() {
	emptyImage.Fill(color.White)
}

type Game struct {
	m *Map
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.m.chunks[AxialPos{0, 0}].DrawChunk(screen)
	g.m.chunks[AxialPos{0, 1}].DrawChunk(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Time Travel Game")

	g := &Game{
		m: NewMap(),
	}

	g.m.GetChunk(AxialPos{0, 0})
	g.m.GetChunk(AxialPos{0, 1})

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
