package lexer

import (
	"testing"
)

func Test_Basic(t *testing.T) {
	l := Lex("1d20", nil)
	for _, token := range l.Items() {
		t.Logf("{Token: %s, Raw:%q}\n", token.Type, token.Raw)
	}
}
