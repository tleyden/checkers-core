package checkerscore

type EvaluationFunction func(player Player, board Board) float64

func DefaultEvaluationFunction() EvaluationFunction {
	evalFunc := func(player Player, board Board) float64 {
		return board.WeightedScore(player)
	}
	return evalFunc
}
