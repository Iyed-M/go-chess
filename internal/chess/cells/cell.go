package cells

import (
	"fmt"
)

type Direction int

func (d Direction) Next(x, y int8) (int8, int8) {
	switch d {
	case DirUp:
		return x, y + 1
	case DirDown:
		return x, y - 1
	case DirLeft:
		return x - 1, y
	case DirRight:
		return x + 1, y
	case DirUpLeft:
		return x - 1, y + 1
	case DirUpRight:
		return x + 1, y + 1
	case DirDownLeft:
		return x - 1, y - 1
	case DirDownRight:
		return x + 1, y - 1
	default:
		panic(fmt.Sprintf("Direction.Next : invalid direction %v", d))
	}
}

const (
	DirNotAligned Direction = iota
	DirUp
	DirDown
	DirLeft
	DirRight
	DirUpLeft
	DirUpRight
	DirDownLeft
	DirDownRight
)

type Cell struct {
	X int8
	Y int8
}

func (c Cell) InBoard() bool {
	return c.X <= 7 && c.Y <= 7 && c.Y >= 0 && c.X >= 0
}

func (c Cell) String() string {
	return fmt.Sprintf("(X=%v,Y=%v)", c.X, c.Y)
}

func (c Cell) MustBeInBoard() {
	if !c.InBoard() {
		panic(fmt.Sprintf("cell not in board %s", c.String()))
	}
}

func (c Cell) Equal(o Cell) bool {
	return c.X == o.X && c.Y == o.Y
}

func (c Cell) Direction(o Cell) Direction {
	if c.X == o.X {
		if c.Y == o.Y {
			return DirNotAligned
		}
		if c.Y < o.Y {
			return DirUp
		}
		return DirDown
	}
	if c.Y == o.Y {
		if c.X < o.X {
			return DirRight
		}
		return DirLeft
	}
	if Abs(c.X-o.X) == Abs(c.Y-o.Y) {
		if c.X < o.X {
			if c.Y < o.Y {
				return DirUpRight
			}
			return DirDownRight
		}
		if c.Y < o.Y {
			return DirUpLeft
		}
		return DirDownLeft
	}
	return DirNotAligned
}
