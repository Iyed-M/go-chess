package utils

import (
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

func CheckInBoardRange(c types.Cell) bool {
	return c.X < 8 && c.Y < 8
}

func AppendCellInBoardRange(cells []types.Cell, c types.Cell) ([]types.Cell, bool) {
	if !c.InBoard() {
		return cells, false
	}
	cells = append(cells, c)
	return cells, true
}

// Move(newPosition types.Cell)
//
// Name() name
// IsWhite() bool
