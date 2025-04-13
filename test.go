package main

import (
	"reflect"
	"z/lexer"
)

func tokenize_assert(source string, expected []lexer.Token) {
	l := lexer.New(source)
	l.Tokenize()

	if !reflect.DeepEqual(expected, l.Tokens) {
		panic("lexer.Tokenization test failed")
	}
}

func test_tokenizer() {
	tokenize_assert("x := 1", []lexer.Token{{Kind: lexer.IDENTIFIER, Value: "x"}, {Kind: lexer.SYMBOL_COLON, Value: ""}, {Kind: lexer.SYMBOL_EQUALS, Value: ""}, {Kind: lexer.LITERAL_NUMBER, Value: "1"}, {Kind: lexer.EOF, Value: ""}})
	tokenize_assert("const e=2.718", []lexer.Token{{Kind: lexer.KEYWORD_CONST, Value: ""}, {Kind: lexer.IDENTIFIER, Value: "e"}, {Kind: lexer.SYMBOL_EQUALS, Value: ""}, {Kind: lexer.LITERAL_NUMBER, Value: "2.718"}, {Kind: lexer.EOF, Value: ""}})
	tokenize_assert("let a: str = \"happy birthday!\" ;", []lexer.Token{{Kind: lexer.KEYWORD_LET, Value: ""}, {Kind: lexer.IDENTIFIER, Value: "a"}, {Kind: lexer.SYMBOL_COLON, Value: ""}, {Kind: lexer.TYPE_STRING, Value: ""}, {Kind: lexer.SYMBOL_EQUALS, Value: ""}, {Kind: lexer.LITERAL_STRING, Value: "happy birthday!"}, {Kind: lexer.SYMBOL_SEMI_COLON, Value: ""}, {Kind: lexer.EOF, Value: ""}})
	tokenize_assert("const z:num= 125;", []lexer.Token{{Kind: lexer.KEYWORD_CONST, Value: ""}, {Kind: lexer.IDENTIFIER, Value: "z"}, {Kind: lexer.SYMBOL_COLON, Value: ""}, {Kind: lexer.TYPE_NUMBER, Value: ""}, {Kind: lexer.SYMBOL_EQUALS, Value: ""}, {Kind: lexer.LITERAL_NUMBER, Value: "125"}, {Kind: lexer.SYMBOL_SEMI_COLON, Value: ""}, {Kind: lexer.EOF, Value: ""}})
}

// todo: test parser
