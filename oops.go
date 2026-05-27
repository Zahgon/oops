package oops

import (
	"context"
	"net/http"
	"time"
)

// Wrap wraps an error into an `oops.OopsError` object that satisfies `error`.
func Wrap(err error) error { _ = "STUB: not implemented"; return nil }

// Wrapf wraps an error into an `oops.OopsError` object that satisfies `error` and formats an error message.
func Wrapf(err error, format string, args ...any) error { _ = "STUB: not implemented"; return nil }

// New returns `oops.OopsError` object that satisfies `error`.
func New(message string) error { _ = "STUB: not implemented"; return nil }

// Errorf formats an error and returns `oops.OopsError` object that satisfies `error`.
func Errorf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

func FromContext(ctx context.Context) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

func Join(e ...error) error { _ = "STUB: not implemented"; return nil }

// Recover handle panic and returns `oops.OopsError` object that satisfies `error`.
func Recover(cb func()) (err error) { _ = "STUB: not implemented"; return nil }

// Recoverf handle panic and returns `oops.OopsError` object that satisfies `error` and formats an error message.
func Recoverf(cb func(), msg string, args ...any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Assert panics if condition is false. Panic payload will be of type oops.OopsError.
// Assertions can be chained.
func Assert(condition bool) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Assertf panics if condition is false. Panic payload will be of type oops.OopsError.
// Assertions can be chained.
func Assertf(condition bool, msg string, args ...any) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Code set a code or slug that describes the error.
// Error messages are intended to be read by humans, but such code is expected to
// be read by machines and even transported over different services.
func Code(code any) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// Time set the error time.
// Default: `time.Now()`.
func Time(time time.Time) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Since set the error duration.
func Since(time time.Time) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Duration set the error duration.
func Duration(duration time.Duration) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// In set the feature category or domain.
func In(domain string) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// Tags adds multiple tags, describing the feature returning an error.
func Tags(tags ...string) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Trace set a transaction id, trace id or correlation id...
func Trace(trace string) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// Span represents a unit of work or operation.
func Span(span string) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// With supplies a list of attributes declared by pair of key+value.
func With(kv ...any) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// WithContext supplies a list of values declared in context.
func WithContext(ctx context.Context, keys ...any) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Hint set a hint for faster debugging.
func Hint(hint string) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// Public sets a message that is safe to show to an end user.
func Public(public string) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Owner set the name/email of the colleague/team responsible for handling this error.
// Useful for alerting purpose.
func Owner(owner string) OopsErrorBuilder { _ = "STUB: not implemented"; return *new(OopsErrorBuilder) }

// User supplies a user id with optional attributes.
// Attributes can be provided as alternating string key/value pairs,
// map[string]any values, and slog.Attr values.
func User(userID string, data ...any) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Tenant supplies a tenant id with optional attributes.
// Attributes can be provided as alternating string key/value pairs,
// map[string]any values, and slog.Attr values.
func Tenant(tenantID string, data ...any) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Request supplies a http.Request.
func Request(req *http.Request, withBody bool) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// Response supplies a http.Response.
func Response(res *http.Response, withBody bool) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// CallerSkip sets the number of additional callers to skip when capturing
// the stack trace. This is useful when oops is wrapped in helper functions.
func CallerSkip(skip int) OopsErrorBuilder {
	_ = "STUB: not implemented"
	return *new(OopsErrorBuilder)
}

// FrameSkip registers a frame filter that excludes matching frames from stack trace
// output. Both file and fun are matched using strings.Contains against the raw values
// returned by runtime.CallersFrames — the absolute file path and the fully-qualified
// function name respectively. An empty string matches anything (i.e., acts as a wildcard).
//
// Filtering is applied at output time (when Stacktrace(), Sources(), or StackFrames()
// is called), not at error creation time. This means patterns registered after an error
// is created will still filter frames from that error's stack trace.
//
// Calling FrameSkip with the same (file, fun) pair more than once is a no-op —
// duplicate entries are silently ignored.
//
// Example:
//
//	oops.FrameSkip("myproject/pkg/errutil", "")          // match by file path substring
//	oops.FrameSkip("", "myproject/pkg/errutil.WrapErr")  // match by full function name substring
func FrameSkip(file string, fun string) { _ = "STUB: not implemented"; return }

// GetPublic returns a message that is safe to show to an end user, or a default generic message.
func GetPublic(err error, defaultPublicMessage string) string { _ = "STUB: not implemented"; return "" }
