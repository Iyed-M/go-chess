package state

import (
	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

type player struct {
	pieces       pieces.Pieces
	pinnedPieces []pieces.Piece
}

func (p *player) syncPinnedPiecs(from, cell types.Cell) {
}
