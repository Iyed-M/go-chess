package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/types"
	"github.com/Iyed-M/go-chess/internal/chess/utils"
)

type Pawn struct {
	*basePiece
}

func (p Pawn) PossibleMoves() []types.Cell {
	c := p.position
	if !utils.CheckInBoardRange(c) {
		panic(fmt.Sprintf("pawn in invalid cell : %s", c.String()))
	}
	moves := make([]types.Cell, 0, 14)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: i, Y: c.Y})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X, Y: i})
	}
	return moves
}

func (p Pawn) Name() Name {
	return NamePawn
}

func newPawn(x, y int8, isWhite bool) Pawn {
	return Pawn{basePiece: NewBasePiece(x, y, isWhite)}
}
