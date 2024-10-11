package lexer

import (
	"errors"
	"log/slog"
)

var ErrNoMatches = errors.New("lexer did not match any tokens")

type pattern struct {
	pattern matcher
	handler tokenHandler
}

type matcher func(string) bool

type tokenHandler func(lex *lexer, pattern pattern)

type lexer struct {
	patterns []pattern
	tokens   []Token
	source   string
	pos      int
	logger   *slog.Logger
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

func (lex *lexer) advance() {
	lex.pos += 1
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) push(token Token) {
	lex.tokens = append(lex.tokens, token)
}

func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

func createLexer(source string, logger *slog.Logger) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		tokens: make([]Token, 0),
		logger: logger,
	}
}

func Lex(source string, logger *slog.Logger) ([]Token, error) {
	lex := createLexer(source, logger)
	for !lex.at_eof() {
		matched := false
		for _, pattern := range lex.patterns {
			logger.With("pattern", pattern).Info("matching pattern")
			// find the index of the pattern in lex.remainder()
		}

		if !matched {
			return lex.tokens, ErrNoMatches
		}
	}

	return lex.tokens, nil
}
