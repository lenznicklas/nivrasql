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

	Asterisk
	Comma
	Semicolon
	Plus
	Minus
	Slash
	Dot

	Integer
	Float
	String
	Boolean

	True
	False
	Null
)

type Token struct {
	Type  TokenType
	Value string
}
