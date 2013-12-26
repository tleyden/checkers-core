package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestIsJumpNegCase(t *testing.T) {
	move := Move{
		from: Location{row: 4, col: 0},
		to:   Location{row: 3, col: 1},
	}
	assert.False(t, move.IsJump())

	move = Move{
		from: Location{row: 4, col: 0},
		to:   Location{row: 4, col: 0},
	}
	assert.False(t, move.IsJump())

}

func TestIsJumpEasy(t *testing.T) {
	move := Move{
		from: Location{row: 4, col: 0},
		to:   Location{row: 2, col: 2},
	}
	assert.True(t, move.IsJump())

}

func TestIsJumpHard(t *testing.T) {

	submove1 := Move{
		from: Location{row: 4, col: 0},
		to:   Location{row: 2, col: 2},
	}
	submove2 := Move{
		from: Location{row: 2, col: 2},
		to:   Location{row: 0, col: 0},
	}
	submoves := []Move{submove1, submove2}

	move := NewMove(submoves)

	assert.True(t, move.IsJump())

}

func TestIsJumpHard2(t *testing.T) {

	submove1 := Move{
		from: Location{row: 4, col: 0},
		to:   Location{row: 2, col: 2},
	}
	submove2 := Move{
		from: Location{row: 2, col: 2},
		to:   Location{row: 4, col: 4},
	}
	submoves := []Move{submove1, submove2}

	move := NewMove(submoves)

	assert.True(t, move.IsJump())

}
