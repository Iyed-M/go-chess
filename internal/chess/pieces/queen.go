package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/types"
	"github.com/Iyed-M/go-chess/internal/chess/utils"
)

type Queen struct {
	*basePiece
}

func (q Queen) PossibleMoves() []types.Cell {
	c := q.position
	if !utils.CheckInBoardRange(c) {
		panic(fmt.Sprintf("quuen in invalid cell : %s", c.String()))
	}
	moves := make([]types.Cell, 0, 28)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X, Y: i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: i, Y: c.Y})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + i, Y: c.Y + i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + i, Y: c.Y - i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - i, Y: c.Y + i})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - i, Y: c.Y - i})
	}
	return moves
}

func (q Queen) Name() Name {
	return NameQueen
}

func newQueen(x, y int8, isWhite bool) Queen {
	return Queen{basePiece: NewBasePiece(x, y, isWhite)}
}
