package oopszap

import (
	"github.com/samber/oops"
	"go.uber.org/zap/zapcore"
)

// OopsStackMarshaller returns the stack trace string for use in zap.
// Usage: zap.String("stacktrace", oopszap.OopsStackMarshaller(err)).
func OopsStackMarshaller(err error) string { _ = "STUB: not implemented"; return "" }

// For normal errors, we might not want to return anything or just empty string,
// but to be safe/useful let's return nothing or handle it at call site.
// Actually, matching zerolog implementation logic:

// OopsMarshalFunc returns a zapcore.ObjectMarshaler that logs the error details.
// Usage: zap.Object("error", oopszap.OopsMarshalFunc(err)).
func OopsMarshalFunc(err error) zapcore.ObjectMarshaler {
	_ = "STUB: not implemented"
	return *new(zapcore.ObjectMarshaler)
}

type simpleErrorMarshaller struct {
	err error
}

func (m *simpleErrorMarshaller) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type zapErrorMarshaller struct {
	err oops.OopsError
}

func (m *zapErrorMarshaller) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip stacktrace in the main object - handled separately if desired
