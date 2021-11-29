package core

import (
	"math"
)

type Color struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	return c.r, c.g, c.b, c.a
}

var Red = Color{math.MaxUint32, 0, 0, math.MaxUint32}
