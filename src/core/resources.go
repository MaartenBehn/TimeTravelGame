package core

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"image/color"
	"strconv"
)

/*
- 100% — FF
- 99% — FC
- 98% — FA
- 97% — F7
- 96% — F5
- 95% — F2
- 94% — F0
- 93% — ED
- 92% — EB
- 91% — E8
- 90% — E6
- 89% — E3
- 88% — E0
- 87% — DE
- 86% — DB
- 85% — D9
- 84% — D6
- 83% — D4
- 82% — D1
- 81% — CF
- 80% — CC
- 79% — C9
- 78% — C7
- 77% — C4
- 76% — C2
- 75% — BF
- 74% — BD
- 73% — BA
- 72% — B8
- 71% — B5
- 70% — B3
- 69% — B0
- 68% — AD
- 67% — AB
- 66% — A8
- 65% — A6
- 64% — A3
- 63% — A1
- 62% — 9E
- 61% — 9C
- 60% — 99
- 59% — 96
- 58% — 94
- 57% — 91
- 56% — 8F
- 55% — 8C
- 54% — 8A
- 53% — 87
- 52% — 85
- 51% — 82
- 50% — 80
- 49% — 7D
- 48% — 7A
- 47% — 78
- 46% — 75
- 45% — 73
- 44% — 70
- 43% — 6E
- 42% — 6B
- 41% — 69
- 40% — 66
- 39% — 63
- 38% — 61
- 37% — 5E
- 36% — 5C
- 35% — 59
- 34% — 57
- 33% — 54
- 32% — 52
- 31% — 4F
- 30% — 4D
- 29% — 4A
- 28% — 47
- 27% — 45
- 26% — 42
- 25% — 40
- 24% — 3D
- 23% — 3B
- 22% — 38
- 21% — 36
- 20% — 33
- 19% — 30
- 18% — 2E
- 17% — 2B
- 16% — 29
- 15% — 26
- 14% — 24
- 13% — 21
- 12% — 1F
- 11% — 1C
- 10% — 1A
- 9% — 17
- 8% — 14
- 7% — 12
- 6% — 0F
- 5% — 0D
- 4% — 0A
- 3% — 08
- 2% — 05
- 1% — 03
- 0% — 00
*/

const (
	backgroundColor = "131a22"

	textIdleColor     = "dff4ff"
	textDisabledColor = "5a7a91"

	labelIdleColor     = textIdleColor
	labelDisabledColor = textDisabledColor

	buttonIdleColor     = textIdleColor
	buttonDisabledColor = labelDisabledColor

	listSelectedBackground         = "4b687a"
	listDisabledSelectedBackground = "2a3944"

	headerColor = textIdleColor

	textInputCaretColor         = "e7c34b"
	textInputDisabledCaretColor = "766326"

	toolTipColor = backgroundColor

	separatorColor = listDisabledSelectedBackground
)

type uiResources struct {
	fonts *fonts

	background *image.NineSlice

	separatorColor color.Color

	text        *textResources
	button      *buttonResources
	label       *labelResources
	checkbox    *checkboxResources
	comboButton *comboButtonResources
	list        *listResources
	slider      *sliderResources
	panel       *panelResources
	tabBook     *tabBookResources
	header      *headerResources
	textInput   *textInputResources
	toolTip     *toolTipResources
}

type textResources struct {
	idleColor     color.Color
	disabledColor color.Color
	face          font.Face
	titleFace     font.Face
	bigTitleFace  font.Face
	smallFace     font.Face
}

type buttonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	padding widget.Insets
}

type checkboxResources struct {
	image   *widget.ButtonImage
	graphic *widget.CheckboxGraphicImage
	spacing int
}

type labelResources struct {
	text *widget.LabelColor
	face font.Face
}

type comboButtonResources struct {
	image   *widget.ButtonImage
	text    *widget.ButtonTextColor
	face    font.Face
	graphic *widget.ButtonImageImage
	padding widget.Insets
}

type listResources struct {
	image        *widget.ScrollContainerImage
	track        *widget.SliderTrackImage
	trackPadding widget.Insets
	handle       *widget.ButtonImage
	handleSize   int
	face         font.Face
	entry        *widget.ListEntryColor
	entryPadding widget.Insets
}

type sliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize int
}

type panelResources struct {
	image   *image.NineSlice
	padding widget.Insets
}

type tabBookResources struct {
	idleButton     *widget.ButtonImage
	selectedButton *widget.ButtonImage
	buttonFace     font.Face
	buttonText     *widget.ButtonTextColor
	buttonPadding  widget.Insets
}

type headerResources struct {
	background *image.NineSlice
	padding    widget.Insets
	face       font.Face
	color      color.Color
}

type textInputResources struct {
	image   *widget.TextInputImage
	padding widget.Insets
	face    font.Face
	color   *widget.TextInputColor
}

type toolTipResources struct {
	background *image.NineSlice
	padding    widget.Insets
	face       font.Face
	color      color.Color
}

func newUIResources() (*uiResources, error) {
	background := image.NewNineSliceColor(hexToColor(backgroundColor))

	fonts, err := loadFonts()
	if err != nil {
		return nil, err
	}

	button, err := newButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	checkbox, err := newCheckboxResources()
	if err != nil {
		return nil, err
	}

	comboButton, err := newComboButtonResources(fonts)
	if err != nil {
		return nil, err
	}

	list, err := newListResources(fonts)
	if err != nil {
		return nil, err
	}

	slider, err := newSliderResources()
	if err != nil {
		return nil, err
	}

	panel, err := newPanelResources()
	if err != nil {
		return nil, err
	}

	tabBook, err := newTabBookResources(fonts)
	if err != nil {
		return nil, err
	}

	header, err := newHeaderResources(fonts)
	if err != nil {
		return nil, err
	}

	textInput, err := newTextInputResources(fonts)
	if err != nil {
		return nil, err
	}

	toolTip, err := newToolTipResources(fonts)
	if err != nil {
		return nil, err
	}

	return &uiResources{
		fonts: fonts,

		background: background,

		separatorColor: hexToColor(separatorColor),

		text: &textResources{
			idleColor:     hexToColor(textIdleColor),
			disabledColor: hexToColor(textDisabledColor),
			face:          fonts.face,
			titleFace:     fonts.titleFace,
			bigTitleFace:  fonts.bigTitleFace,
			smallFace:     fonts.toolTipFace,
		},

		button:      button,
		label:       newLabelResources(fonts),
		checkbox:    checkbox,
		comboButton: comboButton,
		list:        list,
		slider:      slider,
		panel:       panel,
		tabBook:     tabBook,
		header:      header,
		textInput:   textInput,
		toolTip:     toolTip,
	}, nil
}

func newButtonResources(fonts *fonts) (*buttonResources, error) {
	idle, err := loadImageNineSlice("res/graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	i := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &buttonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		face: fonts.face,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newCheckboxResources() (*checkboxResources, error) {
	idle, err := loadImageNineSlice("res/graphics/checkbox-idle.png", 20, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/checkbox-hover.png", 20, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/checkbox-disabled.png", 20, 0)
	if err != nil {
		return nil, err
	}

	checked, err := loadGraphicImages("res/graphics/checkbox-checked-idle.png", "res/graphics/checkbox-checked-disabled.png")
	if err != nil {
		return nil, err
	}

	unchecked, err := loadGraphicImages("res/graphics/checkbox-unchecked-idle.png", "res/graphics/checkbox-unchecked-disabled.png")
	if err != nil {
		return nil, err
	}

	greyed, err := loadGraphicImages("res/graphics/checkbox-greyed-idle.png", "res/graphics/checkbox-greyed-disabled.png")
	if err != nil {
		return nil, err
	}

	return &checkboxResources{
		image: &widget.ButtonImage{
			Idle:     idle,
			Hover:    hover,
			Pressed:  hover,
			Disabled: disabled,
		},

		graphic: &widget.CheckboxGraphicImage{
			Checked:   checked,
			Unchecked: unchecked,
			Greyed:    greyed,
		},

		spacing: 10,
	}, nil
}

func newLabelResources(fonts *fonts) *labelResources {
	return &labelResources{
		text: &widget.LabelColor{
			Idle:     hexToColor(labelIdleColor),
			Disabled: hexToColor(labelDisabledColor),
		},

		face: fonts.face,
	}
}

func newComboButtonResources(fonts *fonts) (*comboButtonResources, error) {
	idle, err := loadImageNineSlice("res/graphics/combo-button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/combo-button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/combo-button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/combo-button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	i := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	arrowDown, err := loadGraphicImages("res/graphics/arrow-down-idle.png", "res/graphics/arrow-down-disabled.png")
	if err != nil {
		return nil, err
	}

	return &comboButtonResources{
		image: i,

		text: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		face:    fonts.face,
		graphic: arrowDown,

		padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newListResources(fonts *fonts) (*listResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/list-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/list-disabled.png")
	if err != nil {
		return nil, err
	}

	mask, _, err := ebitenutil.NewImageFromFile("res/graphics/list-mask.png")
	if err != nil {
		return nil, err
	}

	trackIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/list-track-idle.png")
	if err != nil {
		return nil, err
	}

	trackDisabled, _, err := ebitenutil.NewImageFromFile("res/graphics/list-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	return &listResources{
		image: &widget.ScrollContainerImage{
			Idle:     image.NewNineSlice(idle, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(disabled, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Mask:     image.NewNineSlice(mask, [3]int{26, 10, 23}, [3]int{26, 10, 26}),
		},

		track: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Hover:    image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(trackDisabled, [3]int{0, 5, 0}, [3]int{25, 12, 25}),
		},

		trackPadding: widget.Insets{
			Top:    5,
			Bottom: 24,
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleIdle, 0, 5),
		},

		handleSize: 5,
		face:       fonts.face,

		entry: &widget.ListEntryColor{
			Unselected:         hexToColor(textIdleColor),
			DisabledUnselected: hexToColor(textDisabledColor),

			Selected:         hexToColor(textIdleColor),
			DisabledSelected: hexToColor(textDisabledColor),

			SelectedBackground:         hexToColor(listSelectedBackground),
			DisabledSelectedBackground: hexToColor(listDisabledSelectedBackground),
		},

		entryPadding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    2,
			Bottom: 2,
		},
	}, nil
}

func newSliderResources() (*sliderResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-track-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	handleDisabled, _, err := ebitenutil.NewImageFromFile("res/graphics/slider-handle-disabled.png")
	if err != nil {
		return nil, err
	}

	return &sliderResources{
		trackImage: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Hover:    image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Disabled: image.NewNineSlice(disabled, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleDisabled, 0, 5),
		},

		handleSize: 6,
	}, nil
}

func newPanelResources() (*panelResources, error) {
	i, err := loadImageNineSlice("res/graphics/panel-idle.png", 10, 10)
	if err != nil {
		return nil, err
	}

	return &panelResources{
		image: i,
		padding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    20,
			Bottom: 20,
		},
	}, nil
}

func newTabBookResources(fonts *fonts) (*tabBookResources, error) {
	selectedIdle, err := loadImageNineSlice("res/graphics/button-selected-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedHover, err := loadImageNineSlice("res/graphics/button-selected-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedPressed, err := loadImageNineSlice("res/graphics/button-selected-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedDisabled, err := loadImageNineSlice("res/graphics/button-selected-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selected := &widget.ButtonImage{
		Idle:     selectedIdle,
		Hover:    selectedHover,
		Pressed:  selectedPressed,
		Disabled: selectedDisabled,
	}

	idle, err := loadImageNineSlice("res/graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("res/graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("res/graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("res/graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	unselected := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &tabBookResources{
		selectedButton: selected,
		idleButton:     unselected,
		buttonFace:     fonts.face,

		buttonText: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},

		buttonPadding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

func newHeaderResources(fonts *fonts) (*headerResources, error) {
	bg, err := loadImageNineSlice("res/graphics/header.png", 446, 9)
	if err != nil {
		return nil, err
	}

	return &headerResources{
		background: bg,

		padding: widget.Insets{
			Left:   25,
			Right:  25,
			Top:    4,
			Bottom: 4,
		},

		face:  fonts.bigTitleFace,
		color: hexToColor(headerColor),
	}, nil
}

func newTextInputResources(fonts *fonts) (*textInputResources, error) {
	idle, _, err := ebitenutil.NewImageFromFile("res/graphics/text-input-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := ebitenutil.NewImageFromFile("res/graphics/text-input-disabled.png")
	if err != nil {
		return nil, err
	}

	return &textInputResources{
		image: &widget.TextInputImage{
			Idle:     image.NewNineSlice(idle, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
			Disabled: image.NewNineSlice(disabled, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
		},

		padding: widget.Insets{
			Left:   8,
			Right:  8,
			Top:    4,
			Bottom: 4,
		},

		face: fonts.face,

		color: &widget.TextInputColor{
			Idle:          hexToColor(textIdleColor),
			Disabled:      hexToColor(textDisabledColor),
			Caret:         hexToColor(textInputCaretColor),
			DisabledCaret: hexToColor(textInputDisabledCaretColor),
		},
	}, nil
}

func newToolTipResources(fonts *fonts) (*toolTipResources, error) {
	bg, _, err := ebitenutil.NewImageFromFile("res/graphics/tool-tip.png")
	if err != nil {
		return nil, err
	}

	return &toolTipResources{
		background: image.NewNineSlice(bg, [3]int{19, 6, 13}, [3]int{19, 5, 13}),

		padding: widget.Insets{
			Left:   15,
			Right:  15,
			Top:    10,
			Bottom: 10,
		},

		face:  fonts.toolTipFace,
		color: hexToColor(toolTipColor),
	}, nil
}

func (u *uiResources) close() {
	u.fonts.close()
}

func hexToColor(h string) color.Color {

	u, err := strconv.ParseUint(h, 16, 0)
	if err != nil {
		panic(err)
	}

	if len(h) == 6 {
		return color.RGBA{
			R: uint8(u & 0xff0000 >> 16),
			G: uint8(u & 0xff00 >> 8),
			B: uint8(u & 0xff),
			A: 255,
		}
	} else if len(h) == 8 {
		return color.RGBA{
			R: uint8(u & 0xff000000 >> 24),
			G: uint8(u & 0xff0000 >> 16),
			B: uint8(u & 0xff00 >> 8),
			A: uint8(u & 0xff),
		}
	}
	return color.RGBA{}
}
