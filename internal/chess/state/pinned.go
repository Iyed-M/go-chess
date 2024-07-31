package state

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

func (s *State) pinSyncerAfterMove(from, to cells.Cell) func() error {
	return func() error {
		// pin(o, from)
		ownedKingPos := s.current().pieces.King().Position()
		if cellsAligned(ownedKingPos, from) {
			p, err := firstPiece(ownedKingPos, from, s.board)
			if err != nil {
				return fmt.Errorf("pinSyncerAfterMove() : %w", err)
			}
			if p != nil {
				// attacker, err := firstPiece(ownedKingPos, end cells.Cell, board [][]pieces.Piece)
			}
		}

		// unpin(o , to)
		//
		// pin(e, from)
		// unpin(e, from)
		// pin(to, piece)
		// unpin(to, piece)
		return nil
	}
}
