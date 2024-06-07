package preroll

import (
	// "fmt"
	"strings"
	"sync"
	"unicode/utf8"
)

type Lexeme struct {
	Type LexemeType
	Val  string
}

type LexemeType int

const (
	AdvantageToken LexemeType = iota
	DisadvantageToken
	DieToken
	FacesToken
	DieQuantityToken
	AdditionToken
	SubtractionToken
)

const (
	EOF rune = 0
)

func Lex(name, input string) *Lexer {

	l := &Lexer{
		name:  name,
		input: input,
		state: lexDiceQuantity,
		Items: make([]Lexeme, 0),
	}
	// go l.run()
	l.run()

	return l
}

type Lexer struct {
	// logger
	name, input                      string
	start, pos, width                int
	state                            stateFn
	Items                            []Lexeme
	widthMutex, startMutex, posMutex sync.Mutex
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()

	return r
}

func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	return false
}

func (l *Lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

func (l *Lexer) emit(t Lexeme) {
	l.SetStart(l.GetPos())
	l.Items = append(l.Items, t)
}

func (l *Lexer) next() rune {
	var r rune
	if l.GetPos() >= len(l.input) {
		l.SetWidth(0)
		return EOF
	}
	r, width := utf8.DecodeRuneInString(l.input[l.GetPos():])
	l.SetWidth(width)
	l.SetPos(l.GetPos() + l.GetWidth())
	return r
}

func (l *Lexer) run() {
	for state := lexDiceQuantity; state != nil; {
		state = state(l)
	}
	/* original concurrent implementation
	defer l.chanMutex.Unlock()
		for state := lexText; state != nil; {
					state = state(l)
						}
							l.chanMutex.Lock()
								close(l.items)
	*/
}

func (l *Lexer) GetPos() int {
	defer l.posMutex.Unlock()
	l.posMutex.Lock()
	return l.pos
}

func (l *Lexer) SetPos(pos int) {
	defer l.posMutex.Unlock()
	l.posMutex.Lock()
	l.pos = pos
}

func (l *Lexer) GetWidth() int {
	defer l.widthMutex.Unlock()
	l.widthMutex.Lock()
	return l.width
}

func (l *Lexer) SetWidth(width int) {
	defer l.widthMutex.Unlock()
	l.widthMutex.Lock()
	l.width = width
}

func (l *Lexer) GetStart() int {
	defer l.startMutex.Unlock()
	l.startMutex.Lock()
	return l.start
}

func (l *Lexer) SetStart(start int) {
	defer l.startMutex.Unlock()
	l.startMutex.Lock()
	l.start = start
}

func (l *Lexer) ignore() {
	l.SetStart(l.pos)
}

func (l *Lexer) backup() {
	l.SetPos(l.GetPos() - l.GetWidth())
}
