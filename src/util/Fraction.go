package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
)

type Fraction struct {
	Name       string
	Color      color.Color
	ColorLigth color.Color

	Images map[string]*ebiten.Image
}

var Fractions = []Fraction{
	{
		Name:  "red",
		Color: colornames.Red,
	},
	{
		Name:  "blue",
		Color: colornames.Blue,
	},
}

func GetFractionIndex(f *Fraction) int {
	for i, fraction := range Fractions {
		if fraction.Name == f.Name {
			return i
		}
	}
	return -1
}
