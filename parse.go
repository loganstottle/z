package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type NodeKind int

const (
	ROOT NodeKind = iota

	FACTOR_LITERAL_NUMBER
	FACTOR_IDENTIFIER
	FACTOR_EXPRESSION
	FACTOR_NEGATE
	FACTOR_FUNCTION_CALL

	EXPRESSION_ADD
	EXPRESSION_SUBTRACT
	EXPRESSION_MULTIPLY
	EXPRESSION_DIVIDE

	STATEMENT_LET_DECLARATION
	STATEMENT_CONST_DECLARATION
	STATEMENT_LET_ASSIGN

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
	root   Node
}

func parser_new(tokens []Token) Parser {
	return Parser{tokens, 0, Node{}}
}

func (p *Parser) inbounds() bool {
	return p.i < len(p.tokens) && p.peek().kind != EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.i]
}

func (p *Parser) consume() Token {
	tok := p.peek()
	p.i++
	return tok
}

func (p *Parser) expect(kind TokenKind) Token {
	if p.peek().kind == kind {
		return p.consume()
	} else {
		fmt.Printf("\nError: Expected %s\n", token_kind_to_str(kind))
		//fmt.Printf("%s:%d:%d Syntax Error: Expected %s\n", token_kind_to_str(kind), filename, p.line, p.col)
		fmt.Println("\nParse tree")
		debug_tree(p.root, 1)
		os.Exit(1)
		return Token{}
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
		if p.peek().kind != SYMBOL_OPEN_PAREN {
			return Node{FACTOR_IDENTIFIER, tok.value, make([]Node, 0)}, nil
		}

		p.consume()

		args := []Node{}

		if p.peek().kind == SYMBOL_CLOSE_PAREN {
			p.consume()

			return Node{FACTOR_FUNCTION_CALL, tok.value, args}, nil
		}

		for {
			arg, err := p.parse_expr()
			if err != nil {
				return Node{}, err
			}

			args = append(args, arg)

			if p.peek().kind == SYMBOL_CLOSE_PAREN {
				p.consume()
				break
			} else if p.peek().kind == SYMBOL_COMMA {
				p.consume()
			}
		}

		return Node{FACTOR_FUNCTION_CALL, tok.value, args}, nil
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

func (p *Parser) parse_declaration() (Node, error) {
	tok := p.peek()

	if tok.kind == KEYWORD_LET {
		p.expect(KEYWORD_LET)
		name := p.expect(IDENTIFIER)
		p.expect(SYMBOL_COLON)

		kind := p.consume()

		if kind.kind == TYPE_NUMBER {
			if p.peek().kind == SYMBOL_SEMI_COLON {
				// let a: number;
				p.consume()
				return Node{STATEMENT_LET_DECLARATION, "", []Node{{FACTOR_IDENTIFIER, name.value, []Node{}}, Node{FACTOR_LITERAL_NUMBER, "0", []Node{}}}}, nil
			}

			p.expect(SYMBOL_EQUALS)
			E, err := p.parse_expr()
			if err != nil {
				fmt.Printf("invalid expression following =\n")
				fmt.Printf("p.parse_expr ERROR: %s\n", err.Error())
				os.Exit(1)
			}

			p.expect(SYMBOL_SEMI_COLON)

			return Node{STATEMENT_LET_DECLARATION, "", []Node{{FACTOR_IDENTIFIER, name.value, []Node{}}, E}}, nil
		} else {
			fmt.Printf("unknown variable type: %s\n", token_kind_to_str(kind.kind))
			os.Exit(1)
		}
	} else if tok.kind == KEYWORD_CONST {
		p.expect(KEYWORD_CONST)
		name := p.expect(IDENTIFIER)
		p.expect(SYMBOL_COLON)

		kind := p.consume()

		if kind.kind == TYPE_NUMBER {
			if p.peek().kind == SYMBOL_SEMI_COLON {
				fmt.Println("constants must be initialized")
				os.Exit(1)
			}

			p.expect(SYMBOL_EQUALS)
			E, err := p.parse_expr()
			if err != nil {
				fmt.Printf("invalid expression following =\n")
				fmt.Printf("p.parse_expr ERROR: %s\n", err.Error())
				os.Exit(1)
			}

			p.expect(SYMBOL_SEMI_COLON)

			return Node{STATEMENT_CONST_DECLARATION, "", []Node{{FACTOR_IDENTIFIER, name.value, []Node{}}, E}}, nil
		} else {
			fmt.Printf("unknown constant type: %s\n", token_kind_to_str(kind.kind))
			os.Exit(1)
		}
	} else {
		fmt.Println("unknown statement")
		os.Exit(1)
	}

	return Node{}, nil
}

func (p *Parser) parse_variable_set() (Node, error) {
	name := p.expect(IDENTIFIER)
	p.expect(SYMBOL_EQUALS)
	E, err := p.parse_expr()
	if err != nil {
		return Node{}, err
	}

	p.expect(SYMBOL_SEMI_COLON)
	return Node{STATEMENT_LET_ASSIGN, name.value, []Node{E}}, nil
}

func (p *Parser) parse_statement() (Node, error) {
	tok := p.peek()

	if tok.kind == KEYWORD_LET || tok.kind == KEYWORD_CONST {
		declaration, err := p.parse_declaration()
		if err != nil {
			return Node{}, err
		}

		return declaration, nil
	} else if tok.kind == IDENTIFIER {
		set, err := p.parse_variable_set()
		if err != nil {
			return Node{}, err
		}

		return set, nil
	} else {
		msg := fmt.Sprintf("failed to parse statement [@%s]\n", token_kind_to_str(tok.kind))

		if tok.value != "" {
			msg = fmt.Sprintf("failed to parse statement @ %s (%s)\n", token_kind_to_str(tok.kind), tok.value)
		}

		return Node{}, errors.New(msg)
	}
}

func (p *Parser) parse() error {
	// todo:
	//   strings
	//   bools
	//   more and more statements
	//     (if/while/for/fn etc)

	p.root = Node{ROOT, "", []Node{}}

	for p.inbounds() {
		S, err := p.parse_statement()

		if err != nil {
			return err
		}

		p.root.children = append(p.root.children, S)
	}

	return nil
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
	case FACTOR_FUNCTION_CALL:
		output += " Function Call"
	case EXPRESSION_ADD:
		output += " Add"
	case EXPRESSION_SUBTRACT:
		output += " Subtract"
	case EXPRESSION_MULTIPLY:
		output += " Multiply"
	case EXPRESSION_DIVIDE:
		output += " Divide"
	case STATEMENT_LET_DECLARATION:
		output += " Variable Declaration"
	case STATEMENT_CONST_DECLARATION:
		output += " Constant Declaration"
	case STATEMENT_LET_ASSIGN:
		output += " Variable Assign"
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
