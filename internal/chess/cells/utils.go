package cells

func CheckInBoardRange(c Cell) bool {
	return c.X < 8 && c.Y < 8
}

func AppendCellInBoardRange(cells []Cell, c Cell) ([]Cell, bool) {
	if !c.InBoard() {
		return cells, false
	}
	cells = append(cells, c)
	return cells, true
}

func Abs(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}
