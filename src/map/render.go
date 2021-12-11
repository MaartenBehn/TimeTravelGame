package gameMap

import (
	. "github.com/Stroby241/TimeTravelGame/src/math"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"math"
	"time"
)

const selectorBlinkIntervall = 1

var (
	emptyImage  *ebiten.Image
	unitImgMask *ebiten.Image

	selectorImgMask *ebiten.Image
	selectorVisable bool
	selectorTime    time.Time

	arrowTipImg      *ebiten.Image
	arrowStraigthImg *ebiten.Image
	arrowCornerImg   *ebiten.Image
	arrowEndImg      *ebiten.Image
)

func Init() {
	emptyImage = ebiten.NewImage(3, 3)
	emptyImage.Fill(color.White)

	var err error
	unitImgMask, _, err = ebitenutil.NewImageFromFile("res/sprites/unit.png")
	checkErr(err)

	selectorImgMask, _, err = ebitenutil.NewImageFromFile("res/sprites/selector.png")
	checkErr(err)
	checkErr(err)
	selectorVisable = false
	selectorTime = time.Now()

	arrowTipImg, _, err = ebitenutil.NewImageFromFile("res/sprites/arrow_tip.png")
	checkErr(err)
	arrowStraigthImg, _, err = ebitenutil.NewImageFromFile("res/sprites/arrow_straigth.png")
	checkErr(err)
	arrowCornerImg, _, err = ebitenutil.NewImageFromFile("res/sprites/arrow_corner.png")
	checkErr(err)
	arrowEndImg, _, err = ebitenutil.NewImageFromFile("res/sprites/arrow_end.png")
	checkErr(err)

}

// draw draws the hexagon for the Tile
func (t Tile) draw(img *ebiten.Image) {
	if !t.Visable {
		return
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe

	img.DrawTriangles(t.vertices, tileIndecies, emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
}

// draw draws the Chunk
func (c *Chunk) draw(img *ebiten.Image) {
	for _, tile := range c.Tiles {
		tile.draw(img)
	}
}

func (u *UnitController) draw(img *ebiten.Image, m *Map) {
	for f, units := range u.Units {
		for _, unit := range units {
			unit.draw(img, u.fractions[f], m)
		}
	}

	for _, units := range u.Units {
		for _, unit := range units {
			if unit.Action.Kind == actionMove || unit.Action.Kind == actionSupport {
				tile, _ := m.GetAxial(unit.Pos)
				totile, _ := m.GetAxial(*unit.Action.ToPos)
				drawArrow(tile.Pos, totile.Pos, img)
			}
		}
	}
}

func (u *Unit) draw(img *ebiten.Image, fraction *Fraction, m *Map) {

	size := unitImgMask.Bounds().Size()

	pointImg := ebiten.NewImage(size.X, size.Y)
	pointImg.Fill(fraction.color)

	op := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeDestinationIn,
	}
	pointImg.DrawImage(unitImgMask, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM = ebiten.GeoM{}
	tile, _ := m.GetAxial(u.Pos)
	op.GeoM.Translate(tile.Pos.X-float64(size.X)/2, tile.Pos.Y-float64(size.Y)/2)
	img.DrawImage(pointImg, op)
}

func (m *Map) Draw(img *ebiten.Image, cam *util.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = *cam.GetMatrix()

	img.DrawImage(m.mapImage, op)

	m.drawSelector(img, cam)
}

func (m *Map) drawSelector(img *ebiten.Image, cam *util.Camera) {
	if time.Since(selectorTime).Seconds() >= selectorBlinkIntervall {
		selectorVisable = !selectorVisable
		selectorTime = time.Now()
	}

	if selectorVisable {
		op := &ebiten.DrawImageOptions{}
		tile, _ := m.GetAxial(m.U.SelectedUnit)

		size := selectorImgMask.Bounds().Size()

		op.GeoM.Translate(tile.Pos.X-float64(size.X)/2, tile.Pos.Y-float64(size.Y)/2)
		op.GeoM.Concat(*cam.GetMatrix())

		img.DrawImage(selectorImgMask, op)
	}
}

const (
	lineNone  = 0
	lineUp    = 1
	lineDown  = 2
	lineRigth = 3
	lineLeft  = 4
)

func drawArrow(fromPos CardPos, toPos CardPos, img *ebiten.Image) {
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
	size := arrowTipImg.Bounds().Size()
	op.GeoM.Translate(-float64(size.X)/2, -float64(size.Y)/2)

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

	img.DrawImage(arrowTipImg, op)

	// Arrow End
	op = &ebiten.DrawImageOptions{}
	size = arrowEndImg.Bounds().Size()
	op.GeoM.Translate(-float64(size.X)/2, -float64(size.Y)/2)

	if firstLine == lineLeft {
		op.GeoM.Scale(-1, 1)
	} else if firstLine == lineUp {
		op.GeoM.Rotate(-math.Pi / 2)
	} else if firstLine == lineDown {
		op.GeoM.Rotate(math.Pi / 2)
	}

	op.GeoM.Translate(fromPos.X, fromPos.Y)

	img.DrawImage(arrowEndImg, op)

	cornerPos := fromPos
	if firstLine == lineRigth || firstLine == lineLeft {
		cornerPos.X += -vec.X
	} else {
		cornerPos.Y += -vec.Y
	}

	// Arrow Corner
	if secondLine != lineNone {
		op = &ebiten.DrawImageOptions{}
		size = arrowCornerImg.Bounds().Size()
		op.GeoM.Translate(-float64(size.X)/2, -float64(size.Y)/2)

		if firstLine == lineLeft {
			op.GeoM.Scale(-1, 1)
		}
		if firstLine == lineDown {
			op.GeoM.Scale(1, -1)
		}
		if secondLine == lineUp {
			op.GeoM.Scale(1, -1)
		}
		if secondLine == lineRigth {
			op.GeoM.Scale(-1, 1)
		}
		op.GeoM.Translate(cornerPos.X, cornerPos.Y)

		img.DrawImage(arrowCornerImg, op)
	}

	// First Line
	// TODO: Uses 10 as Image size change to image size
	op = &ebiten.DrawImageOptions{}

	if firstLine == lineRigth {
		op.GeoM.Scale((vec.X+10)/10, 1)
		op.GeoM.Translate(fromPos.X-vec.X-5, fromPos.Y-5)
	} else if firstLine == lineLeft {
		op.GeoM.Scale((vec.X-10)/10, 1)
		op.GeoM.Translate(fromPos.X-vec.X+5, fromPos.Y-5)
	} else if firstLine == lineDown {
		op.GeoM.Rotate(math.Pi / 2)
		op.GeoM.Scale(1, (vec.Y+10)/10)
		op.GeoM.Translate(fromPos.X+5, fromPos.Y-vec.Y-5)
	} else if firstLine == lineUp {
		op.GeoM.Rotate(math.Pi / 2)
		op.GeoM.Scale(1, (vec.Y-10)/10)
		op.GeoM.Translate(fromPos.X+5, fromPos.Y-vec.Y+5)
	}

	img.DrawImage(arrowStraigthImg, op)

	// Second Line
	if secondLine != lineNone {
		op = &ebiten.DrawImageOptions{}
		if secondLine == lineRigth {
			op.GeoM.Scale((vec.X+10)/10, 1)
			op.GeoM.Translate(cornerPos.X-vec.X-5, cornerPos.Y-5)
		} else if secondLine == lineLeft {
			op.GeoM.Scale((vec.X-10)/10, 1)
			op.GeoM.Translate(cornerPos.X-vec.X+5, cornerPos.Y-5)
		} else if secondLine == lineDown {
			op.GeoM.Rotate(math.Pi / 2)
			op.GeoM.Scale(1, (vec.Y+10)/10)
			op.GeoM.Translate(cornerPos.X+5, cornerPos.Y-vec.Y-5)
		} else if secondLine == lineUp {
			op.GeoM.Rotate(math.Pi / 2)
			op.GeoM.Scale(1, (vec.Y-10)/10)
			op.GeoM.Translate(cornerPos.X+5, cornerPos.Y-vec.Y+5)
		}
	}

	img.DrawImage(arrowStraigthImg, op)
}
