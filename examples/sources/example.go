package main

import (
	"fmt"

	"github.com/samber/oops"
)

// go run examples/sources/example.go | jq .error.sources -r

func f() error { _ = "STUB: not implemented"; return nil }

func e() error { _ = "STUB: not implemented"; return nil }

func d() error { _ = "STUB: not implemented"; return nil }

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	oops.SourceFragmentsHidden = false

	err := a()
	fmt.Println(err.(oops.OopsError).Sources()) //nolint:errcheck,forcetypeassert
}
