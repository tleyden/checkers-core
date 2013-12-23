package checkerscore

import (
	"github.com/couchbaselabs/logg"
	_ "github.com/couchbaselabs/logg"
)

// the possible contents of a square
type Square int

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

type Board [8][8]Square

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
			square := board[row][col]
			movesForSquare := board.legalMovesForSquare(player, square)
			moves = append(moves, movesForSquare...)
		}
	}

	return moves
}

func (board Board) legalMovesForSquare(player Player, square Square) []Move {

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

	playerKingSquare := getPlayerKing(player)
	playerSquare := getPlayerSquare(player)

	logg.Log("%v", playerKingSquare)

	if square == playerSquare || square == playerKingSquare {

		// ...
	}

	return moves
}

func getPlayerSquare(player Player) Square {
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

func getPlayerKing(player Player) Square {
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
