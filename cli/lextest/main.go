package main

import (
	"bufio"
	"fmt"
	"nivrasql/lexer"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("NivraSQL Lexer Test")

	for {
		fmt.Print("lexer> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "exit" {
			break
		}

		l := lexer.Lexer{Command: input}

		if err := l.Lex(); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for n, token := range l.Tokens {
			fmt.Printf(
				"%-4d %-15v %q\n",
				n,
				token.Type,
				token.Value,
			)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
