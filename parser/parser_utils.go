package parser

import (
	"fmt"
	"nivrasql/lexer"
)

func (p *Parser) Current() lexer.Token {
	if p.Position >= len(p.Tokens) {
		return lexer.Token{Type: lexer.EOF}
	}

	return p.Tokens[p.Position]
}

func (p *Parser) Advance() lexer.Token {
	token := p.Current()

	if p.Position < len(p.Tokens) {
		p.Position++
	}
	return token
}

func (p *Parser) Match(t lexer.TokenType) bool {
	if p.Current().Type == t {
		p.Advance()
		return true
	}

	return false
}

func (p *Parser) Expect(t lexer.TokenType) (lexer.Token, error) {
	token := p.Current()

	if token.Type != t {
		return lexer.Token{}, fmt.Errorf(
			"expected %v, got %v at %d",
			t,
			token.Type,
			p.Position,
		)
	}

	p.Advance()
	return token, nil
}
