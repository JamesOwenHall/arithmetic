package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	evalArg := flag.String("e", "", "evaluate argument")
	interactive := flag.Bool("i", false, "interactive mode")
	flag.Parse()

	if *interactive {
		repl()
		return
	}

	if *evalArg != "" {
		result, err := Evaluate(strings.NewReader(*evalArg))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(result)
		return
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func repl() {
replloop:
	for {
		fmt.Print("> ")
		byteBuffer := []byte{0x00}
		progBuffer := new(bytes.Buffer)

	readloop:
		for {
			_, err := os.Stdin.Read(byteBuffer)
			if err != nil {
				fmt.Println()
				break replloop
			}
			if byteBuffer[0] == '\n' {
				break readloop
			}
			progBuffer.WriteByte(byteBuffer[0])
		}

		if progBuffer.String() == "exit" {
			break replloop
		}

		result, err := Evaluate(progBuffer)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
