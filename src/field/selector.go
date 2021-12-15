package field

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const selectorBlinkIntervall = 1

var (
	selectorVisable bool
	selectorTime    time.Time
)

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

func (s *Selector) Draw(img *ebiten.Image, cam *util.Camera, f *Field) {
	if time.Since(selectorTime).Seconds() >= selectorBlinkIntervall {
		selectorVisable = !selectorVisable
		selectorTime = time.Now()
	}

	if selectorVisable {
		op := &ebiten.DrawImageOptions{}
		tile := f.GetAxial(s.Pos)

		size := selectorImgMask.Bounds().Size()

		op.GeoM.Translate(tile.Pos.X-float64(size.X)/2, tile.Pos.Y-float64(size.Y)/2)
		op.GeoM.Concat(*cam.GetMatrix())

		img.DrawImage(selectorImgMask, op)
	}
}
