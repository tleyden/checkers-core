package checkerscore

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*

Lexer used for parsing compact board representations, eg converting:

		"|- x - x - x - x|"
		"|x - x - x - x -|"
		"|- x - x - x - x|"
		"|- - - - - - - -|"
		"|- - - - - - - -|"
		"|o - o - o - o -|"
		"|- o - o - o - o|"
		"|o - o - o - o -|"

into a Board struct.  Here are is what the lexer does:

* `\` - pipe delimiter ignored
* ` ` - whitespace ignored
* `x` - emit item w/ item.typ: itemSquareRed
* `X` - emit item w/ item.typ: itemSquareRedKing
* `o` - emit item w/ item.typ: itemSquareBlack
* `O` - emit item w/ item.typ: itemSquareBlacKing
* `-` - emit item w/ item.typ: itemSquareEmpty

*/

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	items chan item // channel of scanned items.
}

type item struct {
	typ itemType // Type, such as itemNumber.
	val string   // Value, such as "23.2".
}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFn func(*lexer) stateFn

type itemType int

const (
	itemError itemType = iota
	itemSquareEmpty
	itemSquareRed
	itemSquareRedKing
	itemSquareBlack
	itemSquareBlackKing
	itemEOF
)

const (
	pipe            = "|"
	squareEmpty     = '-'
	squareRed       = 'x'
	squareRedKing   = 'X'
	squareBlack     = 'o'
	squareBlackKing = 'O'
)

const eof = -1

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

func lex(name, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run() // Concurrently run state machine.
	return l, l.items
}

// run lexes the input by executing state functions until
// the state is nil.
func (l *lexer) run() {
	for state := lexOutsideRow; state != nil; {
		state = state(l)
	}
	close(l.items) // No more tokens will be delivered.
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// next returns the next rune in the input.
func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func lexOutsideRow(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], pipe) {
			l.next()
			return lexInsideRow // Next state.
		}
		if l.next() == eof {
			break
		}
	}
	// Correctly reached EOF.
	l.emit(itemEOF) // Useful to make EOF a token.
	return nil      // Stop the run loop.
}

func lexInsideRow(l *lexer) stateFn {

	for {
		if strings.HasPrefix(l.input[l.pos:], pipe) {
			l.next()
			return lexOutsideRow // Next state.
		}
		switch r := l.next(); {
		case r == eof || r == '\n':
			return l.errorf("unclosed action")
		case isSpace(r):
			l.ignore()
		case r == squareEmpty:
			l.emit(itemSquareEmpty)
			return lexInsideRow
		case r == squareRed:
			l.emit(itemSquareRed)
			return lexInsideRow
		case r == squareRedKing:
			l.emit(itemSquareRedKing)
			return lexInsideRow
		case r == squareBlack:
			l.emit(itemSquareBlack)
			return lexInsideRow
		case r == squareBlackKing:
			l.emit(itemSquareBlackKing)
			return lexInsideRow
		}
	}

	// Correctly reached EOF.
	l.emit(itemEOF) // Useful to make EOF a token.
	return nil      // Stop the run loop.
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...)}
	return nil
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}
