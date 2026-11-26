package lexer

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	l := Lexer{
		Command: "SELECT * FROM users WHERE age >= 18;",
	}

	if err := l.Lex(); err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		{Select, "SELECT"},
		{Asterisk, "*"},
		{From, "FROM"},
		{Identifier, "users"},
		{Where, "WHERE"},
		{Identifier, "age"},
		{GreaterEquals, ">="},
		{Integer, "18"},
		{Semicolon, ";"},
		{EOF, ""},
	}

	if !reflect.DeepEqual(l.Tokens, expected) {
		t.Errorf("expected %v, got %v",
			expected, l.Tokens)
	}
}
