package pieces

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/types"
	"github.com/Iyed-M/go-chess/internal/chess/utils"
)

type Rook struct {
	*basePiece
}

func (r Rook) PossibleMoves() []types.Cell {
	c := r.basePiece.position
	if !utils.CheckInBoardRange(c) {
		panic(fmt.Sprintf("rock in invalid cell : %s", c.String()))
	}
	moves := make([]types.Cell, 0, 28)
	for idx := range 8 {
		i := int8(idx)
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: i, Y: c.Y})
		moves, _ = utils.AppendCellInBoardRange(moves, types.Cell{X: c.X, Y: i})
	}
	return moves
}

func (r Rook) Name() Name {
	return NameRock
}

func newRook(x, y int8, isWhite bool) Rook {
	return Rook{basePiece: NewBasePiece(x, y, isWhite)}
}
