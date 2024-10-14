package lexer

import (
	"strings"
)

func isDecimalDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

type stateFunc func(*lexer) stateFunc

func startState(l *lexer) stateFunc {
	for {
		switch {
		case strings.HasPrefix(l.input[l.pos:], keepHighest):
			l.pos += len(keepHighest)
			l.emit(keepHighestToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], keepLowest):
			l.pos += len(keepLowest)
			l.emit(keepLowestToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], dropHighest):
			l.pos += len(dropHighest)
			l.emit(dropHighestToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], dropLowest):
			l.pos += len(dropLowest)
			l.emit(keepLowestToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], explode):
			l.pos += len(explode)
			l.emit(explodeToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], die):
			l.pos += len(die)
			l.emit(dieToken)
			// TODO: create a stateFunc for lexing die faces instead of jumping straight into numbers
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], die):
			l.pos += len(die)
			l.emit(dieToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], addition):
			l.pos += len(die)
			l.emit(dieToken)
			return lexNumber
		case isDecimalDigit(l.peek()):
			return lexNumber
		default:
			return nil
		}

		if l.next() == EOF {
			break
		}
	}
	return nil
}

func lexNumber(l *lexer) stateFunc {
	digits := "0123456789"
	l.acceptRun(digits)
	l.emit(numberToken)

	return startState
}
