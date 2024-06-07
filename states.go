package preroll

import (
// "unicode"
)

type stateFn func(*Lexer) stateFn

func lexDiceQuantity(l *Lexer) stateFn {
	l.acceptRun("0123456789")
	l.emit(Lexeme{
		Type: DieQuantityToken,
		Val:  l.input[l.GetStart() : l.GetPos()+1],
	})
	return lexDie
}

func lexDie(l *Lexer) stateFn {
	r := l.peek()
	if r == 'd' {
		l.emit(Lexeme{
			Type: DieToken,
			Val:  l.input[l.GetStart():l.GetPos()],
		})
	}

	return nil
}
