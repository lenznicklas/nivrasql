package parser

import "nivrasql/lexer"

func (p *Parser) ParseInsert() (*InsertStatement, error) {

	stmt := &InsertStatement{}

	// INSERT
	if _, err := p.Expect(lexer.Insert); err != nil {
		return nil, err
	}

	// INTO
	if _, err := p.Expect(lexer.Into); err != nil {
		return nil, err
	}

	// Tablename
	table, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}
	stmt.Table = table.Value

	// VALUES
	if _, err := p.Expect(lexer.Values); err != nil {
		return nil, err
	}

	// Left Parenthesis
	if _, err := p.Expect(lexer.LParen); err != nil {
		return nil, err
	}

	// Values
	values, err := p.ParseValues()
	if err != nil {
		return nil, err
	}
	stmt.Values = values

	// Right Parenthesis
	if _, err := p.Expect(lexer.RParen); err != nil {
		return nil, err
	}

	// SEMICOLON
	if _, err := p.Expect(lexer.Semicolon); err != nil {
		return nil, err
	}

	return stmt, nil
}
