package field

import . "github.com/Stroby241/TimeTravelGame/src/math"

type Selector struct {
	Pos     AxialPos
	Visible bool

	blinkVisible bool
	blinkTime    float64
}

func NewSelector() *Selector {
	return &Selector{
		Visible: false,
		Pos:     AxialPos{},
	}
}
