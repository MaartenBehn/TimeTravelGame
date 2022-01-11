package util

import (
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	FontFaceRegular = "res/fonts/NotoSans-Regular.ttf"
	FontFaceBold    = "res/fonts/NotoSans-Bold.ttf"
)

type Font struct {
	Face         font.Face
	TitleFace    font.Face
	BigTitleFace font.Face
	ToolTipFace  font.Face
}

func LoadFonts() (*Font, error) {
	fontFace, err := LoadFont(FontFaceRegular, 20)
	if err != nil {
		return nil, err
	}

	titleFontFace, err := LoadFont(FontFaceBold, 24)
	if err != nil {
		return nil, err
	}

	bigTitleFontFace, err := LoadFont(FontFaceBold, 28)
	if err != nil {
		return nil, err
	}

	toolTipFace, err := LoadFont(FontFaceRegular, 15)
	if err != nil {
		return nil, err
	}

	return &Font{
		Face:         fontFace,
		TitleFace:    titleFontFace,
		BigTitleFace: bigTitleFontFace,
		ToolTipFace:  toolTipFace,
	}, nil
}

func (f *Font) Close() {
	if f.Face != nil {
		_ = f.Face.Close()
	}

	if f.TitleFace != nil {
		_ = f.TitleFace.Close()
	}

	if f.BigTitleFace != nil {
		_ = f.BigTitleFace.Close()
	}
}

func LoadFont(path string, size float64) (font.Face, error) {
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
