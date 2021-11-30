package ui

import (
	"fmt"
	"github.com/Stroby241/TimeTravelGame/src/event"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"image"

	"github.com/blizzy78/ebitenui/widget"
)

func createMapEditor(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {
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
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventEditorNewMap, CardPos{500, 500})
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Save Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			openNewMapWindow(res, ui)
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Load Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventEditorLoadMap, "test")
		}),
	))

	return c
}

func openNewMapWindow(res *uiResources, ui func() *ebitenui.UI) {
	var rw ebitenui.RemoveWindowFunc

	c := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(res.panel.image),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(res.panel.padding),
			widget.RowLayoutOpts.Spacing(15),
		)),
	)

	c.AddChild(widget.NewText(
		widget.TextOpts.Text("Save", res.text.bigTitleFace, res.text.idleColor),
	))

	c.AddChild(widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.TextInputOpts.Image(res.textInput.image),
		widget.TextInputOpts.Color(res.textInput.color),
		widget.TextInputOpts.Padding(widget.Insets{
			Left:   13,
			Right:  13,
			Top:    7,
			Bottom: 7,
		}),
		widget.TextInputOpts.Face(res.textInput.face),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(res.textInput.face, 2),
		),
		widget.TextInputOpts.Placeholder("Enter text here"),
		widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
			fmt.Println(args)
		}),
	))

	c.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.Text("Save", res.button.face, res.button.text),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventEditorSaveMap, "test")
			rw()
		}),
	))

	w := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(c),
	)

	ww, wh := ebiten.WindowSize()
	r := image.Rect(0, 0, ww/2, wh/2)
	r = r.Add(image.Point{ww * 4 / 10, wh / 2 / 2})
	w.SetLocation(r)

	rw = ui().AddWindow(w)
}
