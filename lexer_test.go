package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
	"testing"
)

func TestLexer(t *testing.T) {

	name := "testlexer"
	inputString := "| - |"
	_, tokensChannel := lex(name, inputString)
	item := <-tokensChannel
	logg.Log("item type: %v item value: %v", item.typ, item.val)
	assert.Equals(t, item.typ, itemSquareEmpty)
	assert.Equals(t, item.val, "-")

}
