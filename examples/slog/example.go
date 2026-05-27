package main

import (
	"log/slog"
	"os"
)

// go run examples/slog/example.go | jq
// go run examples/slog/example.go | jq .error.stacktrace -r

func d() error { _ = "STUB: not implemented"; return nil }

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := a()
	if err != nil {
		logger.Error(
			err.Error(),
			slog.Any("error", err),
		)
	}
}
