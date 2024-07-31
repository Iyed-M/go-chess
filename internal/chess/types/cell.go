package types

import (
	"fmt"
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
