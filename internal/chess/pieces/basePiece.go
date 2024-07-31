package pieces

import "github.com/Iyed-M/go-chess/internal/chess/types"

type basePiece struct {
	position types.Cell
	isWhite  bool
}

func (bp *basePiece) Position() types.Cell {
	return bp.position
}

func NewBasePiece(x, y int8, isWhite bool) *basePiece {
	return &basePiece{position: types.Cell{X: x, Y: y}, isWhite: isWhite}
}

func (bp *basePiece) IsWhite() bool {
	return bp.isWhite
}

func (bp *basePiece) Move(newPosition types.Cell) {
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
