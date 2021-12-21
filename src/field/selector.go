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
	TimePos
	Visible bool

	blinkVisible bool
	blinkTime    float64
}

func NewSelector() *Selector {
	return &Selector{
		Visible: false,
		TimePos: TimePos{
			FieldPos: CardPos{},
			TilePos:  AxialPos{},
		},
	}
}

func (s *Selector) Draw(img *ebiten.Image, cam *util.Camera, f *Field) {
	if time.Since(selectorTime).Seconds() >= selectorBlinkIntervall {
		selectorVisable = !selectorVisable
		selectorTime = time.Now()
	}

	if selectorVisable {
		op := &ebiten.DrawImageOptions{}
		tile := f.GetAxial(s.TilePos)

		size := selectorImgMask.Bounds().Size()

		op.GeoM.Translate(f.Pos.X+tile.CalcPos().X-float64(size.X)/2, f.Pos.Y+tile.CalcPos().Y-float64(size.Y)/2)
		op.GeoM.Concat(*cam.GetMatrix())

		img.DrawImage(selectorImgMask, op)
	}
}
