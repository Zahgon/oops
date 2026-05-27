package main

import (
	"log"
)

// go run examples/log/example.go

func d() error { _ = "STUB: not implemented"; return nil }

// lazy evaluation

func c() error { _ = "STUB: not implemented"; return nil }

func b() error { _ = "STUB: not implemented"; return nil }

func a() error { _ = "STUB: not implemented"; return nil }

func main() {
	err := a()
	if err != nil {
		log.Printf("%+v", err)
	}
}
