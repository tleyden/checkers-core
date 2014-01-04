package checkerscore

type Player int

const (
	RED_PLAYER = Player(iota)
	BLACK_PLAYER
)

func (player Player) Opponent() Player {
	switch player {
	case RED_PLAYER:
		return BLACK_PLAYER
	default:
		return RED_PLAYER
	}
}
