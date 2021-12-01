package ui

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/blizzy78/ebitenui"
	"image"
	"strconv"

	"github.com/blizzy78/ebitenui/widget"
)

func createMapEditor(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {

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
		widget.LabelOpts.Text("Map Editor", res.text.face, res.label.text)),
	)

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("New Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			openNewMapPopUp(res, ui)
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Save Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			openSaveMapPopUp(res, ui)
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Load Map", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			openLoadMapPopUp(res, ui)
		}),
	))

	return c
}

func openMapEditorPopUp(res *uiResources, ui func() *ebitenui.UI, name string, textPlaceholder string, f func(text string) bool, horizontalPos int) {
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
		widget.TextOpts.Text(name, res.text.bigTitleFace, res.text.idleColor),
	))

	text := widget.NewTextInput(
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
		widget.TextInputOpts.Placeholder(textPlaceholder),
	)
	c.AddChild(text)

	c1 := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(10),
		)),
	)
	c.AddChild(c1)

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.Text(name, res.button.face, res.button.text),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			if f(text.InputText) {
				rw()
			}
		}),
	))

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.Text("Cancel", res.button.face, res.button.text),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			rw()
		}),
	))

	w := widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(c),
	)

	r := image.Rect(0, 0, 400, 200)
	r = r.Add(image.Point{horizontalPos, 100})
	w.SetLocation(r)

	rw = ui().AddWindow(w)
	text.Focus(true)
}

func openNewMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openMapEditorPopUp(res, ui, "New Map", "Set Map Size", func(text string) bool {
		x, err := strconv.Atoi(text)
		if err != nil {
			return false
		}
		event.Go(event.EventEditorNewMap, CardPos{float64(x), float64(x)})
		return true
	}, 120)
}

func openSaveMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openMapEditorPopUp(res, ui, "Save", "Set Name", func(text string) bool {
		event.Go(event.EventEditorSaveMap, text)
		return true
	}, 280)
}

func openLoadMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openMapEditorPopUp(res, ui, "Load", "Set Name", func(text string) bool {
		event.Go(event.EventEditorLoadMap, text)
		return true
	}, 430)
}
