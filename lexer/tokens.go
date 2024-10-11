package lexer

//go:generate stringer -type=TokenKind
type TokenKind int

type Token struct {
	Kind  TokenKind
	Value string
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
