package parser

import (
	"PLI/ast"
	"PLI/lexer"
	"PLI/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.NextToken()
	p.NextToken()

	return p

}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
