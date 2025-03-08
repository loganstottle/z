package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type NodeKind int

const (
	LITERAL_NUMBER_FACTOR NodeKind = iota
	IDENTIFIER_FACTOR
	EXPRESSION_FACTOR
	NEGATION_FACTOR

	FACTOR_LITERAL_NUMBER
	FACTOR_IDENTIFIER
	FACTOR_EXPRESSION
	FACTOR_NEGATE

	EXPRESSION_ADD
	EXPRESSION_SUBTRACT
	EXPRESSION_MULTIPLY
	EXPRESSION_DIVIDE

	ROOT
	END
)

type Node struct {
	kind     NodeKind
	value    string
	children []Node
}

type Parser struct {
	tokens []Token
	i      int
}

func parser_new(tokens []Token) Parser {
	return Parser{tokens, 0}
}

func (p *Parser) inbounds() bool {
	return p.i < len(p.tokens)
}

func (p *Parser) peek() Token {
	return p.tokens[p.i]
}

func (p *Parser) consume() Token {
	tok := p.peek()
	p.i++
	return tok
}

func (p *Parser) expect(kind TokenKind) {
	if p.peek().kind == kind {
		p.consume()
	} else {
		fmt.Printf("Expected %s\n", token_kind_to_str(kind))
		//fmt.Printf("%s:%d:%d Syntax Error: Expected %s\n", token_kind_to_str(kind), filename, p.line, p.col)
		os.Exit(1)
	}
}

// E -> T [+|- T]*
// T -> F | F*T | F/T
// F -> literal | identifier | (E) | -F

func (p *Parser) parse_factor() (Node, error) {
	var result Node

	tok := p.consume()

	if tok.kind == LITERAL_NUMBER {
		return Node{FACTOR_LITERAL_NUMBER, tok.value, make([]Node, 0)}, nil
	} else if tok.kind == IDENTIFIER {
		return Node{FACTOR_IDENTIFIER, tok.value, make([]Node, 0)}, nil
	} else if tok.kind == SYMBOL_OPEN_PAREN {
		E, err := p.parse_expr()

		if err != nil {
			return result, err
			// return nil, errors.New("expected expression after \"(\"")
		}

		p.expect(SYMBOL_CLOSE_PAREN)

		return E, nil
	} else if tok.kind == SYMBOL_MINUS {
		F, err := p.parse_factor()

		if err != nil {
			return result, err
			// return nil, errors.new("expected expression after \"-\"")
		}

		return Node{FACTOR_NEGATE, "", []Node{F}}, nil
	} else if tok.kind == EOF {
		return Node{END, "", []Node{}}, nil
	} else {
		fmt.Println("parsing failed!")
		return Node{}, errors.New("unknown factor")
	}
}

func (p *Parser) parse_term() (Node, error) {
	result, err := p.parse_factor()

	if err != nil {
		return result, err
	}

	if p.inbounds() {
		tok := p.peek()

		if tok.kind == SYMBOL_STAR {
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_MULTIPLY, "", []Node{result, T2}}
		} else if tok.kind == SYMBOL_SLASH {
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_DIVIDE, "", []Node{result, T2}}
		}
	}

	return result, nil
}

func (p *Parser) parse_expr() (Node, error) {
	result, err := p.parse_term()

	if err != nil {
		return result, err
	}

	for p.inbounds() {
		tok := p.peek()

		if tok.kind == SYMBOL_PLUS {
			T1 := result
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_ADD, "", []Node{T1, T2}}
		} else if tok.kind == SYMBOL_MINUS {
			T1 := result
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_SUBTRACT, "", []Node{T1, T2}}
		} else {
			break
		}
	}

	return result, nil
}

func (p *Parser) parse() (Node, error) {
	result := Node{ROOT, "", []Node{}}

	for p.inbounds() {
		E, err := p.parse_expr()
		if err != nil {
			return result, err
		}

		if E.kind != END {
			result.children = append(result.children, E)
		}
	}

	return result, nil
}

func debug_tree(node Node, depth int) {
	output := "  " + strings.Repeat("-", depth*2)

	switch node.kind {
	case FACTOR_LITERAL_NUMBER:
		output += " Number Literal"
	case FACTOR_IDENTIFIER:
		output += " Identifier"
	case FACTOR_NEGATE:
		output += " Negate"
	case EXPRESSION_ADD:
		output += " Add"
	case EXPRESSION_SUBTRACT:
		output += " Subtract"
	case EXPRESSION_MULTIPLY:
		output += " Multiply"
	case EXPRESSION_DIVIDE:
		output += " Divide"
	case ROOT:
		output += " Root"
	}

	if node.value != "" {
		output += fmt.Sprintf(" (%s)", node.value)
	}

	fmt.Println(output)

	for _, n := range node.children {
		debug_tree(n, depth+1)
	}
}
