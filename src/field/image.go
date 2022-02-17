package field

import (
	"encoding/gob"
	"github.com/Stroby241/TimeTravelGame/src/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"time"
)

var (
	emptyImage      *ebiten.Image
	selectorImgMask *ebiten.Image
	tileImgae       *ebiten.Image
)

type customColor struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (c customColor) RGBA() (r, g, b, a uint32) {
	return c.r, c.g, c.b, c.a
}

func Init() {
	gob.Register(BasicMovePattern{})

	loadFont()

	emptyImage = ebiten.NewImage(3, 3)
	emptyImage.Fill(color.White)

	var err error
	unitImgMask, _, err := ebitenutil.NewImageFromFile("res/sprites/unit.png")
	checkErr(err)

	selectorImgMask, _, err = ebitenutil.NewImageFromFile("res/sprites/selector.png")
	checkErr(err)
	checkErr(err)
	selectorVisable = false
	selectorTime = time.Now()

	arrowTipImgMask, _, err := ebitenutil.NewImageFromFile("res/sprites/arrow_tip.png")
	checkErr(err)
	arrowStraigthImgMask, _, err := ebitenutil.NewImageFromFile("res/sprites/arrow_straigth.png")
	checkErr(err)
	arrowCornerImgMask, _, err := ebitenutil.NewImageFromFile("res/sprites/arrow_corner.png")
	checkErr(err)
	arrowEndImgMask, _, err := ebitenutil.NewImageFromFile("res/sprites/arrow_end.png")
	checkErr(err)

	tileImgae, _, err = ebitenutil.NewImageFromFile("res/sprites/tile.png")
	checkErr(err)

	for i, fraction := range util.Fractions {
		r, g, b, a := fraction.Color.RGBA()
		fraction.ColorLigth = customColor{r, g, b, a / 2}

		op := &ebiten.DrawImageOptions{
			CompositeMode: ebiten.CompositeModeDestinationIn,
		}

		w, h := unitImgMask.Size()
		unitImg := ebiten.NewImage(w, h)

		w, h = arrowTipImgMask.Size()
		arrowTipImg := ebiten.NewImage(w, h)
		arrowTipImgL := ebiten.NewImage(w, h)

		w, h = arrowStraigthImgMask.Size()
		arrowStraigthImg := ebiten.NewImage(w, h)
		arrowStraigthImgL := ebiten.NewImage(w, h)

		w, h = arrowCornerImgMask.Size()
		arrowCornerImg := ebiten.NewImage(w, h)
		arrowCornerImgL := ebiten.NewImage(w, h)

		w, h = arrowEndImgMask.Size()
		arrowEndImg := ebiten.NewImage(w, h)
		arrowEndImgL := ebiten.NewImage(w, h)

		unitImg.Fill(fraction.Color)
		unitImg.DrawImage(unitImgMask, op)

		arrowTipImg.Fill(fraction.Color)
		arrowTipImg.DrawImage(arrowTipImgMask, op)
		arrowTipImgL.Fill(fraction.ColorLigth)
		arrowTipImgL.DrawImage(arrowTipImgMask, op)

		arrowStraigthImg.Fill(fraction.Color)
		arrowStraigthImg.DrawImage(arrowStraigthImgMask, op)
		arrowStraigthImgL.Fill(fraction.ColorLigth)
		arrowStraigthImgL.DrawImage(arrowStraigthImgMask, op)

		arrowCornerImg.Fill(fraction.Color)
		arrowCornerImg.DrawImage(arrowCornerImgMask, op)
		arrowCornerImgL.Fill(fraction.ColorLigth)
		arrowCornerImgL.DrawImage(arrowCornerImgMask, op)

		arrowEndImg.Fill(fraction.Color)
		arrowEndImg.DrawImage(arrowEndImgMask, op)
		arrowEndImgL.Fill(fraction.ColorLigth)
		arrowEndImgL.DrawImage(arrowEndImgMask, op)

		fraction.Images = map[string]*ebiten.Image{}
		fraction.Images["unit"] = unitImg
		fraction.Images["arrow_tip"] = arrowTipImg
		fraction.Images["arrow_tip_ligth"] = arrowTipImgL

		fraction.Images["arrow_straigth"] = arrowStraigthImg
		fraction.Images["arrow_straigth_ligth"] = arrowStraigthImgL

		fraction.Images["arrow_corner"] = arrowCornerImg
		fraction.Images["arrow_corner_ligth"] = arrowCornerImgL

		fraction.Images["arrow_end"] = arrowEndImg
		fraction.Images["arrow_end_ligth"] = arrowEndImgL

		util.Fractions[i] = fraction
	}
}
