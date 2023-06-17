package main

import (
	"fmt"
	"os"

	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
)

func main() {
	sourceCode, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens := lexer.Lexer([]rune(string(sourceCode)))
	fmt.Println(tokens)
}
