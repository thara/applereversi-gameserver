package applereversi

type BoardMove struct {
	color CellState
	row int
	column int
}

func (m *BoardMove) CountFlippableDisks(direction Direction, cells *[boardSize][boardSize]CellState) int {
	y := int(direction.vertical)
	x := int(direction.horizontal)
	opponent := OppnentColor(m.color)
	count := 1

	for (m.row + count * y) < boardSize && (m.column + count * x) < boardSize && cells[m.row + count * y][m.column + count * x] == opponent {
		count++
	}
	if cells[m.row + count * y][m.column + count * x] == m.color {
		return count - 1
	} else {
		return 0
	}
}

func (m *BoardMove) CanPlace(cells *[boardSize][boardSize]CellState) bool {
	c := cells[m.row][m.column]
	if c != cellStateEmpty {
		return false
	}

	for i := range allLines {
		v := allLines[i]
		for j := range allLines {
			h := allLines[j]
			if v == lineHold && h == lineHold {
				continue
			}
			if 0 < m.CountFlippableDisks(Direction{vertical: v, horizontal: h}, cells) {
				return true
			}
		}
	}

	return false
}