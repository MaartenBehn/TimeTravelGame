package math

import (
	"math"
)

type CubePos struct {
	X float64
	Y float64
	Z float64
}

var CubeDirections = [6]CubePos{
	{1, -1, 0}, {1, 0, -1}, {0, 1, -1},
	{-1, 1, 0}, {-1, 0, 1}, {0, -1, 1},
}

func (c CubePos) Add(b CubePos) CubePos {
	return CubePos{c.X + b.X, c.Y + b.Y, c.Z + b.Z}
}

func (c CubePos) AddFloat(b float64) CubePos {
	return CubePos{c.X + b, c.Y + b, c.Z + b}
}

func (c CubePos) Sub(b CubePos) CubePos {
	return CubePos{c.X - b.X, c.Y - b.Y, c.Z - b.Z}
}

func (c CubePos) Mul(b CubePos) CubePos {
	return CubePos{c.X * b.X, c.Y * b.Y, c.Z * b.Z}
}

func (c CubePos) MulFloat(b float64) CubePos {
	return CubePos{c.X * b, c.Y * b, c.Z * b}
}

func (c CubePos) Div(b CubePos) CubePos {
	return CubePos{c.X / b.X, c.Y / b.Y, c.Z / b.Z}
}

func (c CubePos) DivFloat(b float64) CubePos {
	return CubePos{c.X / b, c.Y / b, c.Z / b}
}

func (c CubePos) ToAxial() AxialPos {
	return AxialPos{c.X, c.Z}
}
func CubeArrayToAxialArray(cubes []CubePos) (axials []AxialPos) {
	axials = make([]AxialPos, len(cubes))
	for i, cube := range cubes {
		axials[i] = cube.ToAxial()
	}
	return axials
}

func (c CubePos) Distance(b CubePos) float64 {
	return (math.Abs(c.X-b.X) + math.Abs(c.Y-b.Y) + math.Abs(c.Z-b.Z)) / 2
}

func (c CubePos) GetLine(b CubePos) (results []CubePos) {
	N := c.Distance(b)
	for i := 0.0; i < N; i++ {
		results = append(results, c.Lerp(b, 1/N*i).Round())
	}
	return results
}

func (c CubePos) MoveRange(r float64) (results []CubePos) {
	for x := -r; x <= r; x++ {
		for y := math.Max(-r, -x-r); y < math.Min(r, -x+r); y++ {
			results = append(results, CubePos{x, y, -x - y}.Add(c))
		}
	}
	return results
}

func (c CubePos) Round() CubePos {
	rx := math.Round(c.X)
	ry := math.Round(c.Y)
	rz := math.Round(c.Z)

	xdiff := math.Abs(rx - c.X)
	ydiff := math.Abs(ry - c.Y)
	zdiff := math.Abs(rz - c.Z)

	if (xdiff > ydiff) && (xdiff > zdiff) {
		rx = -ry - rz
	} else if ydiff > zdiff {
		ry = -rx - rz
	} else {
		rz = -rx - ry
	}
	return CubePos{rx, ry, rz}
}

func (c CubePos) Lerp(b CubePos, t float64) CubePos {
	return CubePos{
		X: Lerp(c.X, b.X, t),
		Y: Lerp(c.Y, b.Y, t),
		Z: Lerp(c.Z, b.Z, t),
	}
}
