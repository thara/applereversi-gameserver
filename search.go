package applereversi

import ("math")

const (
	minScore = math.MinInt64
	maxScore = math.MaxInt64
)


type Search interface {
	GetBestScore(board *Board, color CellState) int
}

type EvaluationFunc func(*Board, CellState) int

type SearchAlgorithm struct {
	eval EvaluationFunc
	maxDepth int
}

type MiniMaxMethod struct {
	*SearchAlgorithm
}

func (m *MiniMaxMethod) GetBestScore(board *Board, color CellState) int {
	return m.minMax(board, color, 1)
}

func (m *MiniMaxMethod) minMax(node *Board, color CellState, depth int) int {
	if m.maxDepth == depth {
		return m.eval(node, color)
	}
	moves := node.getValidMoves(color)
	if depth % 2 == 0 {
		// opponent turn
		worstScore := maxScore
		for i := range moves {
			mv := moves[i]
			test := node.CloneBoard()
			test.MakeMove(&mv)
			score := m.minMax(test, OppnentColor(color), depth + 1)
			worstScore = min(worstScore, score)
		}
		return worstScore
	} else {
		// self turn
		bestScore := minScore
		for i := range moves{
			mv := moves[i]
			test := node.CloneBoard()
			test.MakeMove(&mv)
			score := m.minMax(test, OppnentColor(color), depth + 1)
			bestScore = max(bestScore, score)
		}
		return bestScore
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}