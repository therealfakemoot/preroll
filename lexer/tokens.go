package lexer

//go:generate stringer -type=TokenKind
type TokenKind int

type Token struct {
	Kind  TokenKind
	Value string
}

func (t Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == t.Kind {
			return true
		}
	}

	return false
}

const (
	EOF TokenKind = iota
	ADVANTAGE
	DISADVANTAGE
	DIE
	FACES
	QUANTITY
	ADDITION
	SUBTRACTION
	EXPLODE
)

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}
