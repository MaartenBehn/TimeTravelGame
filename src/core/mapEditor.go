package core

import (
	. "github.com/TimeTravelGame/TimeTravelGame/src/math"

	"github.com/blizzy78/ebitenui/widget"
)

func createMapEditor(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			),
		),
	)

	c1 := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(10),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5))),
		),
		widget.ContainerOpts.BackgroundImage(res.background),
	)
	c.AddChild(c1)

	// Debug Print Spacer
	c1.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("           ", res.text.face, res.label.text)),
	)

	c1.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("Map Editor", res.text.face, res.label.text)),
	)

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("New Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
	))

	return c
}

func EditorNewMap(size CardPos) {

}

func updateMapEditor(res *uiResources) {

}
