package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: 00 <source.0>")
		os.Exit(1)
	}

	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading file.\n")
		os.Exit(1)
	}

	source := string(contents)
	fmt.Printf("Tokenizing: \n\n`%s`\n\n", source)

	l := lexer_new(source)
	l.tokenize()
	l.debug()

	fmt.Printf("\nNumber of tokens: %d\n", len(l.tokens))
}
