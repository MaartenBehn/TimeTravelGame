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
	m   *Map
	cam *Camera
	u   *UnitController
}

func (g *Game) Update() error {
	defer ebitenErrorHandle()

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

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		data := g.m.Save()
		saveMapBufferToFile("test", data)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	defer ebitenErrorHandle()

	g.m.DrawMap(screen, g.cam)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS()))

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	defer ebitenErrorHandle()

	return screenWidth, screenHeight
}

const loadMap = true

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Time Travel Game")

	bounds := CardPos{500, 500}

	var m *Map
	if loadMap {
		m = LoadMap(loadMapBufferFromFile("test"))
	} else {
		m = NewMap(CardPos{500, 500})
	}

	g := &Game{
		m:   m,
		cam: NewCamera(CardPos{0, 0}, bounds, CardPos{1, 1}, CardPos{10, 10}),
		u:   NewUnitController(4),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func ebitenErrorHandle() {
	if err := recover(); err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
