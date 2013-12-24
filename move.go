package checkerscore

type Move struct {
	from Location
	to   Location
	over Location

	// if a move contains submoves, this means it was a double/triple/etc
	// jump.  the moves (jumps) that compose this double jump will be in order
	// and the submoves completely describe the sequence of move and the "outer move"
	// does not provide any additional information.
	submoves []Move
}

type BoardMove struct {
	board Board
	move  Move
}

func NewMove(moveSequence []Move) Move {
	return Move{}
}

func (move Move) IsJump() bool {
	return (move.from.row-move.to.row == 2 || move.from.row-move.to.row == -2)
}

func (move Move) IsInitialized() bool {
	if move.from.row == 0 &&
		move.from.col == 0 &&
		move.to.row == 0 &&
		move.to.col == 0 {
		return false
	}
	return true
}
