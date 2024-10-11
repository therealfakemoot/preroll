package main

import (
	"log/slog"
	"os"

	"github.com/therealfakemoot/preroll/lexer"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.With("token", lexer.EOF).Info("printing raw token")
}
