package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"testing"
)

func TestLocationEquals(t *testing.T) {
	loc1 := Location{row: 4, col: 0}
	loc2 := Location{row: 4, col: 0}
	assert.True(t, loc1 == loc2)
}
