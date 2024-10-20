package lexer

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

var logger = slog.New(
	slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
	),
)

func Test_Simple(t *testing.T) {
	logger := logger.With("test", "Test_Simple")
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
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Dropping(t *testing.T) {
	logger := logger.With("test", "Test_Dropping")
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
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Keeping(t *testing.T) {
	logger := logger.With("test", "Test_Keeping")
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
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_NumberSubtraction(t *testing.T) {
	logger := logger.With("test", "Test_NumberSubtraction")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1d20-4", expected: []Token{
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
				{subtractionToken, "-"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_NumberAddition(t *testing.T) {
	logger := logger.With("test", "Test_NumberAddition")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1d20+4", expected: []Token{
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
				{additionToken, "+"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Faces(t *testing.T) {
	logger := logger.With("test", "Test_Faces")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "1d{red,blue,green}", expected: []Token{
				{numberToken, "1"},
				{dieToken, "d"},
				{facesOpenToken, "{"},
				{faceToken, "red"},
				{facesSeparatorToken, ","},
				{faceToken, "blue"},
				{facesSeparatorToken, ","},
				{faceToken, "green"},
				{facesCloseToken, "}"},
			},
		},
		{
			input: "1d{2,7,194}", expected: []Token{
				{numberToken, "1"},
				{dieToken, "d"},
				{facesOpenToken, "{"},
				{faceToken, "2"},
				{facesSeparatorToken, ","},
				{faceToken, "7"},
				{facesSeparatorToken, ","},
				{faceToken, "194"},
				{facesCloseToken, "}"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Complex(t *testing.T) {
	logger := logger.With("test", "Test_Complex")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "dh3d20-4", expected: []Token{
				{dropHighestToken, "dh"},
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{subtractionToken, "-"},
				{numberToken, "4"},
			},
		},
		{
			input: "kl2d69+4", expected: []Token{
				{keepLowestToken, "kl"},
				{numberToken, "2"},
				{dieToken, "d"},
				{numberToken, "69"},
				{additionToken, "+"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_RollAdditionSimple(t *testing.T) {
	logger := logger.With("test", "Test_RollAdditionSimple")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3d20+1", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{additionToken, "+"},
				{numberToken, "1"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_RollSubtractionSimple(t *testing.T) {
	logger := logger.With("test", "Test_RollAdditionSimple")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3d20-1", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{subtractionToken, "-"},
				{numberToken, "1"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_RollAdditionComplex(t *testing.T) {
	logger := logger.With("test", "Test_RollAdditionComplex")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3d20+1d4", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{additionToken, "+"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_RollSubtractionComplex(t *testing.T) {
	logger := logger.With("test", "Test_RollSubtractionComplex")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3d20-1d4", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{subtractionToken, "-"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "4"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_RollComplexRolls(t *testing.T) {
	logger := logger.With("test", "Test_RollSubtractionComplex")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3d20-1d4+1d5", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{subtractionToken, "-"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "4"},
				{additionToken, "+"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "5"},
			},
		},
		{
			input: "3d20+1d4-1d5", expected: []Token{
				{numberToken, "3"},
				{dieToken, "d"},
				{numberToken, "20"},
				{additionToken, "+"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "4"},
				{subtractionToken, "-"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "5"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fail()
			}
		})
	}
}

func Test_Explode(t *testing.T) {
	logger := logger.With("test", "Test_Simple")
	cases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "!1d20", expected: []Token{
				{explodeToken, "!"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
			},
		},
		{
			input: "!{3}1d20", expected: []Token{
				{explodeToken, "!"},
				{explodeOpenToken, "{"},
				{explodeCountToken, "3"},
				{explodeCloseToken, "}"},
				{numberToken, "1"},
				{dieToken, "d"},
				{numberToken, "20"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			logger := logger.With("case", tc.input)
			actual := Lex(tc.input, logger).Items()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Logf("%v\n", actual)
				t.Fail()
			}
		})
	}
}
