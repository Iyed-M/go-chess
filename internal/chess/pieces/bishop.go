package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/types"
	"github.com/Iyed-M/go-chess/internal/chess/utils"
)

type Bishop struct {
	*basePiece
}

func (q Bishop) PossibleMoves() []types.Cell {
	c := q.basePiece.position
	if !utils.CheckInBoardRange(c) {
		panic(fmt.Sprintf("bishop in invalid cell : %s", c.String()))
	}
	moves := make([]types.Cell, 0, 28)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + i, Y: c.Y + i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + i, Y: c.Y - i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - i, Y: c.Y + i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - i, Y: c.Y - i})
	}
	return moves
}

func (q Bishop) Name() Name {
	return NameBishop
}

func newBishop(x, y int8, isWhite bool) Bishop {
	return Bishop{basePiece: NewBasePiece(x, y, isWhite)}
}
