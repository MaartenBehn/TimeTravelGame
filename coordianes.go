package main

import "math"

type AxialPos struct {
	q float64
	r float64
}

type CubePos struct {
	x float64
	y float64
	z float64
}

type CardPos struct {
	x float64
	y float64
}

var CubeDirections = [6]CubePos{
	{ 1, -1, 0}, { 1, 0, -1}, {0,  1, -1},
	{-1,  1, 0}, {-1, 0,  1}, {0, -1,  1},
}

var AxialDirections = [6]AxialPos{
	{ 1, 0}, { 1, -1}, {0, -1},
	{-1, 0}, {-1,  1}, {0,  1},
}

func CubetoAxial(cube CubePos) AxialPos {
	return AxialPos{
		q: cube.x,
		r: cube.z,
	}
}
func CubeArrayToAxialArray(cubes []CubePos) (axials []AxialPos){
	axials = make([]AxialPos, len(cubes))
	for i, cube := range cubes {
		axials[i] = CubetoAxial(cube)
	}
	return axials
}

func AxialtoCube(axial AxialPos) CubePos {
	return CubePos{
		x: axial.q,
		y: axial.r,
		z: -axial.q - axial.r,
	}
}
func AxialArrayToCubeArray(axials []AxialPos) (cubes []CubePos){
	cubes = make([]CubePos, len(axials))
	for i, axial := range axials {
		cubes[i] = AxialtoCube(axial)
	}
	return cubes
}

func AxialtoCard(pos AxialPos) CardPos{
	return CardPos{pos.r * 0.75, pos.r * 0.5 + pos.q}
}
func AxialArrayToCardArray(axials []AxialPos) (cards []CardPos){
	cards = make([]CardPos, len(axials))
	for i, axial := range axials {
		cards[i] = AxialtoCard(axial)
	}
	return cards
}

func CubeDistance(a CubePos, b CubePos) float64{
	return (math.Abs(a.x - b.x) + math.Abs(a.y - b.y) + math.Abs(a.z - b.z)) / 2
}

func AxialDistance(a AxialPos, b AxialPos) float64{
	return (math.Abs(a.q - b.q) +
		math.Abs(a.q + a.r - b.q + b.r) +
		math.Abs(a.r - a.r)) /2
}

func Lerp(a float64, b float64, t float64) float64{
	return a + (b - a) * t
}

func CubeLerp(a CubePos, b CubePos, t float64) CubePos{
	return CubePos{
		x: Lerp(a.x, b.x, t),
		y: Lerp(a.y, b.y, t),
		z: Lerp(a.z, b.z, t),
	}
}

func CubeGetLine(a CubePos, b CubePos) (results []CubePos){
	N := CubeDistance(a, b)
	for i := 0.0; i < N; i++ {
		results = append(results, CubeRound(CubeLerp(a, b, 1 / N * i)))
	}
	return results
}

func CubeMoveRange(r float64) (results []CubePos){
	for x := -r; x <= r; x++ {
		for y := math.Max(-r, -x-r); y < math.Min(r, -x+r); y++ {
			results = append(results, CubePos{x, y, -x-y})
		}
	}
	return results
}

func CubeRound(cube CubePos) CubePos{
	rx := math.Round(cube.x)
	ry := math.Round(cube.y)
	rz := math.Round(cube.z)

	xdiff := math.Abs(rx - cube.x)
	ydiff := math.Abs(ry - cube.y)
	zdiff := math.Abs(rz - cube.z)

	if (xdiff > ydiff) && (xdiff > zdiff){
		rx = -ry-rz
	}else if ydiff > zdiff {
		ry = -rx-rz
	}else{
		rz = -rx-ry
	}
	return CubePos{rx, ry, rz}
}


