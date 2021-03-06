package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
	_ "github.com/couchbaselabs/logg"
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
	assert.Equals(t, int(board[1][0]), int(RED))
	assert.Equals(t, int(board[7][0]), int(BLACK))
	assert.Equals(t, int(board[7][7]), int(EMPTY))

}

func TestIsOnOpponentsFirstRank(t *testing.T) {
	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - x - - - - -|" +
		"|- - - O - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - X - - -|" +
		"|- - - o - - - -|" +
		"|- o - - - - - -|"
	board := NewBoard(currentBoardStr)
	shouldBeKing := Location{row: 7, col: 1}
	nonKingPiece := Location{row: 6, col: 3}
	assert.False(t, board.isOnOpponentsFirstRank(nonKingPiece, BLACK_PLAYER))
	assert.True(t, board.isOnOpponentsFirstRank(shouldBeKing, BLACK_PLAYER))

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
	boardPostMove := board.ApplyMove(RED_PLAYER, move)
	assert.True(t, boardPostMove.pieceAt(from) == EMPTY)
	assert.True(t, boardPostMove.pieceAt(over) == EMPTY)
	assert.True(t, boardPostMove.pieceAt(to) == board.pieceAt(from))

	// make sure when black moves it becomes a king
	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- x - - - - - -|" +
		"|- - - - - o - -|" +
		"|- - - - - - - -|"
	board = NewBoard(currentBoardStr)
	move = Move{
		from: Location{row: 6, col: 5},
		to:   Location{row: 7, col: 4},
	}
	boardPostMove = board.ApplyMove(BLACK_PLAYER, move)
	expectedBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- x - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - O - - -|"
	assert.Equals(t, boardPostMove.CompactString(false), expectedBoardStr)

}

func TestLegalMovesKingedWhileJumping(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - o|" +
		"|x - - - - - o -|" +
		"|- - - X - o - -|" +
		"|- - - - - - o -|" +
		"|- - - - - x - -|" +
		"|- - - - - - x -|" +
		"|- x - x - x - x|" +
		"|- - - - - - x -|"
	board := NewBoard(currentBoardStr)
	loc := Location{row: 3, col: 6}
	moves, _ := board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move := moves[0]

	logg.Log("move: %v", move)

	// if its king'ed, it can't keep jumping backwards!
	assert.False(t, move.to == Location{row: 5, col: 0})
	assert.True(t, move.to == Location{row: 7, col: 2})

	currentBoardStr = "" +
		"|- o - o - o - o|" +
		"|- - o - - - - -|" +
		"|- o - o - - - -|" +
		"|- - - - - - - -|" +
		"|- x - - - - - -|" +
		"|- - x - o - - -|" +
		"|- x - x - - - -|" +
		"|x - - - - - x -|"

	board = NewBoard(currentBoardStr)
	loc = Location{row: 5, col: 4}
	moves, _ = board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move = moves[0]

	// if its king'ed, it can't keep jumping backwards!
	assert.False(t, move.to == Location{row: 3, col: 2})
	assert.True(t, move.to == Location{row: 7, col: 2})

}

func TestApplyMoveWithSubmoves(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|X - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)

	loc := Location{row: 4, col: 0}
	moves, _ := board.legalMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 1)

	move := moves[0]

	boardPostMove := board.ApplyMove(RED_PLAYER, move)

	expectedBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - X - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	assert.Equals(t, expectedBoardStr, boardPostMove.CompactString(false))

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

func TestDoubleJumpMovesForLocationEasy(t *testing.T) {

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
	moves, _ := board.legalMovesForLocation(RED_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

	expected := []string{
		"{{(4,0)->(4,4)},[{(4,0)->(2,2)},{(2,2)->(4,4)}]}",
		"{{(4,0)->(0,0)},[{(4,0)->(2,2)},{(2,2)->(0,0)}]}",
	}
	assertMovesContains(t, moves, expected)

}

func TestDoubleJumpMovesForLocationMedium(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|X - - - - - - -|" +
		"|- o - o - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)
	loc := Location{row: 4, col: 0}
	moves, _ := board.legalMovesForLocation(RED_PLAYER, loc)

	assert.Equals(t, len(moves), 2)

	expected := []string{
		"{{(4,0)->(4,0)},[{(4,0)->(6,2)},{(6,2)->(4,4)},{(4,4)->(2,2)},{(2,2)->(4,0)}]}",
		"{{(4,0)->(4,0)},[{(4,0)->(2,2)},{(2,2)->(4,4)},{(4,4)->(6,2)},{(6,2)->(4,0)}]}",
	}
	assertMovesContains(t, moves, expected)

}

func TestDoubleJumpMovesForLocationHard(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - o - o - -|" +
		"|- - - - - - - -|" +
		"|- o - o - o - -|" +
		"|X - - - - - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"
	board := NewBoard(currentBoardStr)
	loc := Location{row: 4, col: 0}
	moves, _ := board.legalMovesForLocation(RED_PLAYER, loc)

	expected := []string{
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

func TestMinimax(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - o -|" +
		"|- - - - - - - -|" +
		"|- - - o - - - -|" +
		"|- - x - x - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - x -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	evalFunc := DefaultEvaluationFunction()
	board := NewBoard(currentBoardStr)
	depth := 0
	_, blackScore := board.Minimax(BLACK_PLAYER, depth, evalFunc)
	_, redScore := board.Minimax(RED_PLAYER, depth, evalFunc)
	assert.True(t, redScore > blackScore)
	logg.Log("blackScore: %v, red: %v", blackScore, redScore)

	// after taking a double jump, black will have 2 pieces
	// to red's 1 piece, and so the min score should be 1.0
	depth = 1
	bestMove, blackScorePost1Move := board.Minimax(BLACK_PLAYER, depth, evalFunc)
	logg.Log("blackScore: %v", blackScorePost1Move)
	logg.Log("bestMove: %v", bestMove)
	assert.Equals(t, blackScorePost1Move, 1.0)
	assert.Equals(t, bestMove.To(), Location{row: 6, col: 7})

	// the black piece is trapped and so if the red piece moves
	// forward one piece, it will capture it when the black
	// piece has to move forward
	currentBoardStr = "" +
		"|- - - - - - - o|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - x - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"
	board = NewBoard(currentBoardStr)
	depth = 3
	bestMove, scorePostMove := board.Minimax(RED_PLAYER, depth, evalFunc)
	logg.Log("score: %v", scorePostMove)
	logg.Log("bestMove: %v", bestMove)
	assert.True(t, scorePostMove >= 1.0)
	assert.Equals(t, bestMove.To(), Location{row: 2, col: 5})

	// black has lost the game, so when getting the minimax for that
	// move it will return final score for black
	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - x - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"
	board = NewBoard(currentBoardStr)
	depth = 1
	bestMove2, scorePostMove2 := board.Minimax(BLACK_PLAYER, depth, evalFunc)
	assert.True(t, scorePostMove2 == -1.0)
	assert.False(t, bestMove2.IsInitialized())

	// http://www.usacheckers.com/famouspositions.php - Third Position
	/*
		currentBoardStr = "" +
			"|- - - - - - - -|" +
			"|o - - - - - - -|" +
			"|- - - - - - - -|" +
			"|o - x - - - - -|" +
			"|- - - - - - - -|" +
			"|o - x - - - - -|" +
			"|- - - - - - - -|" +
			"|- - - - - - - -|"
		board = NewBoard(currentBoardStr)
		depth = 7
		bestMove3, scorePostMove3 := board.Minimax(BLACK_PLAYER, depth, evalFunc)
		logg.Log("score: %v", scorePostMove3)
		logg.Log("bestMove: %v", bestMove3)
	*/

}

func TestWeightedScore(t *testing.T) {

	currentBoardStr := "" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- o - o - - - -|" +
		"|X - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	board := NewBoard(currentBoardStr)
	blackScore := board.WeightedScore(BLACK_PLAYER)
	redScore := board.WeightedScore(RED_PLAYER)
	// logg.Log("blackScore: %v, red: %v", blackScore, redScore)
	assert.True(t, blackScore > redScore)

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
	moves, _ := board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 1)
	move := moves[0]
	assert.Equals(t, move.to.row, 4)
	assert.Equals(t, move.to.col, 5)

	loc = Location{row: 2, col: 1}
	moves, _ = board.legalMovesForLocation(BLACK_PLAYER, loc)
	assert.Equals(t, len(moves), 2)

	loc = Location{row: 5, col: 0}
	moves, _ = board.legalMovesForLocation(RED_PLAYER, loc)
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

	currentBoardStr = "" +
		"|- o - o - o - o|" +
		"|- - o - o - o -|" +
		"|- o - o - o - o|" +
		"|- - o - x - - -|" +
		"|- - - - - - - -|" +
		"|x - - - x - x -|" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|"
	board = NewBoard(currentBoardStr)

	legalMoves = board.LegalMoves(BLACK_PLAYER)
	logg.Log("legalMoves: %v", legalMoves)

	illegalMove := Move{
		from: Location{row: 0, col: 1},
		to:   Location{row: 1, col: 0},
	}

	assert.False(t, illegalMove.ContainedIn(legalMoves))

}
