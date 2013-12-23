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

func TestCanJump(t *testing.T) {

	currentBoardStr := "" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|" +
		"|- o - o - o - o|" +
		"|- - - - x - - -|" +
		"|- - - - - - - -|" +
		"|x - x - - - x -|" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|"
	board := NewBoard(currentBoardStr)

	start := Location{row: 2, col: 1}
	intermediate := Location{row: 3, col: 2}
	dest := Location{row: 4, col: 3}

	assert.False(t, board.canJump(BLACK_PLAYER, start, intermediate, dest))

	start = Location{row: 2, col: 3}
	intermediate = Location{row: 3, col: 4}
	dest = Location{row: 4, col: 5}

	assert.True(t, board.canJump(BLACK_PLAYER, start, intermediate, dest))

}

func DISTEstlegalmoves(t *testing.T) {

	// o == black piece
	// O == black king
	// x == red piece
	// X == red king
	// - == unoccupied square (may be legal dark or illegal white)
	currentBoardStr := "" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|" +
		"|- o - o - o - o|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|x - x - x - x -|" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|"
	board := NewBoard(currentBoardStr)

	legalMoves := board.LegalMoves(BLACK_PLAYER)
	assert.Equals(t, len(legalMoves), 7)

}
