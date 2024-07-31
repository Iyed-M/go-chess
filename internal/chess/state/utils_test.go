package state

import (
	"fmt"
	"testing"

	"github.com/Iyed-M/go-chess/internal/chess/cells"
	"github.com/Iyed-M/go-chess/internal/chess/pieces"
)

func Test_firstPiece(t *testing.T) {
	newPieceTest := func(pos cells.Cell) pieces.Piece {
		return _PieceTest{pos: pos}
	}
	newCell := func(x, y int8) cells.Cell {
		return cells.Cell{X: x, Y: y}
	}
	type args struct {
		board [][]pieces.Piece
	}
	tests := []struct {
		name         string
		start        cells.Cell
		end          cells.Cell
		addPieces    []cells.Cell
		wantPieceNil bool
		wantPos      cells.Cell
		wantErr      bool
	}{
		{
			name:      "vertical",
			start:     newCell(3, 3),
			end:       newCell(3, 6),
			addPieces: []cells.Cell{newCell(3, 5), newCell(3, 4)},
			wantPos:   newCell(3, 4),
		},
		{
			name:         "no pieces between",
			start:        newCell(3, 3),
			end:          newCell(3, 6),
			addPieces:    []cells.Cell{},
			wantPieceNil: true,
			wantErr:      false,
		},
		{
			name:      "horizontal",
			start:     newCell(0, 6),
			end:       newCell(7, 6),
			addPieces: []cells.Cell{newCell(3, 6), newCell(5, 4), newCell(4, 6), newCell(7, 7), newCell(7, 6)},
			wantPos:   newCell(3, 6),
			wantErr:   false,
		},
		{
			name:      "up right",
			start:     newCell(1, 3),
			end:       newCell(4, 6),
			addPieces: []cells.Cell{newCell(5, 7), newCell(2, 4), newCell(3, 5), newCell(7, 7), newCell(7, 6)},
			wantPos:   newCell(2, 4),
			wantErr:   false,
		},
		{
			name:      "down right",
			start:     newCell(4, 6),
			end:       newCell(1, 3),
			addPieces: []cells.Cell{newCell(5, 7), newCell(2, 4), newCell(3, 5), newCell(7, 7), newCell(7, 6)},
			wantPos:   newCell(3, 5),
			wantErr:   false,
		},
		{
			name:      "same cell must error",
			start:     newCell(4, 6),
			end:       newCell(4, 6),
			addPieces: []cells.Cell{newCell(5, 7), newCell(2, 4), newCell(3, 5), newCell(7, 7), newCell(7, 6)},
			wantErr:   true,
		},
		{
			name:      "cells not aligned must error",
			start:     newCell(4, 6),
			end:       newCell(1, 0),
			addPieces: []cells.Cell{newCell(5, 7), newCell(2, 4), newCell(3, 5), newCell(7, 7), newCell(7, 6)},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := make([][]pieces.Piece, 8)
			for i := range board {
				board[i] = make([]pieces.Piece, 8)
			}
			for _, c := range tt.addPieces {
				board[c.X][c.Y] = newPieceTest(c)
			}
			board[tt.start.X][tt.start.Y] = newPieceTest(tt.start)
			board[tt.end.X][tt.end.Y] = newPieceTest(tt.end)
			got, err := firstPiece(tt.start, tt.end, board)
			if (err != nil) != tt.wantErr {
				t.Fatalf("firstPiece() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if tt.wantPieceNil {
				if got != nil {
					t.Fatalf("firstPiece() [expected piece to be nil]=%v , [got_piece is null?]=%v", tt.wantPieceNil, got == nil)
				}
				return
			}
			if !got.Position().Equal(tt.wantPos) {
				t.Fatalf("firstPiece() = %v, want %v", got.Position().String(), tt.wantPos)
			}
		})
	}
}

func printBoard(b [][]pieces.Piece) {
	for _, row := range b {
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

type _PieceTest struct {
	pos cells.Cell
}

func _newPieceTest(x, y uint8) _PieceTest {
	return _PieceTest{pos: cells.Cell{X: int8(x), Y: int8(y)}}
}

func (p _PieceTest) Position() cells.Cell {
	return p.pos
}

func (p _PieceTest) PossibleMoves() []cells.Cell {
	return nil
}

func (p _PieceTest) IsWhite() bool {
	return false
}

func (p _PieceTest) Move(newPosition cells.Cell) {
}

func (_PieceTest) Name() pieces.Name {
	return "TP"
}

func FakeStateFromCells(positions [][2]int) *State {
	s := InitState()
	for _, c := range positions {
		s.cacheWhite.pieces = append(s.cacheWhite.pieces, _PieceTest{pos: cells.Cell{X: int8(c[0]), Y: int8(c[1])}})
	}
	return s
}
