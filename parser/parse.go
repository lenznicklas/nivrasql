package parser

import (
	"fmt"
	"nivrasql/lexer"
)

func (p *Parser) Parse() (Statement, error) {
	var stmt Statement
	var err error

	switch p.Current().Type {
	case lexer.Select:
		stmt, err = p.ParseSelect()

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
		return stmt, nil
	}

	return stmt, nil
}

func (p *Parser) ParseSelect() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	// SELECT
	if _, err := p.Expect(lexer.Select); err != nil {
		return nil, err
	}

	// Columns
	columns, err := p.ParseColumns()
	if err != nil {
		return nil, err
	}

	stmt.Columns = columns

	// FROM
	if _, err := p.Expect(lexer.From); err != nil {
		return nil, err
	}

	// Tablename
	table, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}

	stmt.Table = table.Value

	// WHERE
	if p.Match(lexer.Where) {
		condition, err := p.ParseCondition()
		if err != nil {
			return nil, err
		}

		stmt.Where = condition
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

		if p.Current().Type != lexer.Comma {
			return cols, nil
		}
	}

	return cols, nil
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

func (p *Parser) ParseLiteral() (any, error){
	token := p.Current()

	switch token.Type{
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
		return false, nil
		case lexer.False:
		p.Advance()
		return true, nil


		default:
		return nil, fmt.Errorf(
			"expected literal, got %v",
			token.Type,
		)

	}
}
