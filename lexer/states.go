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
		return lexDieCount
	default:
		logger.Debug("no modifier found")
		return lexDieCount
	}
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
	return nil
}

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
			return lexDieXXX
		}

		if l.next() == EOF {
			break
		}
	}
	return nil
}

func lexDieXXX(l *lexer) stateFunc {
	if strings.HasPrefix(l.input[l.pos:], facesOpen) {
		l.pos += len(facesOpen)
		l.emit(facesOpenToken)
		return lexFacesXXX
	}
	return lexNumber
}

func lexFacesXXX(l *lexer) stateFunc {
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
	l.acceptRun(DIGITS)
	l.emit(numberToken)

	return startState
}
