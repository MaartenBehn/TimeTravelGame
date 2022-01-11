package field

import (
	"github.com/Stroby241/TimeTravelGame/src/util"
	"golang.org/x/image/font"
)

var debugFont font.Face

func loadFont() {
	var err error
	debugFont, err = util.LoadFont(util.FontFaceRegular, 10)
	checkErr(err)
}
