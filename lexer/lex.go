package lexer

import (
	"fmt"
)

type Lexer struct {
	source string
	i      int
	Tokens []Token
}

func New(source string) Lexer {
	return Lexer{source, 0, make([]Token, 0)}
}

func (l *Lexer) inbounds() bool {
	return l.i < len(l.source)
}

func (l *Lexer) peek() rune {
	return rune(l.source[l.i])
}

func (l *Lexer) consume() rune {
	r := l.peek()
	l.i++

	return r
}

func (l *Lexer) Debug() {
	for _, token := range l.Tokens {
		kind := Token_kind_to_str(token.Kind)
		val := token.Value

		if val != "" {
			fmt.Printf("  { %s \"%s\" }\n", kind, val)
		} else {
			fmt.Printf("  { %s }\n", kind)
		}
	}
}
