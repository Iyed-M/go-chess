package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type Rook struct {
	*basePiece
}

func (r Rook) PossibleMoves() []cells.Cell {
	c := r.basePiece.position
	if !cells.CheckInBoardRange(c) {
		panic(fmt.Sprintf("rock in invalid cell : %s", c.String()))
	}
	moves := make([]cells.Cell, 0, 28)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: i, Y: c.Y})
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X, Y: i})
	}
	return moves
}

func (r Rook) Name() Name {
	return NameRock
}

func newRook(x, y int8, isWhite bool) Rook {
	return Rook{basePiece: NewBasePiece(x, y, isWhite)}
}
