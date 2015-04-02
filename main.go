package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Usage: arithmetic <program>")
		os.Exit(1)
	}

	result, err := Evaluate(strings.NewReader(flag.Arg(0)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(result)
}
