package main

import (
	"fmt"
	"os"
	"z/lexer"
	"z/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: 00 <source.0>\n")
		os.Exit(1)
	}

	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	source := string(contents)

	l := lexer.New(source)
	l.Tokenize()

	fmt.Printf("\nTokens (%d)\n", len(l.Tokens))
	l.Debug()

	p := parser.New(l.Tokens)
	err = p.Parse()

	if err != nil {
		fmt.Println("\nParsing error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nParse Tree \n")
	p.Debug()
}
