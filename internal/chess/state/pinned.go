package state

import (
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

func (s *State) pinSyncerAfterMove(from, to types.Cell) func() error {
	return func() error {
		// owned
		// pin(o, from)
		// unpin(o , to)
		// enemy
		// pin(e, from)
		// unpin(e, from)
		// pin(to, piece)
		// unpin(to, piece)
		return nil
	}
}
