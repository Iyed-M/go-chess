package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type Bishop struct {
	*basePiece
}

func (q Bishop) PossibleMoves() []cells.Cell {
	c := q.basePiece.position
	if !cells.CheckInBoardRange(c) {
		panic(fmt.Sprintf("bishop in invalid cell : %s", c.String()))
	}
	moves := make([]cells.Cell, 0, 28)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + i, Y: c.Y + i})
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + i, Y: c.Y - i})
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - i, Y: c.Y + i})
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - i, Y: c.Y - i})
	}
	return moves
}

func (q Bishop) Name() Name {
	return NameBishop
}

func newBishop(x, y int8, isWhite bool) Bishop {
	return Bishop{basePiece: NewBasePiece(x, y, isWhite)}
}
