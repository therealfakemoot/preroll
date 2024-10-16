package lexer

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func Test_Simple(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1d20", expected: []Token{
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
			},
		},
		{
			input: "2d4", expected: []Token{
				{numberToken, "2"},
				{dieToken, "d"},
				{numberToken, "4"},
			},
		},
		{
			input: "0d0", expected: []Token{
				{numberToken, "0"},
				{dieToken, "d"},
				{numberToken, "0"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Dropping(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "dh1d20", expected: []Token{
				{dropHighestToken, "dh"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
			},
		},
		{
			input: "dl2d4", expected: []Token{
				{dropLowestToken, "dl"},
				{numberToken, "2"},
				{dieToken, "d"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Keeping(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "kh1d20", expected: []Token{
				{keepHighestToken, "kh"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
			},
		},
		{
			input: "kl2d4", expected: []Token{
				{keepLowestToken, "kl"},
				{numberToken, "2"},
				{dieToken, "d"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}
