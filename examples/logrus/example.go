package main

import (
	oopslogrus "github.com/samber/oops/loggers/logrus"
	"github.com/sirupsen/logrus"
)

// go run examples/logrus/example.go 2>&1 | jq
// go run examples/logrus/example.go 2>&1 | jq .stacktrace -r

func d() error { _ = "STUB: not implemented"; return nil }

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	logrus.SetFormatter(oopslogrus.NewOopsFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	}))

	err := a()
	if err != nil {
		logrus.WithError(err).Error(err)
	}
}
