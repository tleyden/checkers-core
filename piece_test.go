package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestIsKing(t *testing.T) {
	assert.True(t, BLACK_KING.IsKing())
	assert.False(t, BLACK.IsKing())

}

func TestKing(t *testing.T) {
	assert.Equals(t, BLACK_KING.King(), BLACK_KING)
	assert.Equals(t, BLACK.King(), BLACK_KING)
	assert.Equals(t, RED_KING.King(), RED_KING)
	assert.Equals(t, RED.King(), RED_KING)

}
