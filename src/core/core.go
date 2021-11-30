package core

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/event"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"runtime/debug"
)

const (
	screenWidth  = 1024
	screenHeight = 840
	maxTPS       = 30
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

var g *Game

func Init() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Time Travel Game")
	ebiten.SetWindowResizable(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetMaxTPS(maxTPS)

	bounds := CardPos{500, 500}

	ui, closeUI, err := ui.CreateUI()
	checkErr(err)
	defer closeUI()

	g = &Game{
		ui: ui,

		cam: NewCamera(CardPos{0, 0}, bounds, CardPos{1, 1}, CardPos{10, 10}),
		u:   NewUnitController(4),
	}

	checkErr(ebiten.RunGame(g))
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	mouse := CardPos{X: float64(mouseX), Y: float64(mouseY)}

	if g.m != nil && ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mapPos := CardPos{}
		mat := *g.cam.matrix
		mat.Invert()

		mapPos.X, mapPos.Y = mat.Apply(mouse.X, mouse.Y)

		tile, _ := g.m.Get(mapPos.ToAxial())
		tile.vertices[0].ColorA = 1
		g.m.Update()
	}

	event.Go(event.EventCamUpdate, nil)

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {

	}

	g.ui.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.m != nil {
		g.m.DrawMap(screen, g.cam)
	}

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
