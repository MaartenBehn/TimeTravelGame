package core

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/editor"
	"github.com/Stroby241/TimeTravelGame/src/event"
	gameMap "github.com/Stroby241/TimeTravelGame/src/field"
	"github.com/Stroby241/TimeTravelGame/src/game"
	"github.com/Stroby241/TimeTravelGame/src/ui"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	maxTPS = 30
)

type Game struct {
	ui *ebitenui.UI
}

var g *Game

func Init() {
	event.Init()
	gameMap.Init()
	editor.Init()
	game.Init()

	ebiten.SetWindowSize(1024, 840)
	ebiten.SetWindowTitle("Time Travel Game")
	ebiten.SetWindowResizable(true)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetMaxTPS(maxTPS)

	uiObj, closeUI, err := ui.CreateUI()
	checkErr(err)
	defer closeUI()

	g = &Game{
		ui: uiObj,
	}

	event.Go(event.EventUIShowPanel, ui.PageStart)

	checkErr(ebiten.RunGame(g))
}

func (g *Game) Update() error {
	g.ui.Update()

	event.Go(event.EventUpdate, nil)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	event.Go(event.EventDraw, screen)

	g.ui.Draw(screen)

	if util.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.02f\nFPS: %.02f\n",
			ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
