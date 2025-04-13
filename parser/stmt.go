package parser

import (
	"errors"
	"fmt"
	"os"
	"z/lexer"
)

func (p *Parser) parse_declaration() (Node, error) {
	let_const := p.consume().Kind
	name := p.expect(lexer.IDENTIFIER)
	p.expect(lexer.SYMBOL_COLON)

	kind := p.consume()
	typ := Node{}

	if kind.Kind == lexer.KEYWORD_CONST {
		fmt.Println("Error: constants must be initialized")
		os.Exit(1)
	} else if kind.Kind == lexer.TYPE_NUMBER {
		if kind.Kind == lexer.SYMBOL_SEMI_COLON {
			p.consume()
			return Node{STATEMENT_LET_DECLARATION, "", []Node{{FACTOR_IDENTIFIER, name.Value, []Node{}}, Node{FACTOR_LITERAL_NUMBER, "0", []Node{}}}}, nil
		}

		typ.Kind = TYPE_NUMBER
	} else if kind.Kind == lexer.TYPE_STRING {
		if kind.Kind == lexer.SYMBOL_SEMI_COLON {
			p.consume()
			return Node{STATEMENT_LET_DECLARATION, "", []Node{{FACTOR_IDENTIFIER, name.Value, []Node{}}, Node{FACTOR_LITERAL_STRING, "0", []Node{}}}}, nil
		}

		typ.Kind = TYPE_STRING
	} else {
		fmt.Printf("unknown type: %s\n", lexer.Token_kind_to_str(kind.Kind))
		os.Exit(1)
	}

	p.expect(lexer.SYMBOL_EQUALS)
	E, err := p.parse_expr()
	if err != nil {
		fmt.Printf("\nError: invalid expression following \"=\"\n")
		// fmt.Printf("p.parse_expr ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	p.expect(lexer.SYMBOL_SEMI_COLON)

	if let_const == lexer.KEYWORD_LET {
		return Node{STATEMENT_LET_DECLARATION, name.Value, append([]Node{typ}, E)}, nil
	} else {
		return Node{STATEMENT_CONST_DECLARATION, name.Value, append([]Node{typ}, E)}, nil
	}
}

func (p *Parser) parse_variable_set() (Node, error) {
	name := p.expect(lexer.IDENTIFIER)
	p.expect(lexer.SYMBOL_EQUALS)
	E, err := p.parse_expr()
	if err != nil {
		return Node{}, err
	}

	p.expect(lexer.SYMBOL_SEMI_COLON)
	return Node{STATEMENT_LET_ASSIGN, name.Value, []Node{E}}, nil
}

func (p *Parser) parse_return() (Node, error) {
	p.expect(lexer.KEYWORD_RETURN)

	E, err := p.parse_expr()
	if err != nil {
		return Node{}, err
	}

	p.expect(lexer.SYMBOL_SEMI_COLON)

	return Node{STATEMENT_RETURN, "", []Node{E}}, nil
}

func (p *Parser) parse_block() (Node, error) {
	p.expect(lexer.SYMBOL_OPEN_BRACE)

	block := Node{STATEMENT_BLOCK, "", []Node{}}

	for {
		tok := p.peek()

		if tok.Kind == lexer.SYMBOL_OPEN_BRACE {
			B, err := p.parse_block()
			if err != nil {
				return Node{}, err
			}

			block.Children = append(block.Children, B)
		}

		if tok.Kind == lexer.SYMBOL_CLOSE_BRACE {
			p.consume()
			break
		}

		S, err := p.parse_statement()
		if err != nil {
			return Node{}, err
		}

		block.Children = append(block.Children, S)
	}

	return block, nil
}

func (p *Parser) parse_function_declaration() (Node, error) {
	p.expect(lexer.KEYWORD_FN)
	name := p.expect(lexer.IDENTIFIER)
	p.expect(lexer.SYMBOL_OPEN_PAREN)

	params := []Node{}

	if p.peek().Kind == lexer.SYMBOL_CLOSE_PAREN {
		p.expect(lexer.SYMBOL_CLOSE_PAREN)
	} else {
		for {
			if p.peek().Kind != lexer.IDENTIFIER {
				fmt.Println("error: invalid parameter: ", p.peek().Kind)
				os.Exit(1)
			}

			param_name := p.consume()

			p.expect(lexer.SYMBOL_COLON)

			typ := p.consume().Kind

			switch typ {
			case lexer.TYPE_NUMBER:
				params = append(params, Node{PARAMETER_TYPE_NUMBER, param_name.Value, []Node{}})
			case lexer.TYPE_STRING:
				params = append(params, Node{PARAMETER_TYPE_STRING, param_name.Value, []Node{}})
			case lexer.TYPE_BOOL:
				params = append(params, Node{PARAMETER_TYPE_BOOL, param_name.Value, []Node{}})
			default:
				fmt.Println("invalid return type following \"->\"")
				os.Exit(1)
			}

			if p.peek().Kind == lexer.SYMBOL_CLOSE_PAREN {
				p.consume()
				break
			} else if p.peek().Kind == lexer.SYMBOL_COMMA {
				p.consume()
			}
		}
	}

	p.expect(lexer.SYMBOL_ARROW)

	return_type := p.consume()

	B, err := p.parse_block()
	if err != nil {
		return Node{}, err
	}

	var ret NodeKind

	switch return_type.Kind {
	case lexer.TYPE_NUMBER:
		ret = RETURN_TYPE_NUMBER
	case lexer.TYPE_STRING:
		ret = RETURN_TYPE_STRING
	case lexer.TYPE_BOOL:
		ret = RETURN_TYPE_BOOL
	default:
		fmt.Println("invalid return type following \"->\"")
		os.Exit(1)
	}

	return Node{STATEMENT_FUNCTION_DECLARATION, name.Value, append([]Node{Node{ret, "", []Node{}}}, append(params, B)...)}, nil
}

func (p *Parser) parse_while_loop() (Node, error) {
	p.expect(lexer.KEYWORD_WHILE)

	E, err := p.parse_expr()
	if err != nil {
		fmt.Println("Error: invalid expression following \"while\"")
		return Node{}, err
	}

	B, err := p.parse_block()
	if err != nil {
		fmt.Println("Error: expected block following while loop expression")
		return Node{}, err
	}

	return Node{STATEMENT_WHILE_LOOP, "", []Node{E, B}}, nil
}

func (p *Parser) parse_conditional() (Node, error) {
	p.expect(lexer.KEYWORD_IF)
	E, err := p.parse_expr()
	if err != nil {
		fmt.Println("Error: invalid expression following \"if\"")
		return Node{}, err
	}

	B, err := p.parse_block()
	if err != nil {
		fmt.Println("Error: expected block following for loop expression")
		return Node{}, err
	}

	if p.inbounds() && p.peek().Kind == lexer.KEYWORD_ELSE {
		p.consume()
		S, err := p.parse_statement()
		if err != nil {
			fmt.Println("Error: invalid statement following \"else\"")
			return Node{}, err
		}

		return Node{STATEMENT_IF, "", []Node{E, B, S}}, nil
	}

	return Node{STATEMENT_IF, "", []Node{E, B}}, nil
}

func (p *Parser) parse_statement() (Node, error) {
	tok := p.peek()

	if tok.Kind == lexer.KEYWORD_LET || tok.Kind == lexer.KEYWORD_CONST {
		declaration, err := p.parse_declaration()
		if err != nil {
			return Node{}, err
		}

		return declaration, nil
	} else if tok.Kind == lexer.IDENTIFIER && p.Tokens[p.i+1].Kind != lexer.SYMBOL_OPEN_PAREN { //todo:fix
		set, err := p.parse_variable_set()
		if err != nil {
			return Node{}, err
		}

		return set, nil
	} else if tok.Kind == lexer.KEYWORD_RETURN {
		R, err := p.parse_return()
		if err != nil {
			return Node{}, err
		}

		return R, nil
	} else if tok.Kind == lexer.SYMBOL_OPEN_BRACE {
		B, err := p.parse_block()
		if err != nil {
			return Node{}, err
		}

		return B, nil
	} else if tok.Kind == lexer.KEYWORD_FN {
		F, err := p.parse_function_declaration()
		if err != nil {
			return Node{}, err
		}

		return F, nil
	} else if tok.Kind == lexer.KEYWORD_WHILE {
		W, err := p.parse_while_loop()
		if err != nil {
			return Node{}, err
		}

		return W, nil
	} else if tok.Kind == lexer.KEYWORD_IF {
		C, err := p.parse_conditional()
		if err != nil {
			return Node{}, err
		}

		return C, nil
	} else {
		E, err := p.parse_expr()
		if err != nil {
			msg := fmt.Sprintf("unknown expression after \"%s\"", lexer.Token_kind_to_str(tok.Kind))

			if tok.Value != "" {
				msg = fmt.Sprintf("unknown expression after \"%s (%s)\"", lexer.Token_kind_to_str(tok.Kind), tok.Value)
			}

			return Node{}, errors.New(msg)
		}

		p.expect(lexer.SYMBOL_SEMI_COLON)

		return E, nil
	}
}
