package checkerscore

// the possible contents of a square
type Piece int

const (
	EMPTY = Piece(iota)
	RED
	RED_KING
	BLACK
	BLACK_KING
)

func (piece Piece) String() string {

	switch piece {
	case EMPTY:
		return "-"
	case BLACK:
		return "o" // TODO: use unicode ●
	case BLACK_KING:
		return "O" // TODO: use unicode ♚
	case RED:
		return "x" // TODO: use unicode ○
	case RED_KING:
		return "X" // TODO: use unicode ♔
	}
	panic("Unknown piece")

}

func (piece Piece) WeightedValue() float64 {
	switch piece {
	case BLACK:
		return 1.0
	case BLACK_KING:
		return 1.3
	case RED:
		return 1.0
	case RED_KING:
		return 1.3
	default:
		return 0.0 // EMPTY
	}

}

func (piece Piece) OwnedBy(player Player) bool {
	switch player {
	case BLACK_PLAYER:
		switch piece {
		case BLACK:
			return true
		case BLACK_KING:
			return true
		case RED:
			return false
		case RED_KING:
			return false
		default:
			return false
		}
	default:
		switch piece {
		case BLACK:
			return false
		case BLACK_KING:
			return false
		case RED:
			return true
		case RED_KING:
			return true
		default:
			return false
		}

	}
}

func (piece Piece) IsKing() bool {
	switch piece {
	case BLACK:
		return false
	case BLACK_KING:
		return true
	case RED:
		return false
	case RED_KING:
		return true
	default:
		return false
	}

}

func (piece Piece) King() Piece {
	switch piece {
	case BLACK:
		return BLACK_KING
	case BLACK_KING:
		return piece
	case RED:
		return RED_KING
	case RED_KING:
		return piece
	default:
		return piece
	}

}
