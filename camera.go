package main

import (
	. "github.com/TimeTravelGame/TimeTravelGame/math"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	pos      CardPos
	scale    CardPos
	rotation float64
	matrix   *ebiten.GeoM
}

func NewCamera() *Camera {
	cam := &Camera{
		pos:      CardPos{0, 0},
		scale:    CardPos{5, 5},
		rotation: 0.01,
		matrix:   &ebiten.GeoM{},
	}
	cam.UpdateMatrix()
	return cam
}

func (c *Camera) UpdateMatrix() {
	c.matrix.Reset()
	c.matrix.Rotate(c.rotation)
	c.matrix.Translate(c.pos.X, c.pos.Y)
	c.matrix.Scale(c.scale.X, c.scale.Y)
}

func (c *Camera) UpdateInput() {

	needMatrixUpdate := false
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.pos = c.pos.Add(CardPos{0, -1})
		needMatrixUpdate = true
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.pos = c.pos.Add(CardPos{0, 1})
		needMatrixUpdate = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.pos = c.pos.Add(CardPos{1, 0})
		needMatrixUpdate = true
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.pos = c.pos.Add(CardPos{-1, 0})
		needMatrixUpdate = true
	}

	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		needMatrixUpdate = true
	}
	c.scale = c.scale.AddFloat(wheelY)

	if needMatrixUpdate {
		c.UpdateMatrix()
	}
}
