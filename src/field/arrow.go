package field

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	lineNone  = 0
	lineUp    = 1
	lineDown  = 2
	lineRigth = 3
	lineLeft  = 4
)

func drawArrow(fromPos CardPos, toPos CardPos, img *ebiten.Image, fraction *Fraction) {
	if util.Debug {
		fmt.Println("---------------- Draw Arrow --------------")
	}

	vec := fromPos.Sub(toPos)

	if vec.X == 0 && vec.Y == 0 {
		return
	}

	var firstLine int
	var secondLine int

	if vec.X > 0 {
		firstLine = lineLeft
	} else if vec.X < 0 {
		firstLine = lineRigth
	}

	if vec.Y > 0 {
		secondLine = lineUp
	} else if vec.Y < 0 {
		secondLine = lineDown
	}

	if math.Abs(vec.X) < math.Abs(vec.Y) {
		l := secondLine
		secondLine = firstLine
		firstLine = l
	}

	// Arrow Tip
	op := &ebiten.DrawImageOptions{}
	w, h := fraction.images["arrow_tip"].Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	if secondLine == lineLeft {
		op.GeoM.Scale(-1, 1)
	} else if secondLine == lineUp {
		op.GeoM.Rotate(-math.Pi / 2)
	} else if secondLine == lineDown {
		op.GeoM.Rotate(math.Pi / 2)
	} else if secondLine == lineNone {
		if firstLine == lineLeft {
			op.GeoM.Scale(-1, 1)
		} else if firstLine == lineUp {
			op.GeoM.Rotate(-math.Pi / 2)
		} else if firstLine == lineDown {
			op.GeoM.Rotate(math.Pi / 2)
		}
	}
	op.GeoM.Translate(toPos.X, toPos.Y)

	// Arrow End
	op2 := &ebiten.DrawImageOptions{}
	w, h = fraction.images["arrow_end"].Size()
	op2.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	if firstLine == lineLeft {
		op2.GeoM.Scale(-1, 1)
	} else if firstLine == lineUp {
		op2.GeoM.Rotate(-math.Pi / 2)
	} else if firstLine == lineDown {
		op2.GeoM.Rotate(math.Pi / 2)
	}

	op2.GeoM.Translate(fromPos.X, fromPos.Y)

	cornerPos := fromPos
	if firstLine == lineRigth || firstLine == lineLeft {
		cornerPos.X += -vec.X
	} else {
		cornerPos.Y += -vec.Y
	}

	op3 := &ebiten.DrawImageOptions{}
	op5 := &ebiten.DrawImageOptions{}
	// Arrow Corner
	if secondLine != lineNone {
		w, h = fraction.images["arrow_corner"].Size()
		op3.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		if firstLine == lineLeft {
			op3.GeoM.Scale(-1, 1)
		}
		if firstLine == lineDown {
			op3.GeoM.Scale(1, -1)
		}
		if secondLine == lineUp {
			op3.GeoM.Scale(1, -1)
		}
		if secondLine == lineRigth {
			op3.GeoM.Scale(-1, 1)
		}
		op3.GeoM.Translate(cornerPos.X, cornerPos.Y)

		if secondLine == lineRigth {
			op5.GeoM.Scale((vec.X+10)/10, 1)
			op5.GeoM.Translate(cornerPos.X-vec.X-5, cornerPos.Y-5)
		} else if secondLine == lineLeft {
			op5.GeoM.Scale((vec.X-10)/10, 1)
			op5.GeoM.Translate(cornerPos.X-vec.X+5, cornerPos.Y-5)
		} else if secondLine == lineDown {
			op5.GeoM.Rotate(math.Pi / 2)
			op5.GeoM.Scale(1, (vec.Y+10)/10)
			op5.GeoM.Translate(cornerPos.X+5, cornerPos.Y-vec.Y-5)
		} else if secondLine == lineUp {
			op5.GeoM.Rotate(math.Pi / 2)
			op5.GeoM.Scale(1, (vec.Y-10)/10)
			op5.GeoM.Translate(cornerPos.X+5, cornerPos.Y-vec.Y+5)
		}
	}

	// First Line
	// TODO: Uses 10 as Image size change to image size
	op4 := &ebiten.DrawImageOptions{}

	if firstLine == lineRigth {
		op4.GeoM.Scale((vec.X+10)/10, 1)
		op4.GeoM.Translate(fromPos.X-vec.X-5, fromPos.Y-5)
	} else if firstLine == lineLeft {
		op4.GeoM.Scale((vec.X-10)/10, 1)
		op4.GeoM.Translate(fromPos.X-vec.X+5, fromPos.Y-5)
	} else if firstLine == lineDown {
		op4.GeoM.Rotate(math.Pi / 2)
		op4.GeoM.Scale(1, (vec.Y+10)/10)
		op4.GeoM.Translate(fromPos.X+5, fromPos.Y-vec.Y-5)
	} else if firstLine == lineUp {
		op4.GeoM.Rotate(math.Pi / 2)
		op4.GeoM.Scale(1, (vec.Y-10)/10)
		op4.GeoM.Translate(fromPos.X+5, fromPos.Y-vec.Y+5)
	}

	img.DrawImage(fraction.images["arrow_tip"], op)
	img.DrawImage(fraction.images["arrow_end"], op2)
	img.DrawImage(fraction.images["arrow_straigth"], op4)

	if secondLine != lineNone {
		img.DrawImage(fraction.images["arrow_corner"], op3)
		img.DrawImage(fraction.images["arrow_straigth"], op5)
	}
}
