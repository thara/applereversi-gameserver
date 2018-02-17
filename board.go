package applereversi

const (
	boardSize = 8
)

type CellState uint8

const (
	cellStateEmpty CellState = iota
	cellStateBlack
	cellStateWhite
)

func OppnentColor(c CellState) CellState {
	switch c {
	case cellStateBlack:
		return cellStateWhite
	case cellStateWhite:
		return cellStateBlack
	default:
		return c
	}
}

type Line int

const (
	lineBackward Line = iota
	lineHold
	lineForward
)

var allLines = [3]Line{lineBackward, lineHold, lineForward}

type Direction struct {
	vertical   Line
	horizontal Line
}

type Board struct {
	cells [boardSize][boardSize]CellState
}

type BoardCellChanged struct {
	row int
	column int
}

func NewBoard() *Board {
	b := Board{}
	b.cells = [boardSize][boardSize]CellState{}
	for i := range b.cells {
		b.cells[i] = [boardSize]CellState{}
		for j := range b.cells[i] {
			b.cells[i][j] = cellStateEmpty
		}
	}
	b.cells[3][4] = cellStateBlack
	b.cells[4][3] = cellStateBlack
	b.cells[3][3] = cellStateWhite
	b.cells[4][4] = cellStateWhite
	return &b
}

func (b *Board) CloneBoard() *Board {
	b2 := Board{}
	b2.cells = [boardSize][boardSize]CellState{}
	for i := range b.cells {
		copy(b2.cells[i][:], b.cells[i][:])
	}
	return &b2
}

func (b *Board) MakeMove(move *BoardMove) []*BoardCellChanged {
	changes := make([]*BoardCellChanged, 0)
	for i := range allLines {
		v := allLines[i]
		for j := range allLines {
			h := allLines[j]
			if v == lineHold && h == lineHold {
				continue
			}
			d := Direction{vertical: v, horizontal: h}
			n := move.CountFlippableDisks(d, &b.cells)
			if 0 < n {
				y := int(v)
				x := int(h)
				for i := 0; i < n; i++ {
					b.cells[move.row+i*y][move.column+i*x] = move.color
					changes = append(changes, &BoardCellChanged{row: move.row+i*y, column: move.column+i*x})
				}
			}
		}
	}
	b.cells[move.row][move.column] = move.color

	return changes
}

func (b *Board) CountCells(c CellState) int {
	n := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.cells[i][j] == c {
				n++
			}
		}
	}
	return n
}

func (b *Board) HasGameFinished() bool {
	return b.existsValidMove(cellStateBlack) == false && b.existsValidMove(cellStateWhite) == false
}

func (b *Board) existsValidMove(c CellState) bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			m := BoardMove{color: c, row: i, column: j}
			if m.CanPlace(&b.cells) {
				return true
			}
		}
	}
	return false
}

func (b *Board) getValidMoves(c CellState) []BoardMove {
	moves := make([]BoardMove, 0)
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			m := BoardMove{color: c, row: i, column: j}
			if m.CanPlace(&b.cells) {
				moves = append(moves, m)
			}
		}
	}
	return moves
}

func (b *Board) CanPlace(m *BoardMove) bool {
	return m.CanPlace(&b.cells)
}
