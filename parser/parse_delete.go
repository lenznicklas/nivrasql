package parser

import "nivrasql/lexer"

func (p *Parser) ParseDelete() (*DeleteStatement, error) {
	stmt := &DeleteStatement{}

	// DELETE
	if _, err := p.Expect(lexer.Delete); err != nil {
		return nil, err
	}

	// FROM
	if _, err := p.Expect(lexer.From); err != nil {
		return nil, err
	}

	// Tablename
	name, err := p.Expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}

	stmt.Table = name.Value

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
