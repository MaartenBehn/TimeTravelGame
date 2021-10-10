package main

import (
	"fmt"
	. "github.com/TimeTravelGame/TimeTravelGame/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"runtime/debug"
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

	mouseX, mouseY := ebiten.CursorPosition()

	tile, _ := g.m.Get(CardPos{X: float64(mouseX), Y: float64(mouseY)}.ToAxial())

	tile.vertices[0].ColorA = 1

	if panicInfo := recover(); panicInfo != nil {
		fmt.Printf("%v, %s", panicInfo, string(debug.Stack()))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, chunk := range g.m.chunks {
		chunk.DrawChunk(screen)
	}

	mouseX, mouseY := ebiten.CursorPosition()
	axialPos := CardPos{float64(mouseX), float64(mouseY)}.ToAxial()
	roundPos := axialPos.DivFloat(tileSize * 2).Round()

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f\nPos: %d, %d\nAxi: %f, %f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(), mouseX, mouseY, roundPos.Q, roundPos.R))

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

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
