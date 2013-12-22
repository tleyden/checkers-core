package checkerscore

import (
	"github.com/couchbaselabs/go.assert"
	"github.com/couchbaselabs/logg"
	"testing"
)

func TestLexer(t *testing.T) {

	name := "testlexer"
	inputString := "| - x X - o O || x |"

	_, tokensChannel := lex(name, inputString)

	item := <-tokensChannel
	assert.Equals(t, item.typ, itemSquareEmpty)
	assert.Equals(t, item.val, "-")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareRed)
	assert.Equals(t, item.val, "x")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareRedKing)
	assert.Equals(t, item.val, "X")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareEmpty)
	assert.Equals(t, item.val, "-")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareBlack)
	assert.Equals(t, item.val, "o")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareBlackKing)
	assert.Equals(t, item.val, "O")

	item = <-tokensChannel
	assert.Equals(t, item.typ, itemSquareRed)
	assert.Equals(t, item.val, "x")

	item = <-tokensChannel
	logg.Log("item type: %v item value: %v", item.typ, item.val)
	assert.Equals(t, item.typ, itemEOF)

}
