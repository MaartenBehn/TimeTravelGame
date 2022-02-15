package ui

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
)

func createGamePage(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
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

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Main Menu", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventGameUnload, nil)
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Submit", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventGameUISubmit, nil)
		}),
	))

	userLabel := widget.NewLabel(widget.LabelOpts.Text("%s is playing.       ", res.text.face, res.label.text))
	event.On(event.EventGameCurrentUser, func(data interface{}) {
		id := data.(int)
		userLabel.Label = fmt.Sprintf("%s is playing.", util.Fractions[id].Name)
	})
	event.On(event.EventGameWon, func(data interface{}) {
		id := data.(int)
		userLabel.Label = fmt.Sprintf("%s won.", util.Fractions[id].Name)
	})
	c1.AddChild(userLabel)

	return c
}
