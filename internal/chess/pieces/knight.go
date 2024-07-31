package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type Knight struct {
	*basePiece
}

func (kn Knight) PossibleMoves() []cells.Cell {
	c := kn.basePiece.position
	if !cells.CheckInBoardRange(c) {
		panic(fmt.Sprintf("bishop in invalid cell : %s", c.String()))
	}
	moves := []cells.Cell{}
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 1, Y: c.Y + 2})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 1, Y: c.Y - 2})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 2, Y: c.Y + 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X + 2, Y: c.Y - 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 1, Y: c.Y + 2})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 1, Y: c.Y - 2})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 2, Y: c.Y + 1})
	moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X - 2, Y: c.Y - 1})
	return moves
}

func (kn Knight) Name() Name {
	return NameKnight
}

func newKnight(x, y int8, isWhite bool) Knight {
	return Knight{basePiece: NewBasePiece(x, y, isWhite)}
}
