package checkerscore

type Movegen struct{}

func (m Movegen) SetCurrentBoardState(board Board) {

}

func (m Movegen) LegalMoves(player Player) []Move {
	return []Move{}
}
