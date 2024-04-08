package parser

import (
	"monkey-lang/ast"
	"monkey-lang/lexer"
	"monkey-lang/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for !p.curTokenIs(token.EOF) {
		stm := p.parseStament()
		if stm != nil {
			program.Statements = append(program.Statements, stm)
		}
		p.nextToken()

	}
	return program
}

func (p *Parser) parseStament() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStament()
	default:
		return nil
	}
}

func (p *Parser) parseLetStament() *ast.LetStatement {
	stm := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stm.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stm
}

func (p *Parser) curTokenIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

func (p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	}
	return false
}
func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}
