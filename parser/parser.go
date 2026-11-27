package parser

import "nivrasql/lexer"

type Parser struct {
	Tokens   []lexer.Token
	Position int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		Tokens:   tokens,
		Position: 0,
	}
}
