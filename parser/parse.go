package parser

import (
	"fmt"
	"nivrasql/lexer"
	"strconv"
)

func (p *Parser) Parse() (Statement, error) {
	var stmt Statement
	var err error

	switch p.Current().Type {
	case lexer.Select:
		stmt, err = p.ParseSelect()
	case lexer.Insert:
		stmt, err = p.ParseInsert()

	default:
		return nil, fmt.Errorf(
			"unexpected statement: %v",
			p.Current().Type,
		)
	}

	if err != nil {
		return nil, err
	}

	p.Match(lexer.Semicolon)

	if _, err := p.Expect(lexer.EOF); err != nil {
		return nil, err
	}

	return stmt, nil
}

func (p *Parser) ParseColumns() ([]string, error) {
	var cols []string

	if p.Match(lexer.Asterisk) {
		return []string{"*"}, nil
	}

	for {
		token, err := p.Expect(lexer.Identifier)
		if err != nil {
			return nil, err
		}

		cols = append(cols, token.Value)

		if !p.Match(lexer.Comma) {
			break
		}
	}

	return cols, nil
}

func (p *Parser) ParseValues() ([]any, error) {
	var vals []any

	for {
		token, err := p.Expect(lexer.Identifier)
		if err != nil {
			return nil, err
		}

		vals = append(vals, token.Value)

		if !p.Match(lexer.Comma) {
			break
		}
	}

	return vals, nil
}

func (p *Parser) ParseColDefs() ([]ColumnDefinition, error) {
	var coldefs []ColumnDefinition

	for {
		name, err := p.Expect(lexer.Identifier)
		if err != nil {
			return nil, err
		}

		colType, err := p.ParseType()
		if err != nil {
			return nil, err
		}
		coldefs = append(coldefs, ColumnDefinition{name.Value, colType})

		if !p.Match(lexer.Comma) {
			break
		}

		return coldefs, nil
	}
}

func (p *Parser) ParseType() (string, error) {
	cType := p.Current().Type

	if cType == lexer.Integer {
		return "INT", nil
	}
	if cType == lexer.Text {
		return "TEXT", nil
	}
	if cType == lexer.Float {
		return "FLOAT", nil
	}
	if cType == lexer.Boolean {
		return "BOOL", nil
	}
	p.Advance()

	return "", fmt.Errorf("not a type")
}

func (p *Parser) ParseCondition() (*Condition, error) {
	col, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}

	operator, err := p.ParseOperator()
	if err != nil {
		return nil, err
	}

	value, err := p.ParseLiteral()
	if err != nil {
		return nil, err
	}

	return &Condition{col.Value, operator, value}, nil
}

func (p *Parser) ParseOperator() (string, error) {
	token := p.Current()

	switch token.Type {
	case lexer.Equals,
		lexer.EqualsEquals,
		lexer.BangEquals,
		lexer.Smaller,
		lexer.SmallerEquals,
		lexer.Greater,
		lexer.GreaterEquals:
		p.Advance()
		return token.Value, nil
	}

	return "", fmt.Errorf("expected operator, got %v", token.Value)
}

func (p *Parser) ParseLiteral() (any, error) {
	token := p.Current()

	switch token.Type {
	case lexer.Integer:
		p.Advance()
		return strconv.Atoi(token.Value)
	case lexer.Float:
		p.Advance()
		return strconv.ParseFloat(token.Value, 64)
	case lexer.String:
		p.Advance()
		return token.Value, nil
	case lexer.True:
		p.Advance()
		return true, nil
	case lexer.False:
		p.Advance()
		return false, nil

	default:
		return nil, fmt.Errorf(
			"expected literal, got %v",
			token.Type,
		)
	}
}
