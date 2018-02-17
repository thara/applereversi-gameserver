package applereversi

var cellWeights = [][]int{
	{120, -20, 20, 5, 5, 20, -20, 120},
	{-20, -40, -5, -5, -5, -5, -40, -20},
	{20, -5, 15, 3, 3, 15, -5, 20},
	{5, -5, 3, 3, 3, 3, -5, 5},
	{5, -5, 3, 3, 3, 3, -5, 5},
	{20, -5, 15, 3, 3, 15, -5, 20},
	{-20, -40, -5, -5, -5, -5, -40, -20},
	{120, -20, 20, 5, 5, 20, -20, 120},
}

func GetWeightedScore(board *Board, color CellState) int {
	opponent := OppnentColor(color)

	total := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			s := board.cells[i][j]
			if s == color {
				total += cellWeights[i][j]
			} else if s == opponent {
				total -= cellWeights[i][j]
			}

		}
	}
	return total
}
