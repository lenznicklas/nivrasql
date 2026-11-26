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

func TestInsert(t *testing.T) {
	l := Lexer{
		Command: "INSERT INTO users VALUES (1, 'Max Mustermann');",
	}

	if err := l.Lex(); err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		{Insert, "INSERT"},
		{Into, "INTO"},
		{Identifier, "users"},
		{Values, "VALUES"},
		{LParen, "("},
		{Integer, "1"},
		{Comma, ","},
		{String, "Max Mustermann"},
		{RParen, ")"},
		{Semicolon, ";"},
		{EOF, ""},
	}

	if !reflect.DeepEqual(l.Tokens, expected) {
		t.Errorf("expected %v, got %v", expected, l.Tokens)
	}
}

func TestNumbers(t *testing.T) {
	l := Lexer{
		Command: "123 3.14 0 999 0.5",
	}

	if err := l.Lex(); err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		{Integer, "123"},
		{Float, "3.14"},
		{Integer, "0"},
		{Integer, "999"},
		{Float, "0.5"},
		{EOF, ""},
	}

	if !reflect.DeepEqual(l.Tokens, expected) {
		t.Errorf("expected %v, got %v", expected, l.Tokens)
	}
}

func TestOperators(t *testing.T) {
	l := Lexer{
		Command: "= == != < <= > >=",
	}

	if err := l.Lex(); err != nil {
		t.Fatal(err)
	}

	expected := []Token{
		{Equals, "="},
		{EqualsEquals, "=="},
		{BangEquals, "!="},
		{Smaller, "<"},
		{SmallerEquals, "<="},
		{Greater, ">"},
		{GreaterEquals, ">="},
		{EOF, ""},
	}

	if !reflect.DeepEqual(l.Tokens, expected) {
		t.Errorf("expected %v, got %v", expected, l.Tokens)
	}
}

func TestInvalidInput(t *testing.T) {
	inputs := []string{
		"'Hello",
		"!",
		"@",
		"#",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			l := Lexer{Command: input}

			if err := l.Lex(); err == nil {
				t.Errorf("expected error for %q", input)
			}
		})
	}
}
