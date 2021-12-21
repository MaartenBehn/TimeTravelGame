package field

import . "github.com/Stroby241/TimeTravelGame/src/math"

type TimePos struct {
	TilePos  AxialPos
	FieldPos CardPos
}

func (p TimePos) SamePos(p2 TimePos) bool {
	return p.TilePos == p2.TilePos && p.FieldPos == p2.FieldPos
}

func (p TimePos) CalcPos() CardPos {
	return p.CalcTilePos().Add(p.FieldPos)
}

func (p TimePos) CalcTilePos() CardPos {
	return p.TilePos.MulFloat(tileSize * 2).ToCard().AddFloat(tileSize)
}
