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
		"|- o - - - o - o|" +
		"|o - x - o - o -|" +
		"|- o - - - o - o|" +
		"|- - - - X - - -|" +
		"|- - - - - - - -|" +
		"|x - o - o - x -|" +
		"|- x - x - o - x|" +
		"|x - x - - - x -|"
	board := NewBoard(currentBoardStr)

	start := Location{row: 5, col: 0}
	intermediate := Location{row: 4, col: 1}
	dest := Location{row: 3, col: 2}

	assert.False(t, board.canJump(RED_PLAYER, start, intermediate, dest))

	start = Location{row: 2, col: 3}
	intermediate = Location{row: 3, col: 4}
	dest = Location{row: 4, col: 5}

	assert.True(t, board.canJump(BLACK_PLAYER, start, intermediate, dest))

	start = Location{row: 6, col: 3}
	intermediate = Location{row: 5, col: 2}
	dest = Location{row: 4, col: 1}

	assert.True(t, board.canJump(RED_PLAYER, start, intermediate, dest))

	start = Location{row: 6, col: 3}
	intermediate = Location{row: 5, col: 4}
	dest = Location{row: 4, col: 5}

	assert.True(t, board.canJump(RED_PLAYER, start, intermediate, dest))

	// red piece trying to jump backwards
	start = Location{row: 5, col: 6}
	intermediate = Location{row: 6, col: 5}
	dest = Location{row: 7, col: 4}
	assert.False(t, board.canJump(RED_PLAYER, start, intermediate, dest))

	// black piece trying to jump backwards
	start = Location{row: 2, col: 1}
	intermediate = Location{row: 1, col: 2}
	dest = Location{row: 0, col: 3}
	assert.False(t, board.canJump(BLACK_PLAYER, start, intermediate, dest))

}

func TestCanMove(t *testing.T) {

	currentBoardStr := "" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|" +
		"|- - - o - O - o|" +
		"|- - - - x - - -|" +
		"|- - - - - - - -|" +
		"|x - x - o - x -|" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|"
	board := NewBoard(currentBoardStr)

	start := Location{row: 5, col: 0}
	dest := Location{row: 4, col: 1}
	assert.True(t, board.canMove(RED_PLAYER, start, dest))

	start = Location{row: 5, col: 0}
	dest = Location{row: 2, col: 5}
	assert.False(t, board.canMove(RED_PLAYER, start, dest))

	start = Location{row: 5, col: 4}
	dest = Location{row: 4, col: 3}
	assert.False(t, board.canMove(BLACK_PLAYER, start, dest))

	start = Location{row: 3, col: 4}
	dest = Location{row: 5, col: 3}
	assert.False(t, board.canMove(RED_PLAYER, start, dest))

}

func TestJumpMovesForLocation(t *testing.T) {

	currentBoardStr := "" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|" +
		"|- - - o - o - -|" +
		"|- - - - x - - -|" +
		"|- - - - - - - -|" +
		"|x - o - o - x -|" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|"
	board := NewBoard(currentBoardStr)

	loc := Location{row: 3, col: 0}
	moves := board.jumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 0)

	loc = Location{row: 2, col: 3}
	moves = board.jumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move := moves[0]
	assert.Equals(t, move.to.row, 4)
	assert.Equals(t, move.to.col, 5)

	loc = Location{row: 2, col: 5}
	moves = board.jumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)

	loc = Location{row: 6, col: 3}
	moves = board.jumpMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

}

func TestNonJumpMovesForLocation(t *testing.T) {

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

	loc := Location{row: 5, col: 0}
	moves := board.nonJumpMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 1)

	loc = Location{row: 5, col: 2}
	moves = board.nonJumpMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

	loc = Location{row: 7, col: 7}
	moves = board.nonJumpMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 0)

}

func TestLegalMovesForLocation(t *testing.T) {

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

	loc := Location{row: 2, col: 3}
	moves := board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move := moves[0]
	assert.Equals(t, move.to.row, 4)
	assert.Equals(t, move.to.col, 5)

	loc = Location{row: 2, col: 1}
	moves = board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

	loc = Location{row: 5, col: 0}
	moves = board.legalMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 1)

}

func TestLegalMoves(t *testing.T) {

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

	legalMoves = board.LegalMoves(RED_PLAYER)
	assert.Equals(t, len(legalMoves), 7)

}
