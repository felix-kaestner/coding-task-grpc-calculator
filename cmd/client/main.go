package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/felix-kaestner/calculator/internal/app/client"
)

const (
	usage = `Calculator gRPC client.

Usage:
  client client -method <method> -a <a> -b <b>
  client client -host calculator.lovoo.com -method <method> -a <a> -b <b>

Options:`
)

var (
	host = flag.String("host", "localhost:8000", "The address of the gRPC server.")
	m    = flag.String("method", "", "The method to use, [Available methods: 'add', 'sub', 'mul', 'div']")
	a    = flag.Float64("a", 0, "The left-hand side operand.")
	b    = flag.Float64("b", 0, "The right-hand side operand")
)

func isFlagPassed(flags ...string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		for _, ff := range flags {
			if f.Name == ff {
				found = true
			}
		}
	})
	return found
}

func main() {
	flag.Parse()

	// Check if all required flags are passed.
	if !isFlagPassed("method", "a", "b") {
		fmt.Println(usage)
		flag.PrintDefaults()
		os.Exit(1)
	}

	r, err := client.Solve(context.Background(), *host, *m, *a, *b)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(r)
}