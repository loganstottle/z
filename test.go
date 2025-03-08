package main

import (
	"reflect"
)

func tokenize_assert(source string, expected []Token) {
	l := lexer_new(source)
	l.tokenize()

	if !reflect.DeepEqual(expected, l.tokens) {
		panic("tokenization test failed")
	}
}

func test_tokenizer() {
	tokenize_assert("x := 1", []Token{{IDENTIFIER, "x"}, {SYMBOL_COLON, ""}, {SYMBOL_EQUALS, ""}, {LITERAL_NUMBER, "1"}, {EOF, ""}})
	tokenize_assert("const e=2.718", []Token{{KEYWORD_CONST, ""}, {IDENTIFIER, "e"}, {SYMBOL_EQUALS, ""}, {LITERAL_NUMBER, "2.718"}, {EOF, ""}})
	tokenize_assert("let a: string = \"happy birthday!\" ;", []Token{{KEYWORD_LET, ""}, {IDENTIFIER, "a"}, {SYMBOL_COLON, ""}, {TYPE_STRING, ""}, {SYMBOL_EQUALS, ""}, {LITERAL_STRING, "happy birthday!"}, {SYMBOL_SEMI_COLON, ""}, {EOF, ""}})
	tokenize_assert("const z:number= 125;", []Token{{KEYWORD_CONST, ""}, {IDENTIFIER, "z"}, {SYMBOL_COLON, ""}, {TYPE_NUMBER, ""}, {SYMBOL_EQUALS, ""}, {LITERAL_NUMBER, "125"}, {SYMBOL_SEMI_COLON, ""}, {EOF, ""}})
}

// todo: test parser
