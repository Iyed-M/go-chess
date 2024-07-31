package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type Pawn struct {
	*basePiece
}

func (p Pawn) PossibleMoves() []cells.Cell {
	c := p.position
	if !cells.CheckInBoardRange(c) {
		panic(fmt.Sprintf("pawn in invalid cell : %s", c.String()))
	}
	moves := make([]cells.Cell, 0, 14)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: i, Y: c.Y})
		moves, _ = cells.AppendCellInBoardRange(moves, cells.Cell{X: c.X, Y: i})
	}
	return moves
}

func (p Pawn) Name() Name {
	return NamePawn
}

func newPawn(x, y int8, isWhite bool) Pawn {
	return Pawn{basePiece: NewBasePiece(x, y, isWhite)}
}
