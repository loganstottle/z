package parser

import (
	"errors"
	"fmt"
	"z/lexer"
)

// E -> T [+|- T]*
// T -> F | F*T | F/T
// F -> literal | identifier | (E) | -F

func (p *Parser) parse_factor() (Node, error) {
	var result Node

	tok := p.consume()

	if tok.Kind == lexer.LITERAL_NUMBER {
		return Node{FACTOR_LITERAL_NUMBER, tok.Value, make([]Node, 0)}, nil
	} else if tok.Kind == lexer.LITERAL_STRING {
		return Node{FACTOR_LITERAL_STRING, tok.Value, make([]Node, 0)}, nil
	} else if tok.Kind == lexer.IDENTIFIER {
		if p.peek().Kind != lexer.SYMBOL_OPEN_PAREN {
			return Node{FACTOR_IDENTIFIER, tok.Value, make([]Node, 0)}, nil
		}

		p.consume()

		args := []Node{}

		if p.peek().Kind == lexer.SYMBOL_CLOSE_PAREN {
			p.consume()

			return Node{FACTOR_FUNCTION_CALL, tok.Value, args}, nil
		}

		for {
			arg, err := p.parse_expr()
			if err != nil {
				return Node{}, err
			}

			args = append(args, arg)

			if p.peek().Kind == lexer.SYMBOL_CLOSE_PAREN {
				p.consume()
				break
			} else if p.peek().Kind == lexer.SYMBOL_COMMA {
				p.consume()
			}
		}

		return Node{FACTOR_FUNCTION_CALL, tok.Value, args}, nil
	} else if tok.Kind == lexer.SYMBOL_OPEN_PAREN {
		E, err := p.parse_expr()

		if err != nil {
			return result, err
			// return nil, errors.New("expected expression after \"(\"")
		}

		p.expect(lexer.SYMBOL_CLOSE_PAREN)

		return E, nil
	} else if tok.Kind == lexer.SYMBOL_MINUS {
		F, err := p.parse_factor()

		if err != nil {
			return result, err
			// return nil, errors.new("expected expression after \"-\"")
		}

		return Node{FACTOR_NEGATE, "", []Node{F}}, nil
	} else {
		return Node{}, errors.New(fmt.Sprintf("Error: unexpected \"%s\", expected expression", lexer.Token_kind_to_str(tok.Kind)))
	}
}

func (p *Parser) parse_term() (Node, error) {
	result, err := p.parse_factor()

	if err != nil {
		return result, err
	}

	if p.inbounds() {
		tok := p.peek()

		if tok.Kind == lexer.SYMBOL_STAR {
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_MULTIPLY, "", []Node{result, T2}}
		} else if tok.Kind == lexer.SYMBOL_SLASH {
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

		if tok.Kind == lexer.SYMBOL_PLUS {
			T1 := result
			p.consume()
			T2, err := p.parse_term()

			if err != nil {
				return result, err
			}

			result = Node{EXPRESSION_ADD, "", []Node{T1, T2}}
		} else if tok.Kind == lexer.SYMBOL_MINUS {
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
