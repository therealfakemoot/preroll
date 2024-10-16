package lexer

import (
	"strings"
)

func isDecimalDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isASCIIAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
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
			l.emit(dropLowestToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], explode):
			l.pos += len(explode)
			l.emit(explodeToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], die):
			l.pos += len(die)
			l.emit(dieToken)
			return startState
		case strings.HasPrefix(l.input[l.pos:], addition):
			l.pos += len(addition)
			l.emit(additionToken)
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], subtraction):
			l.pos += len(subtraction)
			l.emit(subtractionToken)
			return lexNumber
		case isDecimalDigit(l.peek()):
			return lexNumber
		case strings.HasPrefix(l.input[l.pos:], facesOpen):
			l.pos += len(facesOpen)
			l.emit(facesOpenToken)
			// return lexFaces
			return lexDie
		}

		if l.next() == EOF {
			break
		}
	}
	return nil
}

func lexDie(l *lexer) stateFunc {
	logger := l.logger.WithGroup("lexDie")
	logger.With("input", l.input[l.pos:]).Info("checking for faces")
	if strings.HasPrefix(l.input[l.pos:], facesOpen) {
		l.pos += len(facesOpen)
		l.emit(facesOpenToken)
		return lexFaces
	}
	return lexNumber
}

func lexFaces(l *lexer) stateFunc {
	// TODO: figure out how to lex comma separated strings
	for {
		switch l.next() {
		case ',':
			l.emit(faceToken)
		case '}':
			l.emit(facesCloseToken)
			return startState
		default:
			l.errorf("incorrect faces syntax")
			return nil
		}
	}
}

func lexNumber(l *lexer) stateFunc {
	digits := "0123456789"
	l.acceptRun(digits)
	l.emit(numberToken)

	return startState
}
