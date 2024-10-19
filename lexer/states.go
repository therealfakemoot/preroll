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

func lexModifier(l *lexer) stateFunc {
	logger := l.logger.WithGroup("lexModifier")
	logger.Debug("entering lexModifier")
	switch {
	case strings.HasPrefix(l.input[l.pos:], keepHighest):
		l.pos += len(keepHighest)
		l.emit(keepHighestToken)
		return lexDieCount
	case strings.HasPrefix(l.input[l.pos:], keepLowest):
		l.pos += len(keepLowest)
		l.emit(keepLowestToken)
		return lexDieCount
	case strings.HasPrefix(l.input[l.pos:], dropHighest):
		l.pos += len(dropHighest)
		l.emit(dropHighestToken)
		return lexDieCount
	case strings.HasPrefix(l.input[l.pos:], dropLowest):
		l.pos += len(dropLowest)
		l.emit(dropLowestToken)
		return lexDieCount
	case strings.HasPrefix(l.input[l.pos:], explode):
		l.pos += len(explode)
		l.emit(explodeToken)
		return lexExplode
	default:
		logger.Debug("no modifier found")
		return lexDieCount
	}
}

func lexExplode(l *lexer) stateFunc {
	if strings.HasPrefix(l.input[l.pos:], explodeOpen) {
		l.pos += len(explodeOpen)
		l.emit(explodeOpenToken)
		if l.accept(DIGITS) {
			l.acceptRun(DIGITS)
			l.emit(explodeCountToken)
		}
		if l.peek() == '}' {
			l.next()
			l.emit(explodeCloseToken)
			return lexDieCount
		}
	}
	return lexDieCount
}

func lexDieCount(l *lexer) stateFunc {
	logger := l.logger.WithGroup("lexDieCount")
	logger.With("lexer", l).Debug("entering lexDieCount")
	if l.accept(NON_ZERO_DIGITS) {
		l.acceptRun(DIGITS)
		l.emit(numberToken)
	}
	return lexDie
}

func lexDie(l *lexer) stateFunc {
	if strings.HasPrefix(l.input[l.pos:], die) {
		l.pos += len(die)
		l.emit(dieToken)
		return lexFaces
	}
	return nil
}

func lexFaces(l *lexer) stateFunc {
	switch {
	case strings.HasPrefix(l.input[l.pos:], facesOpen):
		l.pos += len(facesOpen)
		l.emit(facesOpenToken)
		for {
			l.acceptRun(DIGITS + ASCII_ALPHA)
			l.emit(faceToken)
			switch {
			case l.peek() == ',':
				l.next()
				l.emit(facesSeparatorToken)
			case l.peek() == '}':
				l.next()
				l.emit(facesCloseToken)
				return lexAddSubtract
			}
		}
	case l.accept(NON_ZERO_DIGITS):
		l.acceptRun(DIGITS)
		l.emit(numberToken)
		return lexAddSubtract
	}
	return nil
}

func lexAddSubtract(l *lexer) stateFunc {
	switch {
	case strings.HasPrefix(l.input[l.pos:], addition):
		l.pos += len(addition)
		l.emit(additionToken)
	case strings.HasPrefix(l.input[l.pos:], subtraction):
		l.pos += len(subtraction)
		l.emit(subtractionToken)
	}
	switch {
	case l.accept(NON_ZERO_DIGITS):
		l.acceptRun(DIGITS)
		l.emit(numberToken)
	}
	if l.peek() == EOF {
		return nil
	}
	return lexModifier
}
