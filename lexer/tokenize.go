package lexer

import (
	"unicode"
)

func (l *Lexer) tokenize_ident_or_keyword() {
	ident_buf := ""

	for l.inbounds() && (unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_') {
		ident_buf += string(l.consume())
	}

	tok := Token{Token_str_to_kind(ident_buf), ""}

	if tok.Kind == IDENTIFIER {
		tok.Value = ident_buf
	}

	l.Tokens = append(l.Tokens, tok)
}

func (l *Lexer) tokenize_number_literal() {
	number_literal_buf := ""

	for l.inbounds() && (unicode.IsDigit(l.peek()) || l.peek() == '.') {
		number_literal_buf += string(l.consume())
	}

	l.Tokens = append(l.Tokens, Token{LITERAL_NUMBER, number_literal_buf})
}

func (l *Lexer) tokenize_string_literal() {
	string_literal_buf := ""

	l.consume()

	for l.inbounds() && l.peek() != rune('"') {
		string_literal_buf += string(l.consume())
	}

	l.consume()

	l.Tokens = append(l.Tokens, Token{LITERAL_STRING, string_literal_buf})
}

func is_symbol(ch rune) bool {
	is_letter := unicode.IsLetter(ch)
	is_digit := unicode.IsDigit(ch)
	is_space := unicode.IsSpace(ch)

	return !is_space && !is_letter && !is_digit && ch != '"'
}

func (l *Lexer) tokenize_symbol() {
	symbol1 := string(l.consume())
	kind1 := Token_str_to_kind(symbol1)

	if kind1 == IDENTIFIER {
		kind1 = SYMBOL_INVALID
	}

	if l.inbounds() && is_symbol(l.peek()) {
		symbol2 := string(l.consume())
		kind2 := Token_str_to_kind(symbol2)

		combined_symbols := symbol1 + symbol2
		combined_kind := Token_str_to_kind(combined_symbols)

		if combined_kind != IDENTIFIER {
			l.Tokens = append(l.Tokens, Token{combined_kind, ""})
			return
		} else if kind2 == IDENTIFIER {
			kind2 = SYMBOL_INVALID
		}

		l.Tokens = append(l.Tokens, Token{kind1, ""})
		l.Tokens = append(l.Tokens, Token{kind2, ""})
	} else {
		l.Tokens = append(l.Tokens, Token{kind1, ""})
	}
}

func (l *Lexer) handle_whitespace() {
	l.consume()

	// if l.source[l.i] == '\n' {
	//   fmt.Printf("new line encountered\n")
	// }
}

func (l *Lexer) Tokenize() {
	// todo:
	//   fix number literals (e.g. "1..10")
	//   multi-character symbols (i.e. <=, &&, +=, ++, etc)
	//   comments
	//   better errors
	//   token_new (?)

	for l.inbounds() {
		ch := l.peek()

		is_letter := unicode.IsLetter(ch)
		is_digit := unicode.IsDigit(ch)
		is_space := unicode.IsSpace(ch)

		if is_letter || ch == '_' {
			l.tokenize_ident_or_keyword()
		} else if is_digit {
			l.tokenize_number_literal()
		} else if ch == '"' {
			l.tokenize_string_literal()
		} else if !is_space && !is_letter && !is_digit {
			l.tokenize_symbol()
		} else if is_space {
			l.handle_whitespace()
		}
	}

	l.Tokens = append(l.Tokens, Token{EOF, ""})
}
