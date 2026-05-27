package main

import (
	oopszap "github.com/samber/oops/loggers/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go run examples/zap/example.go

func d() error { _ = "STUB: not implemented"; return nil }

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build()
	defer func() { _ = logger.Sync() }()

	err := a()

	logger.Error(err.Error(),
		zap.Object("error", oopszap.OopsMarshalFunc(err)),
		zap.String("stacktrace", oopszap.OopsStackMarshaller(err)),
	)
}
