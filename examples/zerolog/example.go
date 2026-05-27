package main

import (
	"os"

	"github.com/rs/zerolog"
	oopszerolog "github.com/samber/oops/loggers/zerolog"
)

// go run examples/zerolog/example.go 2>&1 | jq
// go run examples/zerolog/example.go 2>&1 | jq .stack -r

func d() error { _ = "STUB: not implemented"; return nil }

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	zerolog.ErrorStackMarshaler = oopszerolog.OopsStackMarshaller
	zerolog.ErrorMarshalFunc = oopszerolog.OopsMarshalFunc
	logger := zerolog.
		New(os.Stderr).
		With().
		Timestamp().
		Logger()

	err := a()

	logger.Error().Stack().Err(err).Msg(err.Error())
}
