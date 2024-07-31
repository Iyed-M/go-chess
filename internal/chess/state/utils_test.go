package state

import (
	"slices"
	"testing"

	"github.com/Iyed-M/go-chess/internal/chess/pieces"
	"github.com/Iyed-M/go-chess/internal/chess/types"
)

type _PieceTest struct {
	pos types.Cell
}

func (p _PieceTest) Position() types.Cell {
	return p.pos
}

func (p _PieceTest) PossibleMoves() []types.Cell {
	return nil
}

func (p _PieceTest) IsWhite() bool {
	return false
}

func (_PieceTest) Move(newPosition types.Cell) {
}

func (_PieceTest) Name() (_ pieces.Name) {
	panic("implement me")
}

func FakeStateFromCells(cells [][2]int) *State {
	s := initState()
	for _, c := range cells {
		s.white.pieces = append(s.white.pieces, _PieceTest{pos: types.Cell{X: int8(c[0]), Y: int8(c[1])}})
	}
	return s
}

func TestState_GetPiecesBetween(t *testing.T) {
	FakeStateFromCells([][2]int{{1, 2}, {}})
	type args struct {
		x1 int8
		y1 int8
		x2 int8
		y2 int8
	}
	tests := []struct {
		name    string
		s       *State
		args    args
		want    []types.Cell
		wantErr bool
	}{
		{
			s:       FakeStateFromCells([][2]int{{1, 2}}),
			wantErr: false,
			name:    "horizontal 0",
			args:    args{x1: 1, y1: 1, x2: 1, y2: 3},
			want:    []types.Cell{{X: 1, Y: 2}},
		},
		{
			s:       FakeStateFromCells([][2]int{}),
			wantErr: false,
			name:    "horizontal 1",
			args:    args{x1: 1, y1: 1, x2: 1, y2: 3},
			want:    []types.Cell{},
		},
		{
			s:       FakeStateFromCells([][2]int{{5, 7}, {6, 7}}),
			wantErr: false,
			name:    "vertical",
			args:    args{x1: 7, y1: 7, x2: 4, y2: 7},
			want:    []types.Cell{{X: 5, Y: 7}, {X: 6, Y: 7}},
		},
		{
			wantErr: false,
			name:    "diag",
			args:    args{x1: 1, y1: 2, x2: 4, y2: 5},
			want:    []types.Cell{{X: 2, Y: 3}, {X: 3, Y: 4}},
		},
		{
			wantErr: true,
			name:    "non aligned",
			args:    args{x1: 1, y1: 2, x2: 4, y2: 3},
			want:    []types.Cell{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPiecesBetween(types.Cell{X: tt.args.x1, Y: tt.args.y1}, types.Cell{X: tt.args.x2, Y: tt.args.y2})
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCellBetween() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetCellBetween() = %v, want %v", got, tt.want)
				return
			}
			for _, c := range got {
				if !slices.Contains(tt.want, c.Position()) {
					t.Errorf("GetCellBetween() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
