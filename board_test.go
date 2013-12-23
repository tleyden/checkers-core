package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestNewBoard(t *testing.T) {

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
	assert.Equals(t, int(board[0][0]), int(EMPTY)) // TODO: why is cast to int() needed?
	assert.Equals(t, int(board[0][1]), int(RED))
	assert.Equals(t, int(board[7][0]), int(BLACK))
	assert.Equals(t, int(board[7][7]), int(EMPTY))

	// etc...

}
