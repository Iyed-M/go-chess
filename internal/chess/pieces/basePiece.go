package pieces

import "github.com/Iyed-M/go-chess/internal/chess/cells"

type basePiece struct {
	position cells.Cell
	isWhite  bool
}

func (bp *basePiece) Position() cells.Cell {
	return bp.position
}

func NewBasePiece(x, y int8, isWhite bool) *basePiece {
	return &basePiece{position: cells.Cell{X: x, Y: y}, isWhite: isWhite}
}

func (bp *basePiece) IsWhite() bool {
	return bp.isWhite
}

func (bp *basePiece) Move(newPosition cells.Cell) {
	newPosition.MustBeInBoard()
	bp.position = newPosition
}

type Name string

const (
	NameQueen  Name = "queen"
	NameKing   Name = "king"
	NameRock   Name = "rock"
	NameKnight Name = "knight"
	NameBishop Name = "bishop"
	NamePawn   Name = "pawn"
)
