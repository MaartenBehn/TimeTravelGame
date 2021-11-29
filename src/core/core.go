package core

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
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
	ui  *ebitenui.UI
}

const loadMap = true

func Init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Time Travel Game")

	bounds := math.CardPos{500, 500}

	var m *Map
	if loadMap {
		m = LoadMap(loadMapBufferFromFile("test"))
	} else {
		m = NewMap(math.CardPos{500, 500})
	}

	ui, closeUI, err := CreateUI()
	checkErr(err)
	defer closeUI()

	g := &Game{
		m:  m,
		ui: ui,

		cam: NewCamera(math.CardPos{0, 0}, bounds, math.CardPos{1, 1}, math.CardPos{10, 10}),
		u:   NewUnitController(4),
	}

	checkErr(ebiten.RunGame(g))
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	mouse := math.CardPos{X: float64(mouseX), Y: float64(mouseY)}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mapPos := math.CardPos{}
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

	g.ui.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.m.DrawMap(screen, g.cam)

	g.ui.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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
