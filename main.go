package main

import (
	"fmt"
	"os"

	"github.com/Luisgustavom1/go-language-scheme/core/interpreter"
	"github.com/Luisgustavom1/go-language-scheme/core/lexer"
	"github.com/Luisgustavom1/go-language-scheme/core/parser"
)

func main() {
	sourceCode, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens := lexer.Lexer([]rune(string(sourceCode)))

	var parseIndex int
	var ast = parser.Ast{
		{
			Kind: parser.LiteralValue,
			Literal: &lexer.Token{Value: "begin"},
		},
	}

	for parseIndex < len(tokens) {
		childAst, nextIndex := parser.Parse(tokens, parseIndex)
		ast = append(ast, parser.AstValue{
			Kind: parser.ListValue,
			List: &childAst,
		})

		parseIndex = nextIndex
	} 

	// Optimizations
	// Type checking
	// Are made in this step, over the AST

	interpreter.InitializeBuiltIns()
	value := interpreter.InitAstWalk(ast, map[string]any{})
	fmt.Println("Result: ", value)
}
