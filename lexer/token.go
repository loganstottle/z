package lexer

type TokenKind int

const (
	IDENTIFIER TokenKind = iota

	KEYWORD_LET
	KEYWORD_CONST
	KEYWORD_FN
	KEYWORD_FOR
	KEYWORD_WHILE
	KEYWORD_RETURN
	KEYWORD_IF
	KEYWORD_ELSE

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
	SYMBOL_STAR
	SYMBOL_SLASH
	SYMBOL_AND
	SYMBOL_OR
	SYMBOL_NOT
	SYMBOL_NOT_EQUALS
	SYMBOL_GREATER_OR_EQUAL
	SYMBOL_LESS_OR_EQUAL
	SYMBOL_ARROW
	SYMBOL_INVALID

	EOF
)

type Token struct {
	Kind  TokenKind
	Value string
}

func Token_str_to_kind(symbol string) TokenKind {
	token_symbols := map[string]TokenKind{
		"let":    KEYWORD_LET,
		"const":  KEYWORD_CONST,
		"fn":     KEYWORD_FN,
		"for":    KEYWORD_FOR,
		"while":  KEYWORD_WHILE,
		"return": KEYWORD_RETURN,
		"if":     KEYWORD_IF,
		"else":   KEYWORD_ELSE,
		"num":    TYPE_NUMBER,
		"str":    TYPE_STRING,
		"bool":   TYPE_BOOL,
		"true":   LITERAL_BOOL,
		"false":  LITERAL_BOOL,
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
		"*":      SYMBOL_STAR,
		"/":      SYMBOL_SLASH,
		"&&":     SYMBOL_AND,
		"||":     SYMBOL_OR,
		"!":      SYMBOL_NOT,
		"!=":     SYMBOL_NOT_EQUALS,
		">=":     SYMBOL_GREATER_OR_EQUAL,
		"<=":     SYMBOL_LESS_OR_EQUAL,
		"->":     SYMBOL_ARROW,
	}

	kind := token_symbols[symbol]

	if kind != 0 {
		return kind
	} else {
		return IDENTIFIER
	}
}

func Token_kind_to_str(kind TokenKind) string {
	token_names := map[TokenKind]string{
		IDENTIFIER:              "identifier",
		KEYWORD_LET:             "let",
		KEYWORD_CONST:           "const",
		KEYWORD_FN:              "fn",
		KEYWORD_FOR:             "for",
		KEYWORD_RETURN:          "return",
		KEYWORD_IF:              "if",
		KEYWORD_ELSE:            "else",
		TYPE_NUMBER:             "number type",
		TYPE_STRING:             "string type",
		TYPE_BOOL:               "bool type",
		LITERAL_NUMBER:          "number literal",
		LITERAL_STRING:          "string literal",
		LITERAL_BOOL:            "bool literal",
		SYMBOL_OPEN_PAREN:       "(",
		SYMBOL_CLOSE_PAREN:      ")",
		SYMBOL_PLUS:             "+",
		SYMBOL_MINUS:            "-",
		SYMBOL_STAR:             "*",
		SYMBOL_SLASH:            "/",
		SYMBOL_LESS_THAN:        "<",
		SYMBOL_GREATER_THAN:     ">",
		SYMBOL_OPEN_BRACE:       "{",
		SYMBOL_CLOSE_BRACE:      "}",
		SYMBOL_OPEN_BRACKET:     "[",
		SYMBOL_CLOSE_BRACKET:    "]",
		SYMBOL_PERIOD:           ".",
		SYMBOL_COMMA:            ",",
		SYMBOL_EQUALS:           "=",
		SYMBOL_COLON:            ":",
		SYMBOL_SEMI_COLON:       ";",
		SYMBOL_AND:              "&&",
		SYMBOL_OR:               "||",
		SYMBOL_NOT:              "!",
		SYMBOL_NOT_EQUALS:       "!=",
		SYMBOL_GREATER_OR_EQUAL: ">=",
		SYMBOL_LESS_OR_EQUAL:    "<=",
		SYMBOL_ARROW:            "->",
		SYMBOL_INVALID:          "invalid symbol",
		EOF:                     "EOF",
	}

	return token_names[kind]
}
