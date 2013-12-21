package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func DisTestNewBoard(t *testing.T) {

	boardStr := "" +
		"|- x - x - x - x|" +
		"|x - x - x - x -|" +
		"|- x - x - x - x|" +
		"|- - - - - - - -|" +
		"|- - - - - - - -|" +
		"|o - o - o - o -|" +
		"|- o - o - o - o|" +
		"|o - o - o - o -|"
	board := NewBoard(boardStr)
	assert.Equals(t, board[0][0], EMPTY)
	assert.Equals(t, board[0][1], RED)
	// etc...

}
