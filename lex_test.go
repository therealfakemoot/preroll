package preroll_test

import (
	"testing"

	"github.com/therealfakemoot/preroll"
)

func Test_BasicDice(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []preroll.Lexeme
	}{
		{
			name:  "one d10",
			input: "d10",
			expected: []preroll.Lexeme{
				preroll.Lexeme{
					Type: preroll.DieToken,
					Val:  "d",
				},
				preroll.Lexeme{
					Type: preroll.FacesToken,
					Val:  "10",
				},
			},
		},
		{
			name:  "two d10",
			input: "2d10",
			expected: []preroll.Lexeme{
				preroll.Lexeme{
					Type: preroll.DieQuantityToken,
					Val:  "2",
				},
				preroll.Lexeme{
					Type: preroll.DieToken,
					Val:  "d",
				},
				preroll.Lexeme{
					Type: preroll.FacesToken,
					Val:  "10",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Fail()
		})
	}

}
