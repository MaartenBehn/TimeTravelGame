package core

import (
	"github.com/TimeTravelGame/TimeTravelGame/src/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	pos    math.CardPos
	minPos math.CardPos
	maxPos math.CardPos

	scale    math.CardPos
	minScale math.CardPos
	maxScale math.CardPos

	rotation float64
	matrix   *ebiten.GeoM
}

func NewCamera(minPos math.CardPos, maxPos math.CardPos, minScale math.CardPos, maxScale math.CardPos) *Camera {
	cam := &Camera{

		pos:    math.CardPos{0, 0},
		minPos: minPos,
		maxPos: maxPos,

		scale:    math.CardPos{5, 5},
		minScale: minScale,
		maxScale: maxScale,

		rotation: 0,
		matrix:   &ebiten.GeoM{},
	}
	cam.updateMatrix()
	return cam
}

// Updates Cam.matrix
func (c *Camera) updateMatrix() {
	c.matrix.Reset()
	c.matrix.Rotate(c.rotation)
	c.matrix.Translate(c.pos.X*-1, c.pos.Y*-1)
	c.matrix.Scale(c.scale.X, c.scale.Y)
}

// Restes the Cam to the set bounds.
func (c *Camera) bounds() {
	if c.pos.X < c.minPos.X {
		c.pos.X = c.minPos.X
	} else if c.pos.X > c.maxPos.X {
		c.pos.X = c.maxPos.X
	}

	if c.pos.Y < c.minPos.Y {
		c.pos.Y = c.minPos.Y
	} else if c.pos.Y > c.maxPos.Y {
		c.pos.Y = c.maxPos.Y
	}

	if c.scale.X < c.minScale.X {
		c.scale.X = c.minScale.X
	} else if c.scale.X > c.maxScale.X {
		c.scale.X = c.maxScale.X
	}

	if c.scale.Y < c.minScale.Y {
		c.scale.Y = c.minScale.Y
	} else if c.scale.Y > c.maxScale.Y {
		c.scale.Y = c.maxScale.Y
	}
}

// Applies user Input to Cam
func (c *Camera) UpdateInput() {

	needMatrixUpdate := false
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.pos = c.pos.Add(math.CardPos{0, -1})
		needMatrixUpdate = true
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.pos = c.pos.Add(math.CardPos{0, 1})
		needMatrixUpdate = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.pos = c.pos.Add(math.CardPos{-1, 0})
		needMatrixUpdate = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.pos = c.pos.Add(math.CardPos{1, 0})
		needMatrixUpdate = true
	}

	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		needMatrixUpdate = true
	}
	c.scale = c.scale.AddFloat(wheelY)

	if needMatrixUpdate {
		c.bounds()
		c.updateMatrix()
	}
}
