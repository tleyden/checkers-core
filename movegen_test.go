package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

/*
This is a move generator.

Given a board state and a player, this generates all the possible moves that player can make.

*/

func TestMoveGen(t *testing.T) {

	// x == red piece
	// X == red king
	// o == blue piece
	// O == blue king
	// . == white square
	// - == unoccupied dark square
	currentBoardStr := "" +
		"|. x . x . x . x|" +
		"|x . x . x . x .|" +
		"|. x . x . x . x|" +
		"|- . - . - . - .|" +
		"|. - . - . - . -|" +
		"|o . o . o . o .|" +
		"|. o . o . o . o|" +
		"|o . o . o . o .|"
	currentBoardState := NewBoard(currentBoardStr)

	movegen := Movegen{}
	movegen.SetCurrentBoardState(currentBoardState)

	legalMoves := movegen.LegalMoves(RED)
	assert.Equals(t, len(legalMoves), 8) // 8??

	/*
		possibleRedMove1 := "" +
			"|. x . x . x . x|" +
			"|x . x . x . x .|" +
			"|. - . x . x . x|" +
			"|x . - . - . - .|" +
			"|. - . - . - . -|" +
			"|o . o . o . o .|" +
			"|. o . o . o . o|" +
			"|o . o . o . o .|"

		possibleRedMove2 := "" +
			"|. x . x . x . x|" +
			"|x . x . x . x .|" +
			"|. x . x . - . x|" +
			"|- . - . x . - .|" +
			"|. - . - . - . -|" +
			"|o . o . o . o .|" +
			"|. o . o . o . o|" +
			"|o . o . o . o .|"

		possibleBoardStates := movegen.LegalMoveBoardStates()
		assert.Equals(t, len(possibleBoardStates), 8) // 8??

		assert.True(t, movegen.IsPossibleBoardstate(possibleRedMove1))
		assert.True(t, movegen.IsPossibleBoardstate(possibleRedMove2))
	*/

}
