package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
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

}

func TestKingCanMoveBackwards(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - x - - - - -|" +
		"|- - - O - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - X - - -|" +
		"|- - - o - - - -|" +
		"|- - - - - - - -|"
	board := NewBoard(currentBoardStr)

	// black king (O) move backwards
	start := Location{row: 2, col: 3}
	dest := Location{row: 1, col: 5}
	assert.True(t, board.canMove(BLACK_PLAYER, start, dest))

	// black king (O) jump backwards
	start = Location{row: 2, col: 3}
	intermediate := Location{row: 1, col: 2}
	dest = Location{row: 0, col: 1}
	assert.True(t, board.canJump(BLACK_PLAYER, start, intermediate, dest))

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

func TestApplyMove(t *testing.T) {
	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - - - - - -|" +
		"|X - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"
	board := NewBoard(currentBoardStr)

	from := Location{row: 4, col: 0}
	to := Location{row: 2, col: 2}
	over := Location{row: 3, col: 1}

	move := Move{
		from: from,
		to:   to,
		over: over,
	}
	boardPostMove := board.applyMove(RED_PLAYER, move)
	assert.True(t, boardPostMove.pieceAt(from) == EMPTY)
	assert.True(t, boardPostMove.pieceAt(over) == EMPTY)
	assert.True(t, boardPostMove.pieceAt(to) == board.pieceAt(from))

}

func TestAlternateSingleStepJumpPaths(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|- - X - - - - -|" +
		"|- o - o - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)
	loc := Location{row: 2, col: 2}
	boardMoves := board.alternateSingleStepJumpPaths(RED_PLAYER, loc)
	assert.Equals(t, len(boardMoves), 4)

}

func TestDoubleJumpMovesForLocation(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- o - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|X - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)
	loc := Location{row: 4, col: 0}
	moves := board.legalMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

	for i, move := range moves {
		logg.Log("0>> move %d: %v", i, move)
	}

	expected := []string{
		"{{(4,0)->(4,4)},[{(4,0)->(2,2)},{(2,2)->(4,4)}]}",
		"{{(4,0)->(0,0)},[{(4,0)->(2,2)},{(2,2)->(0,0)}]}",
	}
	assertMovesContains(t, moves, expected)

	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|X - - - - - - -|" +
		"|- o - o - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board = NewBoard(currentBoardStr)
	loc = Location{row: 4, col: 0}
	moves = board.legalMovesForLocation(RED_PLAYER, loc)

	assert.Equals(t, len(moves), 2)

	expected = []string{
		"{{(4,0)->(4,0)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,2)},{(2,2)->(4,0)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(2,2)},{(2,2)->(4,4)},{(4,4)->(6,2)},{(6,2)->(4,0)}]}",
	}
	assertMovesContains(t, moves, expected)

	for i, move := range moves {
		logg.Log("1>> move %d: %v", i, move)
	}

	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - o - o - -|" +
		"|- - - - - - - -|" +
		"|- o - o - o - -|" +
		"|X - - - - - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"
	board = NewBoard(currentBoardStr)
	loc = Location{row: 4, col: 0}
	moves = board.legalMovesForLocation(RED_PLAYER, loc)

	for i, move := range moves {
		logg.Log("2>> move %d: %v", i, move)
	}

	expected = []string{
		"{{(4,0)->(6,6)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(6,6)}]}",
		"{{(4,0)->(6,6)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,6)},{(2,6)->(0,4)},{(0,4)->(2,2)},{(2,2)->(4,4)},{(4,4)->(6,6)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,6)},{(2,6)->(0,4)},{(0,4)->(2,2)},{(2,2)->(4,0)}]}",
		"{{(4,0)->(6,6)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,2)},{(2,2)->(0,4)},{(0,4)->(2,6)},{(2,6)->(4,4)},{(4,4)->(6,6)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,2)},{(2,2)->(4,0)}]}",
		"{{(4,0)->(6,6)},[{(4,0)->(2,2)},{(2,2)->(4,4)},{(4,4)->(6,6)}]}",
		"{{(4,0)->(2,2)},[{(4,0)->(2,2)},{(2,2)->(4,4)},{(4,4)->(2,6)},{(2,6)->(0,4)},{(0,4)->(2,2)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(2,2)},{(2,2)->(4,4)},{(4,4)->(6,2)},{(6,2)->(4,0)}]}",
		"{{(4,0)->(6,6)},[{(4,0)->(2,2)},{(2,2)->(0,4)},{(0,4)->(2,6)},{(2,6)->(4,4)},{(4,4)->(6,6)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(2,2)},{(2,2)->(0,4)},{(0,4)->(2,6)},{(2,6)->(4,4)},{(4,4)->(6,2)},{(6,2)->(4,0)}]}",
		"{{(4,0)->(2,2)},[{(4,0)->(2,2)},{(2,2)->(0,4)},{(0,4)->(2,6)},{(2,6)->(4,4)},{(4,4)->(2,2)}]}",
	}
	assertMovesContains(t, moves, expected)

}

func assertMovesContains(t *testing.T, actualMoves []Move, expectedMoves []string) {
	for _, expectedMove := range expectedMoves {
		found := false
		for _, actualMove := range actualMoves {
			if actualMove.compactString() == expectedMove {
				found = true
			}
		}
		assert.True(t, found)

	}
}

func DISTestRecursiveExplodeJumpMove1(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - o - - - -|" +
		"|- - - - - - - -|" +
		"|- o - - - - - -|" +
		"|X - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)

	from := Location{row: 4, col: 0}
	over := Location{row: 3, col: 1}
	to := Location{row: 2, col: 2}
	startingMove := Move{
		from: from,
		over: over,
		to:   to,
	}

	boardPostMove := board.applyMove(RED_PLAYER, startingMove)
	boardMove := BoardMove{
		board: boardPostMove,
		move:  startingMove,
	}

	boardMoveSeq := make([]BoardMove, 1000)
	boardMoveSeq[0] = boardMove

	boardMoveSequences := make([][]BoardMove, 1000)
	boardMoveSequences[0] = boardMoveSeq

	curBoardMoveSeqIndex := 0
	boardPostMove.recursiveExplodeJumpMove(RED_PLAYER, boardMoveSeq, &curBoardMoveSeqIndex, boardMoveSequences)

	finalBoardMoveSeq := boardMoveSequences[0]
	boardMoveAdded := finalBoardMoveSeq[1]
	assert.Equals(t, boardMoveAdded.move.to.row, 0)
	assert.Equals(t, boardMoveAdded.move.to.col, 4)

	boardMoveNonExistent := finalBoardMoveSeq[2]
	assert.False(t, boardMoveNonExistent.move.IsInitialized())

}

func TestRecursiveExplodeJumpMove2(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - o - -|" +
		"|X - - - - - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)

	from := Location{row: 4, col: 0}
	over := Location{row: 3, col: 1}
	to := Location{row: 2, col: 2}
	startingMove := Move{
		from: from,
		over: over,
		to:   to,
	}

	boardPostMove := board.applyMove(RED_PLAYER, startingMove)
	boardMove := BoardMove{
		board: boardPostMove,
		move:  startingMove,
	}

	boardMoveSeq := make([]BoardMove, 1000)
	boardMoveSeq[0] = boardMove

	boardMoveSequences := make([][]BoardMove, 1000)
	boardMoveSequences[0] = boardMoveSeq

	curBoardMoveSeqIndex := 0
	boardPostMove.recursiveExplodeJumpMove(RED_PLAYER, boardMoveSeq, &curBoardMoveSeqIndex, boardMoveSequences)

	for i, boardMoveSequence := range boardMoveSequences {
		for j, boardMove := range boardMoveSequence {
			if boardMove.move.IsInitialized() {
				logg.Log("i: %d, j: %d move: %v", i, j, boardMove.move)
			}
		}
	}

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
	moves := board.singleJumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 0)

	loc = Location{row: 2, col: 3}
	moves = board.singleJumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move := moves[0]
	assert.Equals(t, move.to.row, 4)
	assert.Equals(t, move.to.col, 5)

	loc = Location{row: 2, col: 5}
	moves = board.singleJumpMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)

	loc = Location{row: 6, col: 3}
	moves = board.singleJumpMovesForLocation(RED_PLAYER, loc)
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
