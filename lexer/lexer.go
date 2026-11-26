package lexer

import "fmt"

type Lexer struct {
	Command  string
	Tokens   []Token
	Position int
}

func (l *Lexer) Lex() error {
	l.Tokens = nil
	for l.Position = 0; l.Position < len(l.Command); l.Position++ {
		char := l.Command[l.Position]

		switch char {
		case '=':
			if l.Peek('=') {
				l.AddToken(EqualsEquals, "==")
			} else {
				l.AddToken(Equals, "=")
			}
		case '(':
			l.AddToken(LParen, "(")
		case ')':
			l.AddToken(RParen, ")")
		case ',':
			l.AddToken(Comma, ",")
		case ';':
			l.AddToken(Semicolon, ";")
		case '+':
			l.AddToken(Plus, "+")
		case '-':
			l.AddToken(Minus, "-")
		case '*':
			l.AddToken(Asterisk, "*")
		case '/':
			l.AddToken(Slash, "/")
		case '.':
			l.AddToken(Dot, ".")
		case '!':
			if l.Peek('=') {
				l.AddToken(BangEquals, "!=")
			} else {
				return fmt.Errorf("\"!\" must only be used as \"!=\"")
			}
		case '<':
			if l.Peek('=') {
				l.AddToken(SmallerEquals, "<=")
			} else {
				l.AddToken(Smaller, "<")
			}
		case '>':
			if l.Peek('=') {
				l.AddToken(GreaterEquals, ">=")
			} else {
				l.AddToken(Greater, ">")
			}
		case '\'':
			token, err := l.ReadString()
			if err != nil {
				return err
			}
			l.Tokens = append(l.Tokens, token)
		case ' ', '\n', '\t', '\r':
			continue
		default:
			if IsLetter(char) {
				word := l.ReadWord()
				l.AddToken(LookUpKeyword(word), word)
				l.Position--
			} else if IsDigit(char) {
				token := l.ReadNumber()
				l.Tokens = append(l.Tokens, token)
				l.Position--
			} else {
				return fmt.Errorf(
					"unexpected character %q at position %d",
					char,
					l.Position,
				)
			}
		}
	}
	l.AddToken(EOF, "")
	return nil
}

func (l *Lexer) Peek(char byte) bool {
	if l.Position+1 < len(l.Command) && l.Command[l.Position+1] == char {
		l.Position++
		return true
	}
	return false
}

func (l *Lexer) AddToken(tokenType TokenType, val string) {
	l.Tokens = append(l.Tokens, Token{tokenType, val})
}

func (l *Lexer) ReadWord() string {
	start := l.Position

	for l.Position < len(l.Command) &&
		(IsLetter(l.Command[l.Position]) || IsDigit(l.Command[l.Position])) {
		l.Position++
	}

	return l.Command[start:l.Position]
}

func (l *Lexer) ReadNumber() Token {
	start := l.Position
	tokenType := Integer

	for l.Position < len(l.Command) &&
		IsDigit(l.Command[l.Position]) {
		l.Position++
	}

	if l.Position+1 < len(l.Command) &&
		l.Command[l.Position] == '.' &&
		IsDigit(l.Command[l.Position+1]) {
		tokenType = Float
		l.Position++

		for l.Position < len(l.Command) &&
			IsDigit(l.Command[l.Position]) {
			l.Position++
		}
	}

	return Token{tokenType, l.Command[start:l.Position]}
}

func (l *Lexer) ReadString() (Token, error) {
	l.Position++
	start := l.Position

	for l.Position < len(l.Command) && l.Command[l.Position] != '\'' {
		l.Position++
	}

	if l.Position >= len(l.Command) {
		return Token{}, fmt.Errorf("Unterminated string")
	}

	return Token{String, l.Command[start:l.Position]}, nil
}
