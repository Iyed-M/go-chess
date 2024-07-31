package state

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
	"github.com/Iyed-M/go-chess/internal/chess/pieces"
)

func cellsAligned(c1, c2 cells.Cell) bool {
	return c1.X == c2.X || c1.Y == c2.Y || c1.X-c2.X == c1.Y-c2.Y || c1.X-c2.X == c2.Y-c1.Y
}

// firstPiece returns error if the start and end cells are the same or not aligned
func firstPiece(start, end cells.Cell, board [][]pieces.Piece) (pieces.Piece, error) {
	end.MustBeInBoard()
	start.MustBeInBoard()
	if start.Equal(end) {
		return nil, fmt.Errorf("start %s and end %s are the same", end.String(), start.String())
	}

	dir := start.Direction(end)
	if dir == cells.DirNotAligned {
		return nil, fmt.Errorf("cells %s and %s are not aligned", start.String(), end.String())
	}
	for x, y := dir.Next(start.X, start.Y); x != end.X || y != end.Y; x, y = dir.Next(x, y) {
		if piece := board[x][y]; piece != nil {
			return piece, nil
		}
	}

	return nil, nil
}

func firstBehindTargetFromPov(target, pov cells.Cell, board [][]pieces.Piece) (pieces.Piece, error) {
	target.MustBeInBoard()
	pov.MustBeInBoard()
	if pov.Equal(target) {
		return nil, fmt.Errorf("target %s and pov %s  are the same", target.String(), pov.String())
	}
	dir := pov.Direction(target)
	if dir == cells.DirNotAligned {
		return nil, fmt.Errorf("cells %s and %s are not aligned", target.String(), pov.String())
	}
	for x, y := dir.Next(target.X, target.Y); x >= 0 && y >= 0 && x < 8 && y < 8; x, y = dir.Next(x, y) {
		if piece := board[x][y]; piece != nil {
			return piece, nil
		}
	}

	return nil, nil
}
