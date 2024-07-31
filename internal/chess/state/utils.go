package state

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

// error is returned if the cells are not on the same line (horizontal or  vertical or diagonal)
func (s State) GetPiecesBetween(from, to types.Cell) ([]pieces.Piece, error) {
	print("here 1\n")
	if !CellsAligned(from, to) {
		return nil, fmt.Errorf("GetCellBetween : cells [ %s , %s ] are not aligned", from.String(), to.String())
	}

	print("here 2\n")
	dx := to.X - from.X
	dy := to.Y - from.Y

	if dx != 0 {
		dx = dx / abs(dx) // 1 or -1 to get the propeer direction
	}
	if dy != 0 {
		dy = dy / abs(dy)
	}

	ps := []pieces.Piece{}

	print("here 3\n")
	for x, y := from.X+dx, from.Y+dy; x != to.X || y != to.Y; x, y = x+dx, y+dy {
		c := types.Cell{X: x, Y: y}

		print("here 4\n")
		if p := s.white.pieces.FindPieceByPosition(c); p != nil {
			print("here 5\n")
			ps = append(ps, p)
			continue
		}

		if p := s.black.pieces.FindPieceByPosition(c); p != nil {
			print("here 6\n")
			ps = append(ps, p)
		}
	}

	return nil, nil
}

func CellsAligned(c1, c2 types.Cell) bool {
	return c1.X == c2.X || c1.Y == c2.Y || c1.X-c2.X == c1.Y-c2.Y || c1.X-c2.X == c2.Y-c1.Y
}

func abs(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}
