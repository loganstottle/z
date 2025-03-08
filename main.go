package main

import (
	"fmt"
	"os"
)

func main() {
	test_tokenizer()

	if len(os.Args) != 2 {
		fmt.Printf("Usage: 00 <source.0>\n")
		os.Exit(1)
	}

	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading file.\n")
		os.Exit(1)
	}

	source := string(contents)

	l := lexer_new(source)
	l.tokenize()

	fmt.Printf("\nTokens (%d)\n", len(l.tokens))

	l.debug()

	p := parser_new(l.tokens)
	result, err := p.parse()

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nParse Tree \n")

	debug_tree(result, 1)
}
