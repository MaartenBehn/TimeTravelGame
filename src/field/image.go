package field

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"time"
)

var (
	emptyImage      *ebiten.Image
	selectorImgMask *ebiten.Image
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

	for i, fraction := range Fractions {
		r, g, b, a := fraction.color.RGBA()
		fraction.colorLigth = customColor{r, g, b, a / 2}

		op := &ebiten.DrawImageOptions{
			CompositeMode: ebiten.CompositeModeDestinationIn,
		}

		w, h := unitImgMask.Size()
		unitImg := ebiten.NewImage(w, h)

		w, h = arrowTipImgMask.Size()
		arrowTipImg := ebiten.NewImage(w, h)

		w, h = arrowStraigthImgMask.Size()
		arrowStraigthImg := ebiten.NewImage(w, h)

		w, h = arrowCornerImgMask.Size()
		arrowCornerImg := ebiten.NewImage(w, h)

		w, h = arrowEndImgMask.Size()
		arrowEndImg := ebiten.NewImage(w, h)

		unitImg.Fill(fraction.color)
		unitImg.DrawImage(unitImgMask, op)

		arrowTipImg.Fill(fraction.colorLigth)
		arrowTipImg.DrawImage(arrowTipImgMask, op)

		arrowStraigthImg.Fill(fraction.colorLigth)
		arrowStraigthImg.DrawImage(arrowStraigthImgMask, op)

		arrowCornerImg.Fill(fraction.colorLigth)
		arrowCornerImg.DrawImage(arrowCornerImgMask, op)

		arrowEndImg.Fill(fraction.colorLigth)
		arrowEndImg.DrawImage(arrowEndImgMask, op)

		fraction.images = map[string]*ebiten.Image{}
		fraction.images["unit"] = unitImg
		fraction.images["arrow_tip"] = arrowTipImg
		fraction.images["arrow_straigth"] = arrowStraigthImg
		fraction.images["arrow_corner"] = arrowCornerImg
		fraction.images["arrow_end"] = arrowEndImg

		Fractions[i] = fraction
	}

}
