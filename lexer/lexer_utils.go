package lexer

import "strings"

func LookUpKeyword(word string) TokenType {
	if token, ok := Keywords[strings.ToUpper(word)]; ok {
		return token
	}
	return Identifier
}

func IsLetter(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c == '_'
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
