package checkerscore

import (
	_ "github.com/couchbaselabs/logg"
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
			jumpMoveSequences := board.explodeJumpMove(player, jumpMove)
			for _, moveSequence := range jumpMoveSequences {
				nJumpMove := NewMove(moveSequence)
				moves = append(moves, nJumpMove)
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
	// TODO
	return false
}

// Given a starting jump move, return a slice of move slices, where each slice
// of moves represents a particular move sequence of consecutive jumps.
// In each slice of moves, the first move will be startingMove.  In most cases,
// this will return only a single slice of moves, but its possible for jumps to
// "branch", eg, jump to a square where multiple jumps are possible.
// The result will be sorted descending by the longest jump sequence.
func (board Board) explodeJumpMove(player Player, startingMove Move) [][]Move {

	moveSequences := []Move{}
	altMoveSequences := [][]Move{moveSequences}
	return altMoveSequences
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
		moves = append(moves, Move{from: loc, to: downLeftTwo(loc)})
	}
	if board.canJump(player, loc, upRightOne(loc), upRightTwo(loc)) {
		moves = append(moves, Move{from: loc, to: upRightTwo(loc)})
	}
	if board.canJump(player, loc, downRightOne(loc), downRightTwo(loc)) {
		moves = append(moves, Move{from: loc, to: downRightTwo(loc)})
	}
	if board.canJump(player, loc, upLeftOne(loc), upLeftTwo(loc)) {
		moves = append(moves, Move{from: loc, to: upLeftTwo(loc)})
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
