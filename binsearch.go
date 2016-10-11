package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var flagHex = flag.String("hex", "", "search bytes specified as hex pattern")
var flagStr = flag.String("string", "", "search string bytes")

func main() {
	flag.Parse()

	if flag.NArg() != 1 || (len(*flagHex) == 0 && len(*flagStr) == 0) {
		fmt.Printf("Usage: %s <-hex hexPattern | -string stringPattern> fileToSearch\n", os.Args[0])
		os.Exit(1)
	}

	var hexPattern []byte
	if len(*flagHex) > 0 {
		var err error
		hexPattern, err = hex.DecodeString(*flagHex)
		if err != nil {
			panic(err)
		}
	} else if len(*flagStr) > 0 {
		hexPattern = []byte(*flagStr)
	}

	if len(hexPattern) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Hex pattern must be at least one byte long\n")
		os.Exit(1)
	}

	input, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer input.Close()

	fmt.Fprintln(os.Stderr, "Searching for...", hex.Dump(hexPattern))

	oneByte := make([]byte, 1)
	pos := 0
	matched := 0
	for {
		cnt, err := input.Read(oneByte)
		if cnt == 0 || err != nil {
			break
		}

		if oneByte[0] == hexPattern[matched] {
			matched++
			if matched == len(hexPattern) {
				pos -= len(hexPattern) - 1
				fmt.Printf("%d 0x%x\n", pos, pos)
				os.Exit(0)
			}
		} else {
			matched = 0
		}
		pos++
	}

	fmt.Fprintf(os.Stderr, "not found\n")
	os.Exit(1)
}
