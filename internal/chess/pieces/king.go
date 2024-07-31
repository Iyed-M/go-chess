package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/types"
	"github.com/Iyed-M/go-chess/internal/chess/utils"
)

type King struct {
	*basePiece
}

func (k King) PossibleMoves() []types.Cell {
	c := k.basePiece.position
	if !utils.CheckInBoardRange(c) {
		panic(fmt.Sprintf("bishop in invalid cell : %s", c.String()))
	}
	moves := []types.Cell{}
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + 1, Y: c.Y + 1})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + 1, Y: c.Y})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X + 1, Y: c.Y - 1})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X, Y: c.Y - 1})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X, Y: c.Y + 1})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - 1, Y: c.Y})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - 1, Y: c.Y - 1})
	moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X - 1, Y: c.Y + 1})
	return moves
}

func (k King) Name() Name {
	return NameKing
}

func newKing(x, y int8, isWhite bool) King {
	return King{basePiece: NewBasePiece(x, y, isWhite)}
}
