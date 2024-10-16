package lexer

//go:generate stringer -type=tokenType
type tokenType int

const (
	EOF = '\u0000'
)

const (
	keepHighest = "kh"
	keepLowest  = "kl"
	dropHighest = "dh"
	dropLowest  = "dl"
	explode     = "!"
	die         = "d"
	facesOpen   = "{"
	facesClose  = "}"
	addition    = "+"
	subtraction = "-"
)

const (
	errorToken tokenType = iota
	eofToken
	keepHighestToken
	keepLowestToken
	dropHighestToken
	dropLowestToken
	explodeToken
	dieToken
	numberToken
	additionToken
	subtractionToken
	facesOpenToken
	facesCloseToken
	faceToken
)

type Token struct {
	Type tokenType
	Raw  string
}
