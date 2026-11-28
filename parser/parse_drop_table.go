package parser

import "nivrasql/lexer"

func (p *Parser) ParseDropTable() (*DropTableStatement, error) {
	stmt := &DropTableStatement{}

	// DROP
	if _, err := p.Expect(lexer.Drop); err != nil {
		return nil, err
	}

	// TABLE
	if _, err := p.Expect(lexer.Table); err != nil {
		return nil, err
	}

	// Tablename
	name, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}
	stmt.Table = name.Value

	return stmt, nil
}
