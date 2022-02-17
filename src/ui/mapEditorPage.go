package ui

import (
	"github.com/Stroby241/TimeTravelGame/src/event"
	"github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/blizzy78/ebitenui"
	"image"
	"strconv"

	"github.com/blizzy78/ebitenui/widget"
)

func createMapEditorPage(res *uiResources, ui func() *ebitenui.UI) widget.PreferredSizeLocateableWidget {

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

	c1.AddChild(widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Main Menu", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			event.Go(event.EventEditorUnload, nil)
		}),
	))

	entries := []interface{}{
		"Ground",
		"Unit Blue",
		"Unit Red",
		"Arrow",
	}

	c1.AddChild(newListComboButton(
		entries,
		func(e interface{}) string {
			return e.(string)
		},
		func(e interface{}) string {
			return e.(string)
		},
		func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			for i, entry := range entries {
				if entry.(string) == args.Entry.(string) {
					event.Go(event.EventEditorUISetMode, i)
				}
			}
			c.RequestRelayout()
		},
		res))

	return c
}

func openTextInputPopUp(res *uiResources, ui func() *ebitenui.UI, name string, textPlaceholder string, f func(text string) bool, pos math.CardPos) {
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
	r = r.Add(image.Point{X: int(pos.X), Y: int(pos.Y)})
	w.SetLocation(r)

	rw = ui().AddWindow(w)
	text.Focus(true)
}

func openNewMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openTextInputPopUp(res, ui, "New Map", "Set Map Size", func(text string) bool {
		x, err := strconv.Atoi(text)
		if err != nil {
			return false
		}
		event.Go(event.EventEditorUINewMap, x)
		return true
	}, math.CardPos{X: 120, Y: 100})
}

func openSaveMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openTextInputPopUp(res, ui, "Save", "Set Name", func(text string) bool {
		event.Go(event.EventEditorUISaveMap, text)
		return true
	}, math.CardPos{X: 280, Y: 100})
}

func openLoadMapPopUp(res *uiResources, ui func() *ebitenui.UI) {
	openTextInputPopUp(res, ui, "Load", "Set Name", func(text string) bool {
		event.Go(event.EventEditorUILoadMap, text)
		return true
	}, math.CardPos{X: 430, Y: 100})
}
