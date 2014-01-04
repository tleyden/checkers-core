package checkerscore

import (
	"bytes"
	"fmt"
)

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

type MoveFilter func(move Move) bool

func NewMoveFromTo(from, to Location) Move {
	return Move{
		from: from,
		to:   to,
	}
}

func NewMove(moveSequence []Move) Move {

	if len(moveSequence) == 0 {
		panic("move sequence is empty")
	}

	resultMove := Move{}
	resultMove.submoves = []Move{}

	// use the first move in the sequence for the "from" location
	firstMove := moveSequence[0]
	resultMove.from = firstMove.from

	for _, move := range moveSequence {
		resultMove.to = move.to
		resultMove.submoves = append(resultMove.submoves, move)
	}

	return resultMove

}

func (move Move) IsJump() bool {
	if len(move.submoves) > 0 {
		return true
	}
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

/*
Serialize a move into a compact string representation:

    {(0,0)->(2,2)}

or in the case of multiple jumps:

    {{(4,0)->(0,0)},[{(4,0)->(2,2)},{(2,2)->(0,0)}]}

*/
func (move Move) compactString() string {

	buffer := bytes.Buffer{}

	if len(move.submoves) > 0 {
		buffer.WriteString("{")
		buffer.WriteString(move.compactStringWithoutSubmoves())
		buffer.WriteString(",[")
		for i, submove := range move.submoves {
			buffer.WriteString(submove.compactStringWithoutSubmoves())
			if i != (len(move.submoves) - 1) {
				buffer.WriteString(",")
			}
		}
		buffer.WriteString("]}")

	} else {
		buffer.WriteString(move.compactStringWithoutSubmoves())
	}

	return buffer.String()
}

func (move Move) compactStringWithoutSubmoves() string {
	from := fmt.Sprintf("(%d,%d)", move.from.row, move.from.col)
	to := fmt.Sprintf("(%d,%d)", move.to.row, move.to.col)
	from_to := fmt.Sprintf("{%v->%v}", from, to)
	return from_to
}

func (move Move) From() Location {
	return move.from
}

func (move Move) To() Location {
	return move.to
}

func (move Move) String() string {
	return move.compactString()
}

func (move Move) ContainedIn(moves []Move) bool {
	for _, curMove := range moves {
		if move.Equals(curMove) {
			return true
		}
	}
	return false
}

func (move Move) Equals(otherMove Move) bool {

	// TODO: fix this!! it's totally broken, because
	// TODO: moves can end up at the same place but
	// TODO: have different paths (eg, different set of submoves)
	// TODO: and should not be equal.
	// TODO: first, write test that proves this is broken.  then fix it.

	return move.To().Equals(otherMove.To()) &&
		move.From().Equals(otherMove.From())

}

func filterMoves(moves []Move, filter MoveFilter) []Move {
	filteredMoves := []Move{}
	for _, move := range moves {
		if filter(move) {
			filteredMoves = append(filteredMoves, move)
		}
	}
	return filteredMoves
}
