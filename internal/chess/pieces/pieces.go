package pieces

import (
	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type Piece interface {
	Position() cells.Cell
	Move(newPosition cells.Cell)
	PossibleMoves() []cells.Cell
	Name() Name
	IsWhite() bool
}

func InitWhitePieces() []Piece {
	return []Piece{
		newRook(0, 0, false),
		newKnight(1, 0, false),
		newBishop(2, 0, false),
		newQueen(3, 0, false),
		newKing(4, 0, false),
		newBishop(5, 0, false),
		newKnight(6, 0, false),
		newRook(7, 0, false),
		newPawn(0, 1, false),
		newPawn(1, 1, false),
		newPawn(2, 1, false),
		newPawn(3, 1, false),
		newPawn(4, 1, false),
		newPawn(5, 1, false),
		newPawn(6, 1, false),
		newPawn(7, 1, false),
	}
}

type Pieces []Piece

func (ps Pieces) Rook1() Rook {
	return ps[0].(Rook)
}

func (ps Pieces) Knight1() Knight {
	return ps[1].(Knight)
}

func (ps Pieces) Bishop1() Bishop {
	return ps[2].(Bishop)
}

func (ps Pieces) Queen() Queen {
	return ps[3].(Queen)
}

func (ps Pieces) King() King {
	return ps[4].(King)
}

func (ps Pieces) Bishop2() Bishop {
	return ps[5].(Bishop)
}

func (ps Pieces) Knight2() Knight {
	return ps[6].(Knight)
}

func (ps Pieces) Rook2() Rook {
	return ps[7].(Rook)
}

func (ps Pieces) Pawns() []Piece {
	return ps[8:]
}

func (ps Pieces) FindPieceByPosition(pos cells.Cell) Piece {
	for _, p := range ps {
		if p.Position() == pos {
			return p
		}
	}
	return nil
}

func InitBlackPieces() Pieces {
	return []Piece{
		newRook(0, 7, true),
		newKnight(1, 7, true),
		newBishop(2, 7, true),
		newQueen(3, 7, true),
		newKing(4, 7, true),
		newBishop(5, 7, true),
		newKnight(6, 7, true),
		newRook(7, 7, true),
		newPawn(0, 6, true),
		newPawn(1, 6, true),
		newPawn(2, 6, true),
		newPawn(3, 6, true),
		newPawn(4, 6, true),
		newPawn(5, 6, true),
		newPawn(6, 6, true),
		newPawn(7, 6, true),
	}
}
