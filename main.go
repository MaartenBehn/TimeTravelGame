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
	m   *Map
	cam *Camera
}

func (g *Game) Update() error {

	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {

		mapPos := CardPos{}
		mat := *g.cam.matrix
		mat.Invert()
		mapPos.X, mapPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := g.m.Get(mapPos.ToAxial())
		tile.vertices[0].ColorA = 1
		g.m.Update()
	}

	g.cam.UpdateInput()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.m.DrawMap(screen, g.cam)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS()))

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Time Travel Game")

	g := &Game{
		m:   NewMap(CardPos{500, 500}),
		cam: NewCamera(),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
