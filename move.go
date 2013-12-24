package checkerscore

type Move struct {
	from Location
	to   Location

	// if a move contains submoves, this means it was a double/triple/etc
	// jump.  the moves (jumps) that compose this double jump will be in order
	// and the submoves completely describe the sequence of move and the "outer move"
	// does not provide any additional information.
	submoves []Move
}

func (m *Move) AdjustForSubmoves() {
	// TODO: update from/to based on submoves
}
