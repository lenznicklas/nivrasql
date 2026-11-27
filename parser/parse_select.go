package parser

import "nivrasql/lexer"

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
