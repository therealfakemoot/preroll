package lexer

import (
	"log/slog"
	"strings"
	"unicode/utf8"
)

type lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens chan Token
	states []stateFunc
	logger *slog.Logger
}

func Lex(input string, logger *slog.Logger) (*lexer, chan Token) {
	l := &lexer{
		input:  input,
		pos:    0,
		tokens: make(chan Token),
		states: make([]stateFunc, 0),
		logger: logger,
	}

	go l.run()
	return l, l.tokens
}

// next returns the next rune in the input.
func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) emit(t tokenType) {
	logger := l.logger.WithGroup("emit")
	token := Token{t, l.input[l.start:l.pos]}
	logger.With("token", token).Info("found token")
	l.tokens <- token
	l.start = l.pos
}

func (l *lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
