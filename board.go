package checkerscore

import (
	_ "github.com/couchbaselabs/logg"
)

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

	var board Board

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
