package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestNewBoard(t *testing.T) {

	boardStr := "" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|" +
		"|- x - x - x - x|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|o - o - o - o -|" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|"

	board := NewBoard(boardStr)
	assert.Equals(t, int(board[0][0]), int(EMPTY)) // TODO: why is cast to int() needed?
	assert.Equals(t, int(board[0][1]), int(RED))
	assert.Equals(t, int(board[7][0]), int(BLACK))
	assert.Equals(t, int(board[7][7]), int(EMPTY))

	// etc...

}

func TestLegalMoves(t *testing.T) {

	// x == red piece
	// X == red king
	// o == black piece
	// O == black king
	// - == unoccupied square (may be legal dark or illegal white)
	currentBoardStr := "" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|" +
		"|- x - x - x - x|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|o - o - o - o -|" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|"
	currentBoardState := NewBoard(currentBoardStr)

	movegen := Movegen{
		board: currentBoardState
	}

	legalMoves := movegen.LegalMoves(RED)
	assert.Equals(t, len(legalMoves), 7)

	/*
		possibleRedMove1 := "" +
			"|- x - x - x - x|" +
			"|x - x - x - x -|" +
			"|- - - x - x - x|" +
			"|x - - - - - - -|" +
			"|- - - - - - - -|" +
			"|o - o - o - o -|" +
			"|- o - o - o - o|" +
			"|o - o - o - o -|"

		possibleRedMove2 := "" +
			"|- x - x - x - x|" +
			"|x - x - x - x -|" +
			"|- x - x - - - x|" +
			"|- - - - x - - -|" +
			"|- - - - - - - -|" +
			"|o - o - o - o -|" +
			"|- o - o - o - o|" +
			"|o - o - o - o -|"

		possibleBoardStates := movegen.LegalMoveBoardStates()
		assert.Equals(t, len(possibleBoardStates), 8) // 8??

		assert.True(t, movegen.IsPossibleBoardstate(possibleRedMove1))
		assert.True(t, movegen.IsPossibleBoardstate(possibleRedMove2))
	*/

}
