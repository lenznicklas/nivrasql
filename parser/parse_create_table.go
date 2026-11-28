package parser

import "nivrasql/lexer"

func (p *Parser) ParseCreateTable() (*CreateTableStatement, error) {
	stmt := &CreateTableStatement{}

	// CREATE
	if _, err := p.Expect(lexer.Create); err != nil {
		return nil, err
	}

	// TABLE
	if _, err := p.Expect(lexer.Table); err != nil {
		return nil, err
	}

	// Tablename
	table, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}
	stmt.Table = table.Value

	// Left Parenthesis
	if _, err := p.Expect(lexer.LParen); err != nil {
		return nil, err
	}

	// Columns
	coldefs, err := p.ParseColDefs()
	if err != nil {
		return nil, err
	}
	stmt.Columns = coldefs

	// Right Parenthesis
	if _, err := p.Expect(lexer.RParen); err != nil {
		return nil, err
	}

	return stmt, nil
}
