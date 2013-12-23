package checkerscore

type Location struct {
	row int
	col int
}

func (loc Location) isOffBoard() bool {

	if loc.row < 0 || loc.row >= 8 || loc.col < 0 || loc.col >= 8 {
		return true
	}
	return false

}
