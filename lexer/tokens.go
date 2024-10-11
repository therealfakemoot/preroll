package lexer

type TokenKind int

//go:generate stringer -type=Token
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
