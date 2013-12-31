package checkerscore

type Location struct {
	row int
	col int
}

func NewLocation(row, col int) Location {
	return Location{row: row, col: col}
}

func (loc Location) isOffBoard() bool {

	if loc.row < 0 || loc.row >= 8 || loc.col < 0 || loc.col >= 8 {
		return true
	}
	return false

}

func (loc Location) Row() int {
	return loc.row
}

func (loc Location) Col() int {
	return loc.col
}

func (loc Location) Equals(otherLoc Location) bool {
	return loc.row == otherLoc.row && loc.col == otherLoc.col
}
