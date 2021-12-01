package ui

import (
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
)

func createGame(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)

	c1 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.panel.image),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Padding(res.panel.padding),
			widget.RowLayoutOpts.Spacing(15),
		)),
	)
	c.AddChild(c1)

	c1.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("Game", res.text.face, res.label.text)),
	)

	return c
}
