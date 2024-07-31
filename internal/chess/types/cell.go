package types

import "fmt"

type Cell struct {
	X uint8
	Y uint8
}

func (c Cell) InBoard() bool {
	return c.X <= 7 && c.Y <= 7
}

func (c Cell) String() string {
	return fmt.Sprintf("(X=%v,Y=%v)", c.X, c.Y)
}

func (c Cell) MustBeInBoard() {
	if !c.InBoard() {
		panic(fmt.Sprintf("cell not in board %s", c.String()))
	}
}
