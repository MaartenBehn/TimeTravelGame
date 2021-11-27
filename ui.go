package main

import . "github.com/TimeTravelGame/TimeTravelGame/math"

type text struct {
	pos  CardPos
	size CardPos
}

type button struct {
	pos   CardPos
	size  CardPos
	event func()
}

type uiLayout struct {
	buttons []button
}
