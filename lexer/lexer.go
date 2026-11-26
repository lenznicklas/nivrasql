package lexer

type TokenType int

const (
	Illegal TokenType = iota
	EOF

	Select
	Create
	Drop
	Database
	Table
	Insert
	Into
	Values
	From
	Where
	Identifier
	Delete
	Update
	Set
	And
	Or
	Not

	Equals
	EqualsEquals
	BangEquals
	Smaller
	SmallerEquals
	Greater
	GreaterEquals
	LParen
	RParen

	Integer
	Float
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	Command  string
	Tokens   []Token
	Position int
}

func (l *Lexer) Lex() {
	for l.Position = 0; l.Position < len(l.Command); l.Position++ {
		char := l.Command[l.Position]

		switch char {
		case '=':
			l.Tokens = append(l.Tokens, Token{Equals, ""})
		case '(':
			l.Tokens = append(l.Tokens, Token{LParen, ""})
		case ')':
			l.Tokens = append(l.Tokens, Token{RParen, ""})
		case '!':
			if l.Peek('=') {
				l.Tokens = append(l.Tokens, Token{BangEquals, ""})
			}
		case '<':
			if l.Peek('=') {
				l.Tokens = append(l.Tokens, Token{SmallerEquals, ""})
			} else {
				l.Tokens = append(l.Tokens, Token{Smaller, ""})
			}
		case '>':
			if l.Peek('=') {
				l.Tokens = append(l.Tokens, Token{GreaterEquals, ""})
			} else {
				l.Tokens = append(l.Tokens, Token{Greater, ""})
			}
		default:
			if IsLetter(char) {
				word, end := l.ReadWord()

				l.Tokens = append(l.Tokens, Token{LookUpKeyword(word), word})
				l.Position = end - 1
			} else if IsDigit(char) {
				token, end := l.ReadNumber()
				l.Tokens = append(l.Tokens, token)
				l.Position = end - 1
			}
		}
	}
}

func (l *Lexer) Peek(char byte) bool {
	if l.Position+1 < len(l.Command) && l.Command[l.Position+1] == char {
		l.Position++
		return true
	}
	return false
}

func (l *Lexer) ReadWord() (string, int) {
	start := l.Position

	for n < len(l.Command) &&
		(IsLetter(l.Command[l.Position]) || IsDigit(l.Command[l.Position])) {
		l.Position++
	}

	return l.Command[start:l.Position], l.Position
}

func (l *Lexer) ReadNumber() (Token, int) {
	start := l.Position
	tokenType := Integer

	for l.Position < len(l.Command) &&
		l.Command[l.Position] == '.' &&
		IsDigit(l.Command[l.Position+1]) {
		tokenType = Float
		l.Position++

		for l.Position < len(l.Command) && IsDigit(l.Command[l.Position]) {
			l.Position++
		}
	}

	return Token{tokenType, l.Command[start:l.Position]}, l.Position
}
