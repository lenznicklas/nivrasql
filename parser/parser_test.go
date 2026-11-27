package parser

import (
	"nivrasql/lexer"
	"reflect"
	"testing"
)

func TestParseSelect(t *testing.T) {
	l := lexer.Lexer{
		Command: "SELECT name, age FROM users WHERE active = TRUE;",
	}

	if err := l.Lex(); err != nil {
		t.Fatal(err)
	}

	p := NewParser(l.Tokens)

	stmt, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	expected := &SelectStatement{
		Table:   "users",
		Columns: []string{"name", "age"},
		Where: &Condition{
			Column:   "active",
			Operator: "=",
			Value:    true,
		},
	}

	if !reflect.DeepEqual(stmt, expected) {
		t.Errorf("expected %+v, got %+v", expected, stmt)
	}
}
