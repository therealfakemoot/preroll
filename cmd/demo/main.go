package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/therealfakemoot/preroll/lexer"
)

func main() {
	var (
		input    string
		level    string
		logLevel slog.Level
	)

	flag.StringVar(&input, "roll", "1d20+3", "dice roll")
	flag.StringVar(&level, "level", "INFO", "logging level: debug|DEBUG, info|INFO")
	flag.Parse()

	switch {
	case level == "debug" || level == "DEBUG":
		logLevel = slog.LevelDebug
	case level == "info" || level == "INFO":
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	l := lexer.Lex(input, logger.WithGroup("lexer"))
	logger = logger.WithGroup("main").With("input", input)
	logger.Info("lexing input")
	for _, t := range l.Items() {
		logger.With("token", t).Info("found token")
	}
}
