package math

import (
	"math"
)

type AxialPos struct {
	Q float64
	R float64
}

var AxialDirections = [6]AxialPos{
	{1, 0}, {1, -1}, {0, -1},
	{-1, 0}, {-1, 1}, {0, 1},
}

func (a AxialPos) Add(b AxialPos) AxialPos {
	return AxialPos{a.Q + b.Q, a.R + b.R}
}

func (a AxialPos) Sub(b AxialPos) AxialPos {
	return AxialPos{a.Q - b.Q, a.R - b.R}
}

func (a AxialPos) Mul(b AxialPos) AxialPos {
	return AxialPos{a.Q * b.Q, a.R * b.R}
}

func (a AxialPos) MulFloat(b float64) AxialPos {
	return AxialPos{a.Q * b, a.R * b}
}

func (a AxialPos) Div(b AxialPos) AxialPos {
	return AxialPos{a.Q / b.Q, a.R / b.R}
}

func (a AxialPos) DivFloat(b float64) AxialPos {
	return AxialPos{a.Q / b, a.R / b}
}

func (a AxialPos) ToCube() CubePos {
	return CubePos{
		X: a.Q,
		Y: a.R,
		Z: -a.Q - a.R,
	}
}

func AxialArrayToCubeArray(axials []AxialPos) (cubes []CubePos) {
	cubes = make([]CubePos, len(axials))
	for i, axial := range axials {
		cubes[i] = axial.ToCube()
	}
	return cubes
}

func (a AxialPos) ToCard() CardPos {
	return CardPos{a.R * 0.75, a.R*0.5 + a.Q}
}
func AxialArrayToCardArray(axials []AxialPos) (cards []CardPos) {
	cards = make([]CardPos, len(axials))
	for i, axial := range axials {
		cards[i] = axial.ToCard()
	}
	return cards
}

func (a AxialPos) Distance(b AxialPos) float64 {
	return (math.Abs(a.Q-b.Q) +
		math.Abs(a.Q+a.R-b.Q+b.R) +
		math.Abs(a.R-a.R)) / 2
}

func (a AxialPos) Round() AxialPos {
	return a.ToCube().Round().ToAxial()
}
