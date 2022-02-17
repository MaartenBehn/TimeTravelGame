package field

import (
	"fmt"
	. "github.com/Stroby241/TimeTravelGame/src/math"
)

type TimePos struct {
	TilePos     AxialPos
	FieldPos    CardPos
	FieldBounds CardPos
}

func (p TimePos) SamePos(p2 TimePos) bool {
	return p.TilePos == p2.TilePos && p.FieldPos == p2.FieldPos
}

func (p TimePos) CalcPos() CardPos {
	return p.CalcTilePos().Add(p.FieldPos.Mul(p.FieldBounds))
}

func (p TimePos) CalcTilePos() CardPos {
	return p.TilePos.MulFloat(tileSize * 2).ToCard().AddFloat(tileSize)
}

func (p TimePos) ToString() string {
	return fmt.Sprintf("Q: %.0f R: %.0f X: %.0f Y: %.0f", p.TilePos.Q, p.TilePos.R, p.FieldPos.X, p.FieldPos.Y)
}
