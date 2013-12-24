package checkerscore

import (
	"github.com/couchbaselabs/logg"
)

// the possible contents of a square
type Piece int

const (
	EMPTY      = 0
	RED        = 1
	RED_KING   = 2
	BLACK      = 3
	BLACK_KING = 4
)

type Player int

const (
	RED_PLAYER   = 0
	BLACK_PLAYER = 1
)

type Board [8][8]Piece

func NewBoardFromBoard(otherBoard Board) Board {
	board := Board{}
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			board[row][col] = otherBoard[row][col]
		}
	}
	return board
}

func NewBoard(compactBoard string) Board {

	board := Board{}
	name := "boardlexer"
	_, tokensChannel := lex(name, compactBoard)

	i := 0

	for token := range tokensChannel {

		row := int(i / 8)
		col := int(i % 8)

		switch {
		case token.typ == itemSquareEmpty:
			board[row][col] = EMPTY
		case token.typ == itemSquareRed:
			board[row][col] = RED
		case token.typ == itemSquareRedKing:
			board[row][col] = RED_KING
		case token.typ == itemSquareBlack:
			board[row][col] = BLACK
		case token.typ == itemSquareBlackKing:
			board[row][col] = BLACK_KING
		}

		i += 1

	}

	return board
}

func (board Board) LegalMoves(player Player) []Move {

	moves := []Move{}

	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			location := Location{row: row, col: col}
			movesForLocation := board.legalMovesForLocation(player, location)
			moves = append(moves, movesForLocation...)
		}
	}

	return moves
}

func (board Board) legalMovesForLocation(player Player, loc Location) []Move {
	moves := []Move{}

	jumpMoves := board.singleJumpMovesForLocation(player, loc)
	for _, jumpMove := range jumpMoves {
		if board.canExplodeJumpMove(player, jumpMove) {
			jumpMoveSequences := board.explodeJumpMoveDriver(player, jumpMove)
			for _, moveSequence := range jumpMoveSequences {
				multiJumpMove := NewMove(moveSequence)
				moves = append(moves, multiJumpMove)
			}
		} else {
			moves = append(moves, jumpMove)
		}

	}

	// only check for non-jump moves if we don't have any jump moves
	if len(jumpMoves) == 0 {
		nonJumpMoves := board.nonJumpMovesForLocation(player, loc)
		moves = append(moves, nonJumpMoves...)
	}

	return moves
}

func (board Board) canExplodeJumpMove(player Player, startingMove Move) bool {
	jumpMoveSequences := board.explodeJumpMoveDriver(player, startingMove)
	canExplode := len(jumpMoveSequences[0]) > 0
	logg.Log("canExplode: %v", canExplode)
	return canExplode
}

// Given a starting jump move, return a slice of move slices, where each slice
// of moves represents a particular move sequence of consecutive jumps.
// In each slice of moves, the first move will be startingMove.  In most cases,
// this will return only a single slice of moves, but its possible for jumps to
// "branch", eg, jump to a square where multiple jumps are possible.
// The result will be sorted descending by the longest jump sequence.
func (board Board) explodeJumpMoveDriver(player Player, startingMove Move) [][]Move {
	/*
		// this is the starting board before explode jump move is called
		currentBoardStr = "" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - - - - - - -|" +
			"|- o - o - o - -|" +
			"|X - - - - - - -|" +
			"|- o - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - - - - - -|"

		// down-path - this is a different single jump path,
		// we aren't called with this on this invocation to explodeJumpMove.
		// ignore it, just for context.
		currentBoardStr = "" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - - - - - - -|" +
			"|- o - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - x - - - - -|" +
			"|- - - - - - - -|"

		// ----------- start: time step 0 ---------------

		// up-path - this is the single jump path we are called
		// with on this invocation to explode jump move.
		// note there are several paths
		currentBoardStr = "" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - X - - - - -|" +
			"|- - - o - o - -|" +
			"|- - - - - - - -|" +
			"|- o - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - - - - - -|"

		// ----------- time step 1 --------------

		// up-path-up
		currentBoardStr = "" +
			"|- - - - X - - -|" +
			"|- - - - - o - -|" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - - - - - - -|" +
			"|- o - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - - - - - -|"

		// up-path-down
		currentBoardStr = "" +
			"|- - - - - - - -|" +
			"|- - - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - - - o - -|" +
			"|- - - - X - - -|" +
			"|- o - o - o - -|" +
			"|- - - - - - - -|" +
			"|- - - - - - - -|"

		// ----------- /end time step 1 in parellel universes --------------
	*/

	// so at time step 1, the input was a board state and a jump piece / location,
	// and the output was two different board states (it could have been more),
	// each that has it's own jump piece / location.  in both cases, the path is
	// non terminal.

	boardPostMove := board.applyMove(player, startingMove)

	// moveSeq := NewMoveSeq(boardPostMove, startingMove)

	jumpMoves := boardPostMove.singleJumpMovesForLocation(player, startingMove.to)
	if len(jumpMoves) == 0 {
		moveSequence := []Move{}
		altMoveSequences := [][]Move{moveSequence}
		return altMoveSequences

	} else {

		altMoveSequences := make([][]Move, len(jumpMoves))
		for i, jumpMove := range jumpMoves {
			moveSequence := []Move{startingMove, jumpMove}
			altMoveSequences[i] = moveSequence
		}

		return altMoveSequences

	}

}

/*
	// Given this board state, as well as a location that represents
        // the jumper (row: 2, col: 2 in this case)
	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - o - o - -|" +
		"|- - X - - - - -|" +
		"|- - - o - o - -|" +
		"|- - - - - - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	// return the following BoardMoves

	// up-path-up
	currentBoardStr = "" +
		"|- - - - X - - -|" +
		"|- - - - - o - -|" +
		"|- - - - - - - -|" +
		"|- - - o - o - -|" +
		"|- - - - - - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"

	// up-path-down
	currentBoardStr = "" +
		"|- - - - - - - -|" +
		"|- - - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - o - -|" +
		"|- - - - X - - -|" +
		"|- o - o - o - -|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|"


*/

func (board Board) alternateSingleStepJumpPaths(player Player, loc Location) []BoardMove {

	boardMoves := []BoardMove{}

	jumpMoves := board.singleJumpMovesForLocation(player, loc)
	for _, jumpMove := range jumpMoves {
		boardPostMove := board.applyMove(player, jumpMove)
		boardMove := BoardMove{
			board: boardPostMove,
			move:  jumpMove,
		}
		boardMoves = append(boardMoves, boardMove)
	}
	return boardMoves

}

func (board Board) applyMove(player Player, move Move) Board {

	piece := board.pieceAt(move.from)
	boardPostMove := NewBoardFromBoard(board)

	// delete the piece in the move.from location
	boardPostMove[move.from.row][move.from.col] = EMPTY

	// put the piece in the move.to location
	boardPostMove[move.to.row][move.to.col] = piece

	// delete the piece in the middle location (captured)
	if move.IsJump() {
		jumpedLocation := move.over
		boardPostMove[jumpedLocation.row][jumpedLocation.col] = EMPTY
	}

	return boardPostMove
}

func (board Board) singleJumpMovesForLocation(player Player, loc Location) []Move {

	moves := []Move{}

	playerKingPiece := getPlayerKingPiece(player)
	playerPiece := getPlayerPiece(player)

	piece := board.pieceAt(loc)
	if piece != playerPiece && piece != playerKingPiece {
		return moves
	}

	if board.canJump(player, loc, downLeftOne(loc), downLeftTwo(loc)) {
		move := Move{
			from: loc,
			over: downLeftOne(loc),
			to:   downLeftTwo(loc),
		}
		moves = append(moves, move)
	}
	if board.canJump(player, loc, upRightOne(loc), upRightTwo(loc)) {
		move := Move{
			from: loc,
			over: upRightOne(loc),
			to:   upRightTwo(loc),
		}
		moves = append(moves, move)
	}
	if board.canJump(player, loc, downRightOne(loc), downRightTwo(loc)) {
		move := Move{
			from: loc,
			over: downRightOne(loc),
			to:   downRightTwo(loc),
		}
		moves = append(moves, move)
	}
	if board.canJump(player, loc, upLeftOne(loc), upLeftTwo(loc)) {
		move := Move{
			from: loc,
			over: upLeftOne(loc),
			to:   upLeftTwo(loc),
		}
		moves = append(moves, move)
	}

	return moves

}

func (board Board) nonJumpMovesForLocation(player Player, loc Location) []Move {

	moves := []Move{}

	playerKingPiece := getPlayerKingPiece(player)
	playerPiece := getPlayerPiece(player)

	piece := board.pieceAt(loc)
	if piece != playerPiece && piece != playerKingPiece {
		return moves
	}

	if board.canMove(player, loc, downLeftOne(loc)) {
		moves = append(moves, Move{from: loc, to: downLeftOne(loc)})
	}
	if board.canMove(player, loc, upRightOne(loc)) {
		moves = append(moves, Move{from: loc, to: upRightOne(loc)})
	}
	if board.canMove(player, loc, downRightOne(loc)) {
		moves = append(moves, Move{from: loc, to: downRightOne(loc)})
	}
	if board.canMove(player, loc, upLeftOne(loc)) {
		moves = append(moves, Move{from: loc, to: upLeftOne(loc)})
	}

	return moves

}

func (board Board) canMove(player Player, start, dest Location) bool {

	if dest.isOffBoard() {
		return false
	}

	if board.pieceAt(dest) != EMPTY {
		return false // already contains a piece
	}

	switch {
	case player == RED_PLAYER:
		if board.pieceAt(start) == RED && isMovingDown(start, dest) {
			return false // Regular red piece can only move up
		}
		return true // move is legal
	default: // BLACK_PLAYER
		if board.pieceAt(start) == BLACK && isMovingUp(start, dest) {
			return false // Regular black piece can only move down
		}
		return true // move is legal
	}

}

func (board Board) canJump(player Player, start, intermediate, dest Location) bool {

	if dest.isOffBoard() {
		return false
	}

	if board.pieceAt(dest) != EMPTY {
		return false // already contains a piece
	}

	switch {
	case player == RED_PLAYER:
		if board.pieceAt(start) == RED && isMovingDown(start, dest) {
			return false // Regular red piece can only move up
		}
		intermediate := board.pieceAt(intermediate)
		if intermediate != BLACK && intermediate != BLACK_KING {
			return false // there is no black piece to jump
		}
		return true // jump is legal
	default: // BLACK_PLAYER
		if board.pieceAt(start) == BLACK && isMovingUp(start, dest) {
			return false // Regular black piece can only move down
		}
		intermediate := board.pieceAt(intermediate)
		if intermediate != RED && intermediate != RED_KING {
			return false // there is no red piece to jump
		}
		return true // jump is legal
	}

}

func (board Board) pieceAt(loc Location) Piece {
	return board[loc.row][loc.col]
}

// one square "down" (row increasing) and to the "left" from the perspective
// of the piece moving downwards (col increasing)
func downLeftOne(loc Location) Location {
	return Location{
		row: loc.row + 1,
		col: loc.col + 1,
	}
}

func downLeftTwo(loc Location) Location {
	return downLeftOne(downLeftOne(loc))
}

func downRightOne(loc Location) Location {
	return Location{
		row: loc.row + 1,
		col: loc.col - 1,
	}
}

func downRightTwo(loc Location) Location {
	return downRightOne(downRightOne(loc))
}

func upLeftOne(loc Location) Location {
	return Location{
		row: loc.row - 1,
		col: loc.col - 1,
	}
}

func upLeftTwo(loc Location) Location {
	return upLeftOne(upLeftOne(loc))
}

func upRightOne(loc Location) Location {
	return Location{
		row: loc.row - 1,
		col: loc.col + 1,
	}
}

func upRightTwo(loc Location) Location {
	return upRightOne(upRightOne(loc))
}

func isMovingDown(start, dest Location) bool {
	return dest.row > start.row
}

func isMovingUp(start, dest Location) bool {
	return dest.row < start.row
}

func getPlayerPiece(player Player) Piece {
	switch {
	case player == RED_PLAYER:
		return RED
	default: // BLACK_PLAYER
		return BLACK
	}
}

func getPlayerKingPiece(player Player) Piece {
	switch {
	case player == RED_PLAYER:
		return RED_KING
	default: // BLACK_PLAYER
		return BLACK_KING
	}
}
