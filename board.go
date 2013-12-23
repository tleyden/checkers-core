package checkerscore

import (
	"github.com/couchbaselabs/logg"
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

	/*

	   if (board[row][col] == player || board[row][col] == playerKing) {
	       if (canJump(player, row, col, row+1, col+1, row+2, col+2))
	           moves.add(new CheckersMove(row, col, row+2, col+2));
	       if (canJump(player, row, col, row-1, col+1, row-2, col+2))
	           moves.add(new CheckersMove(row, col, row-2, col+2));
	       if (canJump(player, row, col, row+1, col-1, row+2, col-2))
	           moves.add(new CheckersMove(row, col, row+2, col-2));
	       if (canJump(player, row, col, row-1, col-1, row-2, col-2))
	           moves.add(new CheckersMove(row, col, row-2, col-2));
	   }

	*/

	moves := []Move{}

	playerKingPiece := getPlayerKingPiece(player)
	playerPiece := getPlayerPiece(player)

	logg.Log("%v", playerKingPiece)

	piece := board.pieceAt(loc)
	if piece == playerPiece || piece == playerKingPiece {
		if board.canJump(player, loc, downLeftOne(loc), downLeftTwo(loc)) {
			moves = append(moves, Move{from: loc, to: downLeftTwo(loc)})
		}

	}

	return moves
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
	case player == BLACK_PLAYER:
		if board.pieceAt(start) == BLACK && isMovingUp(start, dest) {
			return false // Regular black piece can only move down
		}
		intermediate := board.pieceAt(intermediate)
		if intermediate != RED && intermediate != RED_KING {
			return false // there is no red piece to jump
		}
		return true // jump is legal
	default:
		panic("Invalid player")
		return false
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
	case player == BLACK_PLAYER:
		return BLACK
	default:
		panic("Invalid player")
		return BLACK
	}
}

func getPlayerKingPiece(player Player) Piece {
	switch {
	case player == RED_PLAYER:
		return RED_KING
	case player == BLACK_PLAYER:
		return BLACK_KING
	default:
		panic("Invalid player")
		return BLACK_KING
	}
}
