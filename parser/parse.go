package parser

import (
	"fmt"
	"os"
	"strings"
	"z/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	i      int
	Root   Node
}

func New(tokens []lexer.Token) Parser {
	return Parser{tokens, 0, Node{}}
}

func (p *Parser) inbounds() bool {
	return p.i < len(p.Tokens) && p.peek().Kind != lexer.EOF
}

func (p *Parser) peek() lexer.Token {
	return p.Tokens[p.i]
}

func (p *Parser) consume() lexer.Token {
	tok := p.peek()
	p.i++
	return tok
}

func (p *Parser) expect(kind lexer.TokenKind) lexer.Token {
	if p.peek().Kind == kind {
		return p.consume()
	} else {
		fmt.Printf("\nError: unexpected \"%s\", expected \"%s\"\n", lexer.Token_kind_to_str(p.peek().Kind), lexer.Token_kind_to_str(kind))
		//fmt.Printf("%s:%d:%d Syntax Error: Expected %s\n", lexer.Token_kind_to_str(kind), filename, p.line, p.col)
		os.Exit(1)
		return lexer.Token{}
	}
}

func (p *Parser) Parse() error {
	// todo:
	//   strings
	//   bools
	//   more and more statements
	//     (if/while/for/fn etc)

	p.Root = Node{ROOT, "", []Node{}}

	for p.inbounds() {
		if p.peek().Kind == lexer.SYMBOL_SEMI_COLON {
			p.consume()
			continue
		}

		S, err := p.parse_statement()

		if err != nil {
			return err
		}

		p.Root.Children = append(p.Root.Children, S)
	}

	return nil
}

func debug_node(node Node, depth int) {
	output := "  " + strings.Repeat("-", depth*2)

	switch node.Kind {
	case FACTOR_LITERAL_NUMBER:
		output += " Number Literal"
	case FACTOR_LITERAL_STRING:
		output += " String Literal"
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
	case STATEMENT_RETURN:
		output += " Return"
	case STATEMENT_BLOCK:
		output += " Block"
	case STATEMENT_FUNCTION_DECLARATION:
		output += " Function Declaration"
	case STATEMENT_WHILE_LOOP:
		output += " While Loop"
	case RETURN_TYPE_NUMBER:
		output += " Number Return Type"
	case RETURN_TYPE_STRING:
		output += " String Return Type"
	case RETURN_TYPE_BOOL:
		output += " Boolean Return Type"
	case PARAMETER_TYPE_STRING:
		output += " String Parameter"
	case PARAMETER_TYPE_NUMBER:
		output += " Number Parameter"
	case PARAMETER_TYPE_BOOL:
		output += " Boolean Parameter"
	case TYPE_STRING:
		output += " String Type"
	case TYPE_NUMBER:
		output += " Number Type"
	case TYPE_BOOL:
		output += " Boolean Type"
	case ROOT:
		output += " Root"
	}

	if node.Value != "" {
		output += fmt.Sprintf(" (%s)", node.Value)
	}

	fmt.Println(output)

	for _, n := range node.Children {
		debug_node(n, depth+1)
	}
}

func (p *Parser) Debug() {
	debug_node(p.Root, 1)
}
