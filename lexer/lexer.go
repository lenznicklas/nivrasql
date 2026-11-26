package lexer

type TokenType int

const (
	Select TokenType = iota
	Create
	Drop
	Database
	Table
	LParen
	RParen
	Equals
	BangEquals
	Smaller
	Greater
	SmallerEquals
	GreaterEquals
	Identifier
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Command string
	Tokens  []Token
}

func (l *Lexer) Lex() {
	for n, char := range l.Command {
		switch char {
		case '=':
			l.Tokens = append(l.Tokens, Token{Equals, ""})
		case '(':
			l.Tokens = append(l.Tokens, Token{LParen, ""})
		case ')':
			l.Tokens = append(l.Tokens, Token{RParen, ""})
		case '!':
			if l.Peek('=', n) {
				l.Tokens = append(l.Tokens, Token{BangEquals, ""})
			}
		case '<':
			if l.Peek('=', n) {
				l.Tokens = append(l.Tokens, Token{SmallerEquals, ""})
			} else {
				l.Tokens = append(l.Tokens, Token{Smaller, ""})
			}
		case '>':
			if l.Peek('=', n) {
				l.Tokens = append(l.Tokens, Token{GreaterEquals, ""})
			} else {
				l.Tokens = append(l.Tokens, Token{Greater, ""})
			}
		}
	}
}

func (l *Lexer) Peek(char byte, n int) bool {
	if l.Command[n+1] == char && n < len(l.Command)-1 {
		return true
	}
	return false
}
