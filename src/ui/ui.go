package ui

import (
	"fmt"
	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
)

type pageContainer struct {
	widget    widget.PreferredSizeLocateableWidget
	titleText *widget.Text
	flipBook  *widget.FlipBook
}

func CreateUI() (*ebitenui.UI, func(), error) {
	res, err := newUIResources()
	if err != nil {
		return nil, nil, err
	}

	drag := &dragContents{
		res: res,
	}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(0, 20))),
	//widget.ContainerOpts.BackgroundImage(res.background)
	)

	toolTips := toolTipContents{
		tips: map[widget.HasWidget]string{},
		res:  res,
	}

	toolTip := widget.NewToolTip(
		widget.ToolTipOpts.Container(rootContainer),
		widget.ToolTipOpts.ContentsCreater(&toolTips),
	)

	dnd := widget.NewDragAndDrop(
		widget.DragAndDropOpts.Container(rootContainer),
		widget.DragAndDropOpts.ContentsCreater(drag),
	)

	var ui *ebitenui.UI
	rootContainer.AddChild(createMapEditor(res, func() *ebitenui.UI {
		return ui
	}))

	ui = &ebitenui.UI{
		Container:   rootContainer,
		ToolTip:     toolTip,
		DragAndDrop: dnd,
	}

	return ui, func() {
		res.close()
	}, nil
}

func createBaseContainer(res *uiResources) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(15))),
	)

	button := widget.NewButton(
		widget.ButtonOpts.Image(res.button.image),
		widget.ButtonOpts.Text("Button", res.button.face, res.button.text),
		widget.ButtonOpts.TextPadding(res.button.padding),
	)
	c.AddChild(button)

	check := newCheckbox("Check", func(args *widget.CheckboxChangedEventArgs) {}, res)
	c.AddChild(check)

	// Combo Button
	entries := []interface{}{}
	for i := 1; i <= 20; i++ {
		entries = append(entries, i)
	}
	cb := newListComboButton(
		entries,
		func(e interface{}) string {
			return fmt.Sprintf("Entry %d", e.(int))
		},
		func(e interface{}) string {
			return fmt.Sprintf("Entry %d", e.(int))
		},
		func(args *widget.ListComboButtonEntrySelectedEventArgs) {
			c.RequestRelayout()
		},
		res)
	c.AddChild(cb)

	return c
}

func newCheckbox(label string, changedHandler widget.CheckboxChangedHandlerFunc, res *uiResources) *widget.LabeledCheckbox {
	return widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.Spacing(res.checkbox.spacing),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(widget.ButtonOpts.Image(res.checkbox.image)),
			widget.CheckboxOpts.Image(res.checkbox.graphic),
			widget.CheckboxOpts.ChangedHandler(func(args *widget.CheckboxChangedEventArgs) {
				if changedHandler != nil {
					changedHandler(args)
				}
			})),
		widget.LabeledCheckboxOpts.LabelOpts(widget.LabelOpts.Text(label, res.label.face, res.label.text)))
}

func newPageContentContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)))
}

func newListComboButton(entries []interface{}, buttonLabel widget.SelectComboButtonEntryLabelFunc, entryLabel widget.ListEntryLabelFunc,
	entrySelectedHandler widget.ListComboButtonEntrySelectedHandlerFunc, res *uiResources) *widget.ListComboButton {

	return widget.NewListComboButton(
		widget.ListComboButtonOpts.SelectComboButtonOpts(
			widget.SelectComboButtonOpts.ComboButtonOpts(
				widget.ComboButtonOpts.ButtonOpts(
					widget.ButtonOpts.Image(res.comboButton.image),
					widget.ButtonOpts.TextPadding(res.comboButton.padding),
				),
			),
		),
		widget.ListComboButtonOpts.Text(res.comboButton.face, res.comboButton.graphic, res.comboButton.text),
		widget.ListComboButtonOpts.ListOpts(
			widget.ListOpts.Entries(entries),
			widget.ListOpts.ScrollContainerOpts(
				widget.ScrollContainerOpts.Image(res.list.image),
			),
			widget.ListOpts.SliderOpts(
				widget.SliderOpts.Images(res.list.track, res.list.handle),
				widget.SliderOpts.HandleSize(res.list.handleSize),
				widget.SliderOpts.TrackPadding(res.list.trackPadding)),
			widget.ListOpts.EntryFontFace(res.list.face),
			widget.ListOpts.EntryColor(res.list.entry),
			widget.ListOpts.EntryTextPadding(res.list.entryPadding),
		),
		widget.ListComboButtonOpts.EntryLabelFunc(buttonLabel, entryLabel),
		widget.ListComboButtonOpts.EntrySelectedHandler(entrySelectedHandler))
}

func newList(entries []interface{}, res *uiResources, widgetOpts ...widget.WidgetOpt) *widget.List {
	return widget.NewList(
		widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(widgetOpts...)),
		widget.ListOpts.ScrollContainerOpts(widget.ScrollContainerOpts.Image(res.list.image)),
		widget.ListOpts.SliderOpts(
			widget.SliderOpts.Images(res.list.track, res.list.handle),
			widget.SliderOpts.HandleSize(res.list.handleSize),
			widget.SliderOpts.TrackPadding(res.list.trackPadding),
		),
		widget.ListOpts.HideHorizontalSlider(),
		widget.ListOpts.Entries(entries),
		widget.ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(string)
		}),
		widget.ListOpts.EntryFontFace(res.list.face),
		widget.ListOpts.EntryColor(res.list.entry),
		widget.ListOpts.EntryTextPadding(res.list.entryPadding),
	)
}

func newSeparator(res *uiResources, ld interface{}) widget.PreferredSizeLocateableWidget {
	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    20,
				Bottom: 20,
			}))),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(ld)))

	c.AddChild(widget.NewGraphic(
		widget.GraphicOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch:   true,
			MaxHeight: 2,
		})),
		widget.GraphicOpts.ImageNineSlice(image.NewNineSliceColor(res.separatorColor)),
	))

	return c
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
