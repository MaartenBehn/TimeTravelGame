package field

import "github.com/Stroby241/TimeTravelGame/src/math"

type MovePattern interface {
	GetPositions(pos TimePos, t *Timeline) []TimePos
}

type BasicMovePattern struct {
	Stride int
}

func (b BasicMovePattern) GetPositions(pos TimePos, t *Timeline) []TimePos {

	var moves = []TimePos{pos}
	for x := -b.Stride; x <= b.Stride; x++ {
		for y := -b.Stride; y <= b.Stride; y++ {

			if x != 0 || y != 0 {
				moves = append(moves, TimePos{
					TilePos:     pos.TilePos,
					FieldPos:    pos.FieldPos.Add(math.CardPos{X: float64(x), Y: float64(y)}),
					FieldBounds: pos.FieldBounds,
				})
			}

			for _, axialPos := range math.AxialDirections {
				moves = append(moves, TimePos{
					TilePos:     pos.TilePos.Add(axialPos),
					FieldPos:    pos.FieldPos.Add(math.CardPos{X: float64(x), Y: float64(y)}),
					FieldBounds: pos.FieldBounds,
				})
			}
		}
	}
	return moves
}
