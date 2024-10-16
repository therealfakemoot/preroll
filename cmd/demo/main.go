package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/therealfakemoot/preroll/lexer"
)

func main() {
	var input string

	flag.StringVar(&input, "roll", "1d20+3", "dice roll")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	l := lexer.Lex(input, logger.WithGroup("lexer"))
	logger = logger.WithGroup("main").With("input", input)
	logger.Info("lexing input")
	for t := range l.Items() {
		logger.With("token", t).Debug("found token")
	}
}
