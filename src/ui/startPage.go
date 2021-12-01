package ui

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/blizzy78/ebitenui/widget"
)

func createStartPage(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(5)),
		)),
	)

	c.AddChild(widget.NewLabel(
		widget.LabelOpts.Text("Time Travel Game", res.text.face, res.label.text)),
	)

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Start Game", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Map Editor", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventEditorLoad, nil)
		}),
	))

	return c
}
