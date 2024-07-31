package state

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/cells"
)

type color int

const (
	white color = iota
	black
)

type State struct {
	cacheWhite *cache
	cacheBlack *cache

	board     [][]pieces.Piece
	turn      color
	turnCount int
}

func InitState() *State {
	s := &State{
		turn:       white,
		turnCount:  0,
		cacheWhite: &cache{pieces: pieces.InitWhitePieces(), pinnedPieces: []pieces.Piece{}},
		cacheBlack: &cache{pieces: pieces.InitBlackPieces(), pinnedPieces: []pieces.Piece{}},
	}
	board := make([][]pieces.Piece, 8)
	board[0] = s.cacheWhite.pieces[:8]
	board[1] = s.cacheWhite.pieces[8:]
	for i := range 4 {
		board[2+i] = make([]pieces.Piece, 8)
	}
	board[6] = s.cacheBlack.pieces[8:]
	board[7] = s.cacheBlack.pieces[:8]
	s.board = board

	return s
}

func (s *State) PrintBoard() {
	for _, row := range s.board {
		fmt.Printf("|")
		for _, ps := range row {
			if ps == nil {
				fmt.Printf("  |")
			} else {
				fmt.Printf("%s|", ps.Name()[:2])
			}
		}
		fmt.Printf("\n-------------------------\n")
	}
}

func (s *State) allPieces() []pieces.Piece {
	ps := []pieces.Piece{}
	for _, piece := range s.cacheWhite.pieces {
		if piece != nil {
			ps = append(ps, piece)
		}
	}
	for _, piece := range s.cacheBlack.pieces {
		if piece != nil {
			ps = append(ps, piece)
		}
	}
	return ps
}

func (s *State) current() *cache {
	if s.turn == black {
		return s.cacheBlack
	}
	return s.cacheWhite
}

func (s *State) enemy() *cache {
	if s.turn == black {
		return s.cacheWhite
	}
	return s.cacheBlack
}

func (s *State) PossibleMoves(from, to cells.Cell) ([]cells.Cell, error) {
	from.MustBeInBoard()
	if !to.InBoard() {
		return nil, StateError{fmt.Sprintf("State.MovePiece: to cell %s out of board ", to.String())}
	}

	if from.X == to.X && from.Y == to.Y {
		return nil, StateError{fmt.Sprintf("State.MovePiece: piece already in position %s", to.String())}
	}

	piece := s.current().pieces.FindPieceByPosition(from)
	if piece == nil {
		return nil, newStateError(fmt.Sprintf("cell %s does not contain an owned piece", from.String()))
	}
	moves := []cells.Cell{}
	_ = moves
	for _, mov := range piece.PossibleMoves() {
		if s.current().pieces.FindPieceByPosition(mov) != nil {
			continue
		}
		// check pinned
	}

	return nil, nil
}

func (s *State) MovePiece(from, to cells.Cell) error {
	from.MustBeInBoard()
	if !to.InBoard() {
		return StateError{fmt.Sprintf("State.MovePiece: to cell %s out of board ", to.String())}
	}

	if from.X == to.X && from.Y == to.Y {
		return StateError{fmt.Sprintf("State.MovePiece: piece already in position %s", to.String())}
	}
	if s.turn == white {
		err := movePiece(from, to, s.cacheWhite.pieces, s.cacheBlack.pieces)
		return newStateError(err.Error())
	}

	return nil
}

func movePiece(from, to cells.Cell, ownedPieces pieces.Pieces, enemyPieces pieces.Pieces) error {
	return nil
}
