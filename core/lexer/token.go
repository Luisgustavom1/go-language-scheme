package lexer

import (
	"fmt"
)

type Token struct {
	Value         string        `json:"value"`
	Kind          TokenKind     `json:"kind"`
	Location      int           `json:"location"`
	LexingContext LexingContext `json:"lexingContext"`
}

type LexingContext struct {
	Source         []rune `json:"source"`
	SourceFileName string `json:"sourceFileName"`
}

type TokenKind uint

const (
	// "(", ")"
	SyntaxToken TokenKind = iota
	// "1", "34"
	IntegerToken
	// "+", "define"
	IdentifierToken
)

func (t Token) Debug(description string) {
	var tokenLine []rune
	var tokenLineLength int
	var tokenColumn int
	var i int

	for i < len(t.LexingContext.Source) {
		char := t.LexingContext.Source[i]

		tokenLine = append(tokenLine, char)

		if i < t.Location {
			tokenColumn++
		}

		if char == '\n' {
			tokenLineLength++

			tokenColumn = 1
			tokenLine = nil
		}

		if i == t.Location {
			tokenLineLength++
			break
		}

		i++
	}

	fmt.Printf("%s at line %d, column %d in file %s\n", description, tokenLineLength, tokenColumn, t.LexingContext.SourceFileName)
	fmt.Println(string(tokenLine))
	for tokenColumn >= 1 {
		fmt.Print(" ")
		tokenColumn--
	}
	fmt.Println("^ near here")
}
