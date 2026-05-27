package main

import (
	oopslogrus "github.com/samber/oops/loggers/logrus"
	"github.com/sirupsen/logrus"
)

// go run examples/segfault/example.go 2>&1 | jq
// go run examples/segfault/example.go 2>&1 | jq .stacktrace -r

func nilPointerException() { _ = "STUB: not implemented"; return }

func handlePanic() error { _ = "STUB: not implemented"; return nil }

// ...

// ...

func main() {
	logrus.SetFormatter(oopslogrus.NewOopsFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	}))

	err := handlePanic()
	if err != nil {
		logrus.WithError(err).Error(err)
	}
}
