package main

import (
	"log/slog"
	"os"

	"github.com/therealfakemoot/preroll/lexer"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	tokens, err := lexer.Lex(`1d20`, logger)
	if err != nil {
		logger.With("error", err).Error("error during lexing")
		os.Exit(1)
	}

	logger.With("lexed tokens", tokens).Info("tokenization results")
}
