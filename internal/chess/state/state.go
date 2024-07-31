package state

import (
	"fmt"

	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

type color int

const (
	white color = iota
	black
)

type State struct {
	white *player
	black *player

	turn      color
	turnCount int
}

func initState() *State {
	return &State{
		turn:      white,
		turnCount: 0,
		white:     &player{pieces: pieces.InitWhitePieces()},
		black:     &player{pieces: pieces.InitBlackPieces()},
	}
}

func (s *State) allPieces() []pieces.Piece {
	ps := []pieces.Piece{}
	for _, piece := range s.white.pieces {
		if piece != nil {
			ps = append(ps, piece)
		}
	}
	for _, piece := range s.black.pieces {
		if piece != nil {
			ps = append(ps, piece)
		}
	}
	return ps
}

func (s *State) current() *player {
	if s.turn == black {
		return s.black
	}
	return s.white
}

func (s *State) enemy() *player {
	if s.turn == black {
		return s.white
	}
	return s.black
}

func (s *State) PossibleMoves(from, to types.Cell) ([]types.Cell, error) {
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
	moves := []types.Cell{}
	_ = moves
	for _, mov := range piece.PossibleMoves() {
		if s.current().pieces.FindPieceByPosition(mov) != nil {
			continue
		}
		// check pinned
	}

	return nil, nil
}

func (s *State) MovePiece(from, to types.Cell) error {
	from.MustBeInBoard()
	if !to.InBoard() {
		return StateError{fmt.Sprintf("State.MovePiece: to cell %s out of board ", to.String())}
	}

	if from.X == to.X && from.Y == to.Y {
		return StateError{fmt.Sprintf("State.MovePiece: piece already in position %s", to.String())}
	}
	if s.turn == white {
		err := movePiece(from, to, s.white.pieces, s.black.pieces)
		return newStateError(err.Error())
	}

	return nil
}

func movePiece(from, to types.Cell, ownedPieces pieces.Pieces, enemyPieces pieces.Pieces) error {
	return nil
}
