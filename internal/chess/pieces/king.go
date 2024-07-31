package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type King struct {
	*basePiece
}

func (k King) PossibleMoves() []cells.Cell {
	c := k.basePiece.position
	if !cells.CheckInBoardRange(c) {
		panic(fmt.Sprintf("bishop in invalid cell : %s", c.String()))
	}
	moves := []cells.Cell{}
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 1, Y: c.Y + 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 1, Y: c.Y})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 1, Y: c.Y - 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X, Y: c.Y - 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X, Y: c.Y + 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 1, Y: c.Y})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 1, Y: c.Y - 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 1, Y: c.Y + 1})
	return moves
}

func (k King) Name() Name {
	return NameKing
}

func newKing(x, y int8, isWhite bool) King {
	return King{basePiece: NewBasePiece(x, y, isWhite)}
}
