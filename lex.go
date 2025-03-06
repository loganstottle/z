package main

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	IDENTIFIER TokenKind = iota

	KEYWORD_LET
	KEYWORD_CONST
	KEYWORD_FN
	KEYWORD_FOR

	TYPE_NUMBER
	TYPE_STRING
	TYPE_BOOL

	LITERAL_NUMBER
	LITERAL_STRING
	LITERAL_BOOL

	SYMBOL_OPEN_PAREN
	SYMBOL_CLOSE_PAREN
	SYMBOL_PLUS
	SYMBOL_MINUS
	SYMBOL_LESS_THAN
	SYMBOL_GREATER_THAN
	SYMBOL_OPEN_BRACE
	SYMBOL_CLOSE_BRACE
	SYMBOL_OPEN_BRACKET
	SYMBOL_CLOSE_BRACKET
	SYMBOL_PERIOD
	SYMBOL_COMMA
	SYMBOL_EQUALS
	SYMBOL_COLON
	SYMBOL_SEMI_COLON
	SYMBOL_INVALID

	EOF
)

type Token struct {
	kind  TokenKind
	value string
}

func token_str_to_kind(symbol string) TokenKind {
	token_symbols := map[string]TokenKind{
		"fn":     KEYWORD_FN,
		"for":    KEYWORD_FOR,
		"number": TYPE_NUMBER,
		"string": TYPE_STRING,
		"bool":   TYPE_BOOL,
		"true":   LITERAL_BOOL,
		"false":  LITERAL_BOOL,
		"const":  KEYWORD_CONST,
		"let":    KEYWORD_LET,
		"(":      SYMBOL_OPEN_PAREN,
		")":      SYMBOL_CLOSE_PAREN,
		"+":      SYMBOL_PLUS,
		"-":      SYMBOL_MINUS,
		"<":      SYMBOL_LESS_THAN,
		">":      SYMBOL_GREATER_THAN,
		"{":      SYMBOL_OPEN_BRACE,
		"}":      SYMBOL_CLOSE_BRACE,
		"[":      SYMBOL_OPEN_BRACKET,
		"]":      SYMBOL_CLOSE_BRACKET,
		".":      SYMBOL_PERIOD,
		",":      SYMBOL_COMMA,
		"=":      SYMBOL_EQUALS,
		":":      SYMBOL_COLON,
		";":      SYMBOL_SEMI_COLON,
	}

	kind := token_symbols[symbol]

	if kind != 0 {
		return kind
	} else {
		return IDENTIFIER
	}
}

func token_kind_to_str(kind TokenKind) string {
	token_names := map[TokenKind]string{
		IDENTIFIER:           "identifier",
		KEYWORD_FN:           "fn keyword",
		KEYWORD_FOR:          "for keyword",
		TYPE_NUMBER:          "number type",
		TYPE_STRING:          "string type",
		TYPE_BOOL:            "boolean type",
		LITERAL_NUMBER:       "number literal",
		LITERAL_STRING:       "string literal",
		LITERAL_BOOL:         "boolean literal",
		KEYWORD_CONST:        "const keyword",
		KEYWORD_LET:          "let keyword",
		SYMBOL_OPEN_PAREN:    "opening parenthesis symbol",
		SYMBOL_CLOSE_PAREN:   "closing parenthesis symbol",
		SYMBOL_PLUS:          "plus symbol",
		SYMBOL_MINUS:         "minus symbol",
		SYMBOL_LESS_THAN:     "less than symbol",
		SYMBOL_GREATER_THAN:  "greater than symbol",
		SYMBOL_OPEN_BRACE:    "opening brace symbol",
		SYMBOL_CLOSE_BRACE:   "closing brace symbol",
		SYMBOL_OPEN_BRACKET:  "opening bracket symbol",
		SYMBOL_CLOSE_BRACKET: "closing bracket symbol",
		SYMBOL_PERIOD:        "period symbol",
		SYMBOL_COMMA:         "comma symbol",
		SYMBOL_EQUALS:        "equals symbol",
		SYMBOL_COLON:         "colon symbol",
		SYMBOL_SEMI_COLON:    "semi colon symbol",
		SYMBOL_INVALID:       "invalid symbol",
		EOF:                  "end of file",
	}

	return token_names[kind]
}

type Lexer struct {
	source string
	i      int
	tokens []Token
}

func lexer_new(source string) Lexer {
	return Lexer{source, 0, make([]Token, 0)}
}

func (l *Lexer) tokenize() {
	// todo:
	//   fix number literals (e.g. "1..10")
	//   abstract l.source[l.i] and l.i++ as peek() and consume()
	//   multi-character symbols (i.e. <=, &&, +=, ++, etc)
	//   comments

	for l.i = 0; l.i < len(l.source); l.i++ {
		if unicode.IsLetter(rune(l.source[l.i])) || l.source[l.i] == '_' {
			ident_buf := ""

			for l.i < len(l.source) && (unicode.IsLetter(rune(l.source[l.i])) || unicode.IsDigit(rune(l.source[l.i])) || l.source[l.i] == '_') {
				ident_buf += string(l.source[l.i])
				l.i++
			}

			l.i--

			tok := Token{token_str_to_kind(ident_buf), ""}

			if tok.kind == IDENTIFIER {
				tok.value = ident_buf
			}

			l.tokens = append(l.tokens, tok)
		} else if unicode.IsNumber(rune(l.source[l.i])) {
			number_literal_buf := ""

			for l.i < len(l.source) && (unicode.IsDigit(rune(l.source[l.i]))) || l.source[l.i] == '.' {
				number_literal_buf += string(l.source[l.i])
				l.i++
			}

			l.i--

			l.tokens = append(l.tokens, Token{LITERAL_NUMBER, number_literal_buf})
		} else if l.source[l.i] == '"' {
			string_literal_buf := ""

			l.i++

			for l.i < len(l.source) && l.source[l.i] != '"' {
				string_literal_buf += string(l.source[l.i])
				l.i++
			}

			l.tokens = append(l.tokens, Token{LITERAL_STRING, string_literal_buf})
		} else if !unicode.IsSpace(rune(l.source[l.i])) && !unicode.IsLetter(rune(l.source[l.i])) && !unicode.IsDigit(rune(l.source[l.i])) {
			kind := token_str_to_kind(string(l.source[l.i]))

			if kind == IDENTIFIER {
				kind = SYMBOL_INVALID
			}

			l.tokens = append(l.tokens, Token{kind, ""})
		} else {
			// white space

			// if l.source[l.i] == '\n' {
			//   fmt.Printf("new line encountered\n")
			// }
		}
	}

	l.tokens = append(l.tokens, Token{EOF, ""})
}

func (l *Lexer) debug() {
	for _, token := range l.tokens {
		kind := token_kind_to_str(token.kind)
		val := token.value

		if val != "" {
			fmt.Printf("[%s ( %s )]\n", kind, val)
		} else {
			fmt.Printf("[%s]\n", kind)
		}
	}
}
