package cmd

import (
	"log/slog"
	"os"
)

// initLogger initializes the default logger for the application using slog.
// It does not take any parameters.
// It returns an error if the logger initialization fails, although in this implementation, it always returns nil.
func initLogger(arg *args) error {
	var logLever slog.Level
	if err := logLever.UnmarshalText([]byte(arg.LogLevel)); err != nil {
		return err
	}

	options := &slog.HandlerOptions{
		Level: logLever,
	}

	var logHandler slog.Handler
	if arg.TextFormat {
		logHandler = slog.NewTextHandler(os.Stdout, options)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, options)
	}

	logger := slog.New(logHandler).With(
		slog.String("ver", arg.version),
		slog.String("app", "mcp-code-tools"),
	)

	slog.SetDefault(logger)

	return nil
}
