package state

import (
	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type cache struct {
	pieces       pieces.Pieces
	pinnedPieces []pieces.Piece
}

func (p *cache) syncPinnedPiecs(from, cell cells.Cell) {
}
